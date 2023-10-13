package feeder

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"datamesh.poc/datafeed-process-service/feeder/model"
)

type ClientFeeder struct {
	url   string
	cache [10]*model.Client
}

func NewClientFeeder(url string) *ClientFeeder {
	return &ClientFeeder{
		url:   url,
		cache: [10]*model.Client{},
	}
}

func (f *ClientFeeder) GetClients() []*model.Client {
	return f.cache[:]
}

func (f *ClientFeeder) Feed(amount int) {
	for i := 0; i < amount; i++ {
		client := model.NewClient()
		body, err := client.Json()
		if err != nil {
			continue
		}
		res, err := http.Post(f.url, "application/json", bytes.NewReader(body))
		if err != nil {
			continue
		}
		if res.StatusCode != 200 {
			body, err = io.ReadAll(res.Body)
			log.Println(string(body))
			continue
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			continue
		}
		var m map[string]string
		err = json.Unmarshal(body, &m)
		if err != nil {
			continue
		}
		client.Id = m["id"]
		f.cache[i%(len(f.cache)-1)] = client
	}
}
