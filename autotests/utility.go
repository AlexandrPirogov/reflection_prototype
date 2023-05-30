package main

import (
	"io"
	"net/http"
)

func Post(url string, jwt string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	if err != nil {
		return nil, err
	}

	c := http.Client{}
	return c.Do(req)
}

func Get(url string, jwt string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	if err != nil {
		return nil, err
	}

	c := http.Client{}
	return c.Do(req)
}
