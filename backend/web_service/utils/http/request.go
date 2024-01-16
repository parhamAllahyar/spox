package http

import (
	"io"

	"net/http"
)

func GET(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	
	body, err := io.ReadAll(resp.Body)
	sb := string(body)
	return sb, nil

}
