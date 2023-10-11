package feeder

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"datamesh.poc/datafeed-process-service/feeder/model"
)

type PersonFeeder struct {
	url string
}

func NewPersonFeeder(url string) *PersonFeeder {
	return &PersonFeeder{
		url: url,
	}
}

func (f *PersonFeeder) Feed(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		p := model.NewPerson()

		body, err := json.Marshal(p)
		if err != nil {
			log.Println(err)
			continue
		}
		resp, err := http.Post(f.url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()
		log.Println(resp.StatusCode)
		log.Println(resp.Body)
	}
	wg.Done()
}
