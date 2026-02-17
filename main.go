package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

type Asset struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Change string `json:"change"`
}

type All struct {
	Time        time.Time
	Status      string
	PrimeAssets map[string]Asset
	Others      map[string]Asset
}

func getDovizResponse(ctx context.Context) (*http.Response, error) {

	url := "https://canlidoviz.com/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("Request error to canlidoviz.com!")
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 OPR/94.0.0.0")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return response, errors.New(response.Status)
	}
	return response, ctx.Err()

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var all All
	all.PrimeAssets = make(map[string]Asset)
	all.Others = make(map[string]Asset)
	// context for 3rd party url timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	response, err := getDovizResponse(ctx)
	if err != nil {
		log.Println(err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
	}

	doc.Find("table").Each(func(tableIndex int, table *goquery.Selection) {
		if table.Find("span[dt='amount']").Length() > 0 &&
			table.Find("span[dt='change']").Length() > 0 {
			table.Find("tbody tr, tr.table-row-md, tr.table-row").Each(func(rowIndex int, row *goquery.Selection) {
				// Extract data as above
				codeSel := row.Find(".table-code, .table-name, span.items-center")
				code := strings.TrimSpace(codeSel.First().Text())

				amountSel := row.Find("span[dt='amount']")
				amount := strings.TrimSpace(amountSel.Text())

				changeSel := row.Find("span[dt='change']")
				change := strings.TrimSpace(changeSel.Text())

				if code != "" || amount != "" || change != "" {

					all.Others[code] = Asset{Name: code, Price: amount, Change: change}
				}
			})
		}
	})

	title := doc.Find("span.table-title").FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.Contains(strings.TrimSpace(s.Text()), "Piyasa Özeti")
	}).First()

	if title.Length() == 0 {
		fmt.Println("Title not found")
		return
	}

	container := title.Parent()

	table := container.Find("table.w-full.flex.flex-col")

	if table.Length() == 0 {
		fmt.Println("Table not found")
		return
	}

	// Extract all data

	table.Find("tr.table-row").Each(func(i int, row *goquery.Selection) {
		name := strings.TrimSpace(row.Find("span.table-name").Text())
		price := strings.TrimSpace(row.Find("span.table-price").Text())
		change := strings.TrimSpace(row.Find("span.table-change").Text())

		all.PrimeAssets[name] = Asset{Name: name, Price: price, Change: change}
	})
	time := time.Now()
	all.Time = time
	all.Status = response.Status
	if err := json.NewEncoder(w).Encode(all); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	log.Println("Server started on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
