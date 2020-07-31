package main

import (
	crss "github.com/DmitriyDev/go-RssCommunicator"
	"strings"
)

type RssStream struct {
	City     string
	Category string
}

type RssFacade struct {
	Domain       string
	Communicator crss.Communicator
}

func (rf RssFacade) ReadStream(stream RssStream) []Item {

	request := crss.CommunicationRequest{
		rf.getStreamUrl(stream),
		getHeaders(),
	}
	rawContent, _ := rf.Communicator.Execute(request)
	streamContent, _ := unmarshal(rawContent)

	return streamContent
}

func (rf RssFacade) getStreamUrl(stream RssStream) string {

	streamUrl := rf.Domain

	parameter := []string{}
	if stream.Category != "" {
		parameter = append(parameter, "category="+stream.Category)
	}

	if stream.City != "" {
		parameter = append(parameter, "city="+stream.City)
	}

	if len(parameter) > 0 {
		streamUrl += "?" + strings.Join(parameter, "&")
	}

	return streamUrl
}
