package sheets

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	credentialsFilePath = "CREDENTIALS_FILE_PATH"
	noDataFound         = "No data found, invalid URL"
	readRange           = "Sheet1!A1:ZZ"
)

func GetSheetRawData(sheetURL string) [][]interface{} {
	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentialsFile(os.Getenv(credentialsFilePath)))
	if err != nil {
		panic(err)
	}

	spreadsheetId, err := getSheetID(sheetURL)
	if err != nil {
		return [][]interface{}{}
	}

	readRange := "Sheet1!A1:ZZ"
	log.Printf("getting sheet for sheet ID: %s", spreadsheetId)

	resp, err := client.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		panic(err)
	}

	return resp.Values
}

func getSheetID(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	path := parsedURL.Path
	parts := strings.Split(path, "/")
	id := parts[len(parts)-2]

	return id, nil
}

func GetSheetData(sheetURL string) []string {
	// Set up the Google Sheets API client

	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentialsFile(os.Getenv(credentialsFilePath)))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Specify the spreadsheet ID and range of cells to read
	spreadsheetId, err := getSheetID(sheetURL)
	if err != nil {
		return []string{noDataFound}
	}

	log.Printf("getting sheet for sheet ID: %s", spreadsheetId)

	// Call the API to fetch the data
	resp, err := client.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		// Handle error
		panic(err)
	}

	results := make([]string, 0, len(resp.Values))
	// Parse the response data as needed
	if len(resp.Values) > 0 {
		results = append(results, "Data found")
		for _, row := range resp.Values {
			results = append(results, fmt.Sprintf("%s", row))
		}
	} else {
		results = append(results, "No data found.")

	}
	log.Println(results)

	return results
}
