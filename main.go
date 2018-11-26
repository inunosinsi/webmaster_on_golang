package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	webmasters "google.golang.org/api/webmasters/v3"
)

func main() {
	//認証資格情報を取得するHTTPクライアントを返す
	client, err := google.DefaultClient(
		oauth2.NoContext,
		"https://www.googleapis.com/auth/webmasters.readonly")
	if err != nil {
		log.Fatalf("Unable to read client : %v", err)
	}

	// service, err := searchconsole.New(client)
	service, err := webmasters.New(client)
	if err != nil {
		log.Fatalf("Unable to Access Search Console: %v", err)
	}

	analyticsService := webmasters.NewSearchanalyticsService(service)

	yesterday := time.Now().AddDate(0, 0, -1)
	oneWeekBefore := yesterday.AddDate(0, 0, -7)
	format := "2006-01-02"

	req := &webmasters.SearchAnalyticsQueryRequest{
		StartDate:  oneWeekBefore.Format(format),
		EndDate:    yesterday.Format(format),
		Dimensions: []string{"query"},
		RowLimit:   10,
	}
	resp, err := analyticsService.Query("https://saitodev.co", req).Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range resp.Rows {
		fmt.Println(row)
	}
}
