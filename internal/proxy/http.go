package proxy

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

type HttpProxy struct {
	Headers http.Header
}

func (hp *HttpProxy) AddHeader(key, value string) {
	hp.Headers.Add(key, value)
}

func constructQueryString(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	return values.Encode()
}

func (hp *HttpProxy) makeRequest(method, baseUrl string, params map[string]string, body []byte) ([]byte, error) {
	fullURL := baseUrl + "?" + constructQueryString(params)

	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header = hp.Headers

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (hp *HttpProxy) GET(baseUrl string, params map[string]string) ([]byte, error) {
	return hp.makeRequest("GET", baseUrl, params, nil)
}

func (hp *HttpProxy) POST(baseUrl string, params map[string]string, body []byte) ([]byte, error) {
	return hp.makeRequest("POST", baseUrl, params, body)
}

func (hp *HttpProxy) PUT(baseUrl string, params map[string]string, body []byte) ([]byte, error) {
	return hp.makeRequest("PUT", baseUrl, params, body)
}

func (hp *HttpProxy) DELETE(baseUrl string, params map[string]string) ([]byte, error) {
	return hp.makeRequest("DELETE", baseUrl, params, nil)
}
