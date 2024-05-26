package parseurl

import (
	"News/pkg/rss"
	"News/pkg/storage"
	"encoding/json"
	"log"
	"os"
	"time"
)

type config struct {
	URLS   []string `json:"rss"`
	Period int      `json:"request_period"`
}

var posts = make(chan []storage.Posts)
var errs = make(chan error)

// чтение и раскодирование файла конфигурации пример адреса: "./config.json"
func Conifg(addres string) config {
	var config config
	ReadFile, err := os.ReadFile(addres)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(ReadFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
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
