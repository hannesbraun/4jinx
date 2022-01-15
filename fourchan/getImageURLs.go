package fourchan

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/hannesbraun/4jinx/util"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func GetImageURLs(board string, threadNumber string) *list.List {
	// Getting the html file of the thread
	response, _ := http.Get("http://boards.4chan.org/" + board + "/thread/" + threadNumber)
	byteInput, _ := ioutil.ReadAll(response.Body)
	stringBody := string(byteInput)
	response.Body.Close()

	document, _ := html.Parse(strings.NewReader(stringBody))
	threadNode, err := getThreadNode(document)
	if err != nil {
		fmt.Println(util.RedColorRaw + err.Error() + util.ResetColorRaw)
		return nil
	}

	// Initializing the list with the thread numbers and the extractionFinished variable
	imageURLs := list.New()

	for current := threadNode.FirstChild; current != nil; current = current.NextSibling {
		innerDivNode := current.LastChild
		if innerDivNode != nil {

			fileNode, err := getFileNode(innerDivNode)
			if err != nil {
				// fmt.Println(err.Error())
				continue
			}

			imageNode := fileNode.LastChild
			if imageNode.Attr != nil && len(imageNode.Attr) >= 2 {
				imageURLs.PushBack(imageNode.Attr[1].Val)
			}
		}
	}

	return imageURLs
}

func getThreadNode(doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" && n.Attr != nil {
			if n.Attr[0].Val == "thread" {
				b = n
			}
		}

		for current := n.FirstChild; current != nil; current = current.NextSibling {
			f(current)
		}
	}

	f(doc)

	if b != nil {
		return b, nil
	}
	return nil, errors.New("missing <div class=\"thread\"> in the node tree")
}

func getFileNode(doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" && n.Attr != nil {
			if n.Attr[0].Val == "file" {
				b = n
			}
		}

		for current := n.FirstChild; current != nil; current = current.NextSibling {
			f(current)
		}
	}
	f(doc)

	if b != nil {
		return b, nil
	}
	return nil, errors.New("missing <div class=\"file\"> in the node tree")
}
