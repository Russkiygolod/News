// Пакет для работы с RSS-потоками.
package rss

import (
	"News/pkg/storage"
	"encoding/xml"
	"io"
	"net/http"
)

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Chanel  Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Link        string `xml:"link"`
}

// Parse читает rss-поток и возвращет массив новостей приведенных к типу []storage.Posts
func ParseRss(url string) ([]storage.Posts, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feed Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}

	var Posts []storage.Posts
	var post storage.Posts
	for _, item := range feed.Chanel.Items {
		post.Title = item.Title
		post.Description = item.Description
		post.PubDate = item.PubDate
		post.Link = item.Link
		Posts = append(Posts, post)
	}
	return Posts, nil
}
