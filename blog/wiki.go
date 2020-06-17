package main

import (
	"fmt"
	"io/ioutil"
)

// Structure defining a webpage
type Page struct {
	Title string
	Body []byte
}

// Save Page Method
// This is a method named save that takes as its receiver p,
// a pointer to Page.
func (p *Page) save() error {
	filename := p.title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load Page Method
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a simple Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}