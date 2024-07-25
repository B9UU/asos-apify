package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Payload struct {
	Url string `json:"url"`
}
type Scraper struct {
	Key       string
	Token     string
	DatasetId string
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

func (i *Scraper) Output() error {
	url := fmt.Sprintf("https://api.apify.com/v2/datasets/%s/items?token=%s", i.DatasetId, i.Token)
	// body, err := json.Marshal(dd)
	// if err != nil {
	// 	return err
	// }
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(`["laqo":"test"]`))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func NewScraper() (*Scraper, error) {
	scrp := Scraper{}
	scrp.Key, scrp.Token, scrp.DatasetId = os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"), os.Getenv("APIFY_TOKEN"), os.Getenv("APIFY_DEFAULT_DATASET_ID")
	if scrp.Key == "" || scrp.Token == "" || scrp.DatasetId == "" {
		return nil, errors.New("Missing required env vars")
	}
	err := scrp.Input()
	if err != nil {
		return nil, err
	}
	return &scrp, nil
}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Println("Example actor written in Go.")
	scrp, err := NewScraper()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(scrp.Payload.Url)
	err = scrp.Output()
	fmt.Println(scrp)
}
