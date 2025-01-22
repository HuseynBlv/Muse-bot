package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func YT(apiKey, query string) string {
	baseURL := "https://www.googleapis.com/youtube/v3/search"
	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("q", query)
	params.Add("type", "video")
	params.Add("videoCategoryId", "10") // Music category
	params.Add("key", apiKey)

	resp, err := http.Get(fmt.Sprintf("%s?%s", baseURL, params.Encode()))
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title string `json:"title"`
			} `json:"snippet"`
		} `json:"items"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	if len(result.Items) == 0 {
		return "No results found."
	}

	video := result.Items[0]
	return fmt.Sprintf("Found: %s\nhttps://www.youtube.com/watch?v=%s", video.Snippet.Title, video.ID.VideoID)
}

func YTurl(apiKey, query string) string {
	baseURL := "https://www.googleapis.com/youtube/v3/search"
	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("q", query)
	params.Add("type", "video")
	params.Add("videoCategoryId", "10") // Music category
	params.Add("key", apiKey)

	resp, err := http.Get(fmt.Sprintf("%s?%s", baseURL, params.Encode()))
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title string `json:"title"`
			} `json:"snippet"`
		} `json:"items"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	if len(result.Items) == 0 {
		return "No results found."
	}

	video := result.Items[0]
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", video.ID.VideoID)
	return url
}

func title(apiKey, query string) string {
	baseURL := "https://www.googleapis.com/youtube/v3/search"
	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("q", query)
	params.Add("type", "video")
	params.Add("videoCategoryId", "10") // Music category
	params.Add("key", apiKey)

	resp, err := http.Get(fmt.Sprintf("%s?%s", baseURL, params.Encode()))
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title string `json:"title"`
			} `json:"snippet"`
		} `json:"items"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	if len(result.Items) == 0 {
		return "No results found."
	}

	video := result.Items[0]
	title := fmt.Sprintf("%s", video.Snippet.Title)
	return title
}
