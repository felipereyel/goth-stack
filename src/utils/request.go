package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Headers map[string]string

type Response struct {
	StatusCode int
	Body       []byte
}

type HTTPClient struct {
	Headers Headers
	BaseUrl string
}

func (c *HTTPClient) Request(method string, path string, body []byte) (Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseUrl, path)
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return Response{}, errors.New("failed to create request: " + err.Error())
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	client := http.Client{Timeout: 30 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, errors.New("failed to perform request: " + err.Error())
	}

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, errors.New("could not read response body: " + err.Error())
	}

	return Response{
		StatusCode: res.StatusCode,
		Body:       rawBody,
	}, nil
}
