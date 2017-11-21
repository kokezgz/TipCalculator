package Services

import (
	"bytes"
	"net/http"
	"time"
)

type IclientService interface {
	CallService(url string, endPoint string, contentType string) ([]byte, error)
}

type ClientService struct {
}

func (c ClientService) CallService(url string, endPoint string, contentType string) ([]byte, error) {

	urlString := url + endPoint
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	r, _ := http.NewRequest("GET", urlString, nil)
	r.Header.Add("Content-Type", contentType)
	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	result := responseToBytes(resp)
	return result, nil
}

func responseToBytes(r *http.Response) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	rBytes := buf.Bytes()

	return rBytes
}
