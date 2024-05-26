package main

import (
	"News/pkg/api"
	parseurl "News/pkg/parseUrl"
	"News/pkg/storage"
	"News/pkg/storage/postgres"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var posts = make(chan []storage.Posts)
var errs = make(chan error)

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
	//запускаем чтение и раскодирование из config файла
	Urls, Period := parseurl.Read("./config.json")
	//идем по url и запукаем чтение Rss
	for _, v := range Urls {
		fmt.Println(v)
		go parseurl.Parse(v, posts, errs, Period)
	}
	//отправляем данные из канала в БД
	go func() {
		for p := range posts {
			db.NewPost(p)
		}
	}()
	//вывод ошибок
	go func() {
		for err := range errs {
			fmt.Println(err)
		}
	}()
	api := api.New(*db)
	http.ListenAndServe(":80", api.Router())
}
