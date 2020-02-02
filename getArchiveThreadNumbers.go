package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getArchiveThreadNumbers(board string) *list.List {
	// Getting the html file of the archive
	response, _ := http.Get("http://boards.4chan.org/" + board + "/archive")
	byteInput, _ := ioutil.ReadAll(response.Body)
	stringBody := string(byteInput)
	response.Body.Close()

	// Setting up the html tokenizer
	htmlReader := strings.NewReader(stringBody)
	htmlTokenizer := html.NewTokenizer(htmlReader)

	// Initializing currentToken
	currentToken := htmlTokenizer.Next()

	// Initializing the list with the thread numbers and the extractionFinished variable
	threadNumbers := list.New()
	extractionFinished := false

	// Search for the table (and do the other stuff)
	searchTable(&currentToken, htmlTokenizer, &extractionFinished, threadNumbers)

	if currentToken == html.ErrorToken {
		// Unexpected error
		fmt.Println(RedColorRaw + "An unexpected error occured: " + htmlTokenizer.Err().Error() + ResetColorRaw)
	}

	return threadNumbers
}

func searchTokens(currentToken *html.TokenType, htmlTokenizer *html.Tokenizer, extractionFinished *bool, threadNumbers *list.List) {

	for *currentToken != html.ErrorToken && !*extractionFinished {
		if htmlTokenizer.Token().DataAtom == atom.Td && *currentToken == html.StartTagToken {
			*currentToken = htmlTokenizer.Next()
			if *currentToken == html.TextToken {
				threadNumbers.PushBack(htmlTokenizer.Token().Data)

				// Search for the end of the row: currentToken is closing tr tag after exiting (or error)
				for *currentToken != html.ErrorToken && (*currentToken != html.EndTagToken || htmlTokenizer.Token().DataAtom != atom.Tr) {
					*currentToken = htmlTokenizer.Next()
				}

				*currentToken = htmlTokenizer.Next()

				if htmlTokenizer.Token().DataAtom == atom.Tbody && *currentToken == html.EndTagToken {
					*extractionFinished = true
				}
			}

		} else {
			*currentToken = htmlTokenizer.Next()
		}

	}

}

func searchTr(currentToken *html.TokenType, htmlTokenizer *html.Tokenizer, extractionFinished *bool, threadNumbers *list.List) {
	for *currentToken != html.ErrorToken && !*extractionFinished {
		if htmlTokenizer.Token().DataAtom == atom.Tr {
			searchTokens(currentToken, htmlTokenizer, extractionFinished, threadNumbers)
		}
		*currentToken = htmlTokenizer.Next()
	}
}

func searchTbody(currentToken *html.TokenType, htmlTokenizer *html.Tokenizer, extractionFinished *bool, threadNumbers *list.List) {
	for *currentToken != html.ErrorToken && !*extractionFinished {
		if htmlTokenizer.Token().DataAtom == atom.Tbody {
			searchTr(currentToken, htmlTokenizer, extractionFinished, threadNumbers)
		}
		*currentToken = htmlTokenizer.Next()
	}
}

func searchTable(currentToken *html.TokenType, htmlTokenizer *html.Tokenizer, extractionFinished *bool, threadNumbers *list.List) {
	for *currentToken != html.ErrorToken && !*extractionFinished {
		if htmlTokenizer.Token().DataAtom == atom.Table {

			searchTbody(currentToken, htmlTokenizer, extractionFinished, threadNumbers)

		}
		*currentToken = htmlTokenizer.Next()
	}
}
