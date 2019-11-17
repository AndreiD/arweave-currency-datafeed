package utils

import (
	"arweave-datafeed/utils/log"
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Close error checking for defer close
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// GetRequest executes a generic GET request
func GetRequest(url string) ([]byte, int, error) {

	client := http.Client{Timeout: 180 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, -1, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	reqWithDeadline := req.WithContext(ctx)
	response, err := client.Do(reqWithDeadline)
	if err != nil {
		return nil, -1, err
	}

	data, err := ioutil.ReadAll(response.Body)

	return data, response.StatusCode, err

}

// EncodeToBase64 encodes a byte array to base64 raw url encoding
func EncodeToBase64(toEncode []byte) string {
	return base64.RawURLEncoding.EncodeToString(toEncode)
}

// DecodeString decodes from base64 raw url encoding to byte array
func DecodeString(toDecode string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(toDecode)
}
