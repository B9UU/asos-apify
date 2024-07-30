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
	"sync"
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
	Items     int
	mu        sync.Mutex
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
		return nil
	}
	for _, d := range data.Products {
		fmt.Printf("%+v\n", d)
	}
	i.mu.Lock()
	defer i.mu.Unlock()
	i.Items += len(data.Products)
	fmt.Println(i.Items, "Items Exported")
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
	scrp, err := NewScraper()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	ch := make(chan struct{}, 10)
	for i := scrp.Offset; i < scrp.Max; i += 200 {
		// for i := 0; i < 300; i += 200 {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int) {
			defer func() {
				wg.Done()
				<-ch
			}()
			data, err := Asos("top", fmt.Sprint(i), "200")
			if err != nil {
				log.Println(err)
			}
			err = scrp.Output(data)
			if err != nil {
				log.Println(err)
			}
		}(i)
	}
	wg.Wait()
}
