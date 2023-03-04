package sheets

import (
	"context"
	"fmt"
	"os"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const CREDENTIALS_FILE_PATH = "CREDENTIALS_FILE_PATH"

func GetSheetData() string {
	// Set up the Google Sheets API client

	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentialsFile(os.Getenv(CREDENTIALS_FILE_PATH)))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Specify the spreadsheet ID and range of cells to read
	spreadsheetId := "1DUulwcqZJLdUTXk4Z5l7l4lVsPYBitCGph4gP7T65pg"
	readRange := "Sheet1!A1:D10"

	// Call the API to fetch the data
	resp, err := client.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		// Handle error
		panic(err)
	}

	var builder strings.Builder

	// Parse the response data as needed
	if len(resp.Values) > 0 {
		builder.WriteString("Data found:")
		for _, row := range resp.Values {
			builder.WriteString(fmt.Sprintf("%s\n", row))
		}
	} else {
		builder.WriteString("No data found.")
	}

	return builder.String()
}
