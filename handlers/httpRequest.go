package handlers

import (
	"facebookfix/engine"
	"io"
	"net/http"
)

func FacebookRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	engine.SetHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	request, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return request, nil
}
