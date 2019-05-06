package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type musicData struct {
	Data []album `xml:"album"`
}

type album struct {
	Title string `xml:"title"`
	Year  int    `xml:"year"`
	Songs []song `xml:"song"`
}

type song struct {
	Title    string `xml:"title"`
	Artist   string `xml:"artist"`
	Duration string `xml:"duration"`
}

func main() {
	url := "https://gist.githubusercontent.com/yaron/1a1df8cdc71612c24fedf8562f4c0083/raw/0b9fa04ecd96581abaea478eb6799c387d56aa16/gistfile1.txt"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data musicData
	err = xml.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, album := range data.Data {
		for _, song := range album.Songs {
			fmt.Printf("%s by %s on album %s\n", song.Title, song.Artist, album.Title)
		}
	}
}
