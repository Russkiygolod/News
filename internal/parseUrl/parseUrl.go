package parseurl

import (
	"News/internal/rss"
	posts "News/pkg/model"
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	URLS   []string `json:"rss"`
	Period int      `json:"request_period"`
}

// чтение и раскодирование файла конфигурации пример адреса: "./config.json"
func Read(addres string) (urls []string, perion int) {
	var config Config
	ReadFile, err := os.ReadFile(addres)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(ReadFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config.URLS, config.Period
}

// чтение Rss, приведение к типу []storage.Posts, отправка данных в канал
func Parse(url string, posts chan<- []posts.Posts, errs chan<- error, period int) {
	for {
		news, err := rss.ParseRss(url)
		if err != nil {
			errs <- err
			continue
		}
		posts <- news
		time.Sleep(time.Second * time.Duration(period))
	}
}
