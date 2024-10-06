package main

import (
	"fmt"
)

func main() {
	bh := NewBrowserHistory()

	bh.VisitNewPage("google.com")
	bh.VisitNewPage("quera.org")
	bh.VisitNewPage("stackoverflow.com")

	url := bh.GetCurrentURL()
	fmt.Println("Current URL:", url) // Output: Current URL: stackoverflow.com

	bh.Back()
	url = bh.GetCurrentURL()
	fmt.Println("After Back, Current URL:", url) // Output: After Back, Current URL: quera.org

	bh.Forward()
	url = bh.GetCurrentURL()
	fmt.Println("After Forward, Current URL:", url) // Output: After Forward, Current URL: stackoverflow.com

	bh.VisitNewPage("golang.org")
	url = bh.GetCurrentURL()
	fmt.Println("After Visiting New Page, Current URL:", url) // Output: After Visiting New Page, Current URL: golang.org

	err := bh.Forward()
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: no next page
	}

	bh.ClearHistory()
	err = bh.Back()
	if err != nil {
		fmt.Println("After Clearing History, Error:", err) // Output: After Clearing History, Error: no previous page
	}

	url = bh.GetCurrentURL()
	fmt.Println("Current URL after clearing history:", url) // Output: Current URL after clearing history:
}

// Node represents a node in the doubly linked list
type Node struct {
	url  string
	prev *Node
	next *Node
}

// BrowserHistory manages the browsing history using a doubly linked list
type BrowserHistory struct {
	current *Node
}

func NewBrowserHistory() *BrowserHistory {
	return nil
}

func (bh *BrowserHistory) VisitNewPage(url string) {
}

func (bh *BrowserHistory) Back() error {
	return nil
}

func (bh *BrowserHistory) Forward() error {
	return nil
}

func (bh *BrowserHistory) ClearHistory() {
}

func (bh *BrowserHistory) GetCurrentURL() string {
	return ""
}
