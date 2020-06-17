package main

import (
	"fmt"
)

// TODOs
// add a function that will insert a new page, after a given page
// add a function that will deleta page.

// Linked List
type storyPage struct {
	text string
	// Don't want to store an entire "storyPage"
	// so we use a pointer.
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addToEnd(text string) {
	// Walk down the linked list until we find nil for next page
	for page.nextPage != nil {
		page = page.nextPage
		// Once loop is done we're at the end of the list
	}
	page.nextPage = &storyPage{text, nil}
}

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// Delete (we'll come back)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page1.addToEnd("You are alone, and you need to find the sacred helmet before the bad guys do.")
	page1.addToEnd("You see a troll ahead.")

	page1.addAfter("Testing AddAfter method.")
	page1.playStory()

	// Functions - has return value - may also execute commands
	// Procedures - has no return value, just executes commands
	// Methods - functions attached to a struct/object/etc
}
