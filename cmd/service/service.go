package main

import (
	"News/pkg/api"
	"News/pkg/rss"
	"News/pkg/storage"
	"News/pkg/storage/postgres"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	URLS   []string `json:"rss"`
	Period int      `json:"request_period"`
}

type Service struct {
	Rss           []string
	RequestPeriod int
	db            postgres.Store
	Posts         storage.Posts
	//m             sync.Mutex
	// err           error
}

func New() *Service {
	service := Service{}
	service.db = postgres.Store{}
	service.Posts = storage.Posts{}
	//service.RequestPeriod =
	//service.Rss =
	return &service
}

func main() {

	//////////// Реляционная БД PostgreSQL //////
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "News"))
	if err != nil {
		log.Fatal(err)
	}
	// закрываем подключение
	defer conn.Close()
	// Проверка соединения с БД
	err = conn.Ping()
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}
	db := postgres.New(conn)

	// чтение и раскодирование файла конфигурации
	ReadFile, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config config
	err = json.Unmarshal(ReadFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	posts := make(chan []storage.Posts)
	errs := make(chan error)
	for _, v := range config.URLS {
		fmt.Println(v)
		go parseURL(v, posts, errs, config.Period)
	}
	go func() {
		for p := range posts {
			db.NewPost(p)
		}
	}()

	go func() {
		for err := range errs {
			fmt.Println(err)
		}
	}()
	api := api.New(*db)
	http.ListenAndServe(":80", api.Router())
	//wg.Wait()
}
func parseURL(url string, posts chan<- []storage.Posts, errs chan<- error, period int) {
	for {
		news, err := rss.Parse(url)
		if err != nil {
			errs <- err
			continue
		}
		posts <- news
		time.Sleep(time.Second * time.Duration(period))
	}
}
