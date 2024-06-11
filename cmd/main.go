package main

import (
	"News/internal"
	"News/internal/api"
	parseurl "News/internal/parseUrl"
	"News/internal/postgres"
	posts "News/pkg/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var Posts = make(chan []posts.Posts)
var errs = make(chan error)

func main() {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "News"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	postgresDB := postgres.New(conn)
	//запускаем чтение и раскодирование из config файла Urls
	Urls, Period := parseurl.Read("./News/internal/config/config.json")
	//идем по url и запукаем чтение Rss
	for _, url := range Urls {
		fmt.Println(url)
		go parseurl.Parse(url, Posts, errs, Period)
	}
	//отправляем данные из канала в БД
	go func() {
		for p := range Posts {
			postgresDB.NewPost(p)
		}
	}()
	//вывод ошибок
	go func() {
		for err := range errs {
			fmt.Println(err)
		}
	}()
	internal := internal.New(*postgresDB)
	appi := api.New(internal)
	log.Fatal(http.ListenAndServe(":80", appi.Router()))
}

// для коммита
