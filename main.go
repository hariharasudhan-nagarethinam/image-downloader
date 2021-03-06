package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ImageURL struct {
	Id          string `json:"id"`
	DownloadURL string `json:"download_url"`
}

func getImageURLs(imageRepositoryURL string) ([]ImageURL, error) {
	if response, err := http.Get(imageRepositoryURL); err != nil {
		return []ImageURL{}, err
	} else {
		responseData, _ := ioutil.ReadAll(response.Body)
		var imageURLs []ImageURL
		json.Unmarshal(responseData, &imageURLs)
		return imageURLs, nil
	}
}

func downloadImage(imageUrl string) {
	if _, err := http.Get(imageUrl); err != nil {
		return
	} else {
		fmt.Printf("%s downloaded \n", imageUrl)
		return
	}
}

func main() {
	const imageRepositoryURL = "https://picsum.photos/v2/list"
	urls, err := getImageURLs(imageRepositoryURL)
	if err != nil {
		panic(err)
	}

	for _, data := range urls {
		downloadImage(data.DownloadURL)
	}
}
