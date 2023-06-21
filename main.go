package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

type Channel struct {
	Items []Item `xml:"channel>item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

func main() {
	if res, err := http.Get("https://www.perfil.com/feed/politica"); err != nil {
		fmt.Println("Error retrieving resource:", err)
		os.Exit(1)
	} else {
		channel := Channel{}
		if err := xml.NewDecoder(res.Body).Decode(&channel); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		} else if len(channel.Items) != 0 {

			for _, item := range channel.Items {
				fmt.Println(item.Title)
			}
		}
	}
}
