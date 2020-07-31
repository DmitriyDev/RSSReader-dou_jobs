package main

import (
	"encoding/xml"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Categories  []string `xml:"category"`
	PubDate     string   `xml:"pubDate"`
	Guid        string   `xml:"guid"`
	Media       string   `xml:"media"`
	Creator     string   `xml:"dc:creator"`
}

func unmarshal(content string) ([]Item, error) {

	rss := Rss{}
	if err := xml.Unmarshal([]byte(content), &rss); err != nil {
		return []Item{}, err
	}
	return rss.Channel.Items, nil
}
