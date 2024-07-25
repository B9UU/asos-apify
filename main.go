package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Payload struct {
	Offset int    `json:"start"`
	Max    int    `json:"last_item"`
	Query  string `json:"search_query"`
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

// TODO: need to handle payload more than 5MB
func (i *Scraper) Output(data AsosResp) error {
	url := fmt.Sprintf("https://api.apify.com/v2/datasets/%s/items?token=%s", i.DatasetId, i.Token)
	fmt.Println("Number of Items: ", data.ItemCount)

	dataM, err := json.Marshal(data.Products)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(dataM)
	fmt.Println(reader.Len())
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	dd, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		fmt.Println("ERROR: unable to add to dataset: ", string(dd))
	}
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

func Asos(q, offset, limit string) (AsosResp, error) {

	baseUrl := "https://www.asos.com/api/product/search/v2/"
	params := url.Values{}
	params.Add("includeNonPurchasableTypes", "restocking")
	params.Add("offset", offset)
	params.Add("q", q)
	params.Add("store", "ROW")
	params.Add("lang", "en-GB")
	params.Add("currency", "USD")
	params.Add("limit", limit)

	urll, err := url.Parse(baseUrl)
	if err != nil {
		return AsosResp{}, fmt.Errorf("Unable to parse url: %v", err)
	}
	urll.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", urll.String(), nil)
	if err != nil {
		return AsosResp{}, fmt.Errorf("Unable to structure the request: %v", err)
	}
	req.Header.Add("authority", "www.asos.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.6")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return AsosResp{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return AsosResp{}, err
	}
	var dd AsosResp
	json.NewDecoder(res.Body).Decode(&dd)
	return dd, nil

}

func main() {
	log.Println("Example actor written in Go.")
	scrp, err := NewScraper()
	if err != nil {
		log.Fatal(err)
	}
	data, err := Asos("top", "0", "200")
	if err != nil {
		log.Fatal(err)
	}
	err = scrp.Output(data)
}
