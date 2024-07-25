package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Payload struct {
	Url string `json:"url"`
}
type Scraper struct {
	Key   string
	Token string
	Payload
}

func (i *Scraper) Input() error {
	url := fmt.Sprintf("https://api.apify.com/v2/key-value-stores/%s/records/INPUT?token=%s", i.Key, i.Token)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&i.Payload); err != nil {
		return err
	}
	return nil

}

func NewScraper() (*Scraper, error) {
	scrp := Scraper{}
	scrp.Key, scrp.Token = os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"), os.Getenv("APIFY_TOKEN")
	if scrp.Key == "" || scrp.Token == "" {
		return nil, errors.New("Missing required env vars")
	}
	err := scrp.Input()
	if err != nil {
		return nil, err
	}
	return &scrp, nil
}

func main() {
	log.Println("Example actor written in Go.")
	_, err := NewScraper()
	if err != nil {
		log.Fatal(err)
	}
	_, err = http.Get("https://www.ibrahimboussaa.com")
	if err != nil {
		log.Fatal(err)
		return
	}
}
