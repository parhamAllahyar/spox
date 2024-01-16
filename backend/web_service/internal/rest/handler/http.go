package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var urls = map[string]string{
	"user":      "http://userapp:8085/",
	"workspace": "http://workspaceapp:82/",
	"asset":     "",
	"chat":      "",
	"board":     "",
}

func ForwardURL(w http.ResponseWriter, r *http.Request, targetURL string, service string) {

	proxyURL := urls[service] + targetURL

	fmt.Println(proxyURL)

	// Create a new request to the target URL
	req, err := http.NewRequest(r.Method, proxyURL, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("err A is", err)
	fmt.Println("req is", err)

	// Copy headers from the original request to the new request
	req.Header = make(http.Header)
	for h, val := range r.Header {
		req.Header[h] = val
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Make the request to the target URL
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers from the target URL to the original response
	for h, val := range resp.Header {
		w.Header()[h] = val
	}

	// Set the status code for the original response
	w.WriteHeader(resp.StatusCode)

	// Copy the response body from the target URL to the original response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("err D is", err)
	fmt.Println("body is", body)

	w.Write(body)

}
