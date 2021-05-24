package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	p1 := &Page{Title: "LoremIpsum", Body: []byte("Hipster ipsum dolor sit consectur")}
	p1.save()

	p2, _ := loadPage("LoremIpsum")
	fmt.Print(string(p2.Body))

}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
