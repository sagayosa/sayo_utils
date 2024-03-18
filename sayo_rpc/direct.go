package sayorpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Delete(URL string, data interface{}) (code int, body []byte, err error) {
	return httpRequest("DELETE", URL, data)
}

func Put(URL string, data interface{}) (code int, body []byte, err error) {
	return httpRequest("PUT", URL, data)
}

func Post(URL string, data interface{}) (code int, body []byte, err error) {
	return httpRequest("POST", URL, data)
}

func Get(URL string, data map[string]interface{}) (code int, body []byte, err error) {
	if data == nil {
		data = make(map[string]interface{})
	}
	params := url.Values{}
	for k, v := range data {
		params.Add(k, fmt.Sprintf("%v", v))
	}

	reqURL, err := url.Parse(URL)
	if err != nil {
		return
	}
	reqURL.RawQuery = params.Encode()

	resp, err := http.Get(reqURL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return resp.StatusCode, body, nil
}

func httpRequest(way string, URL string, data interface{}) (code int, body []byte, err error) {
	if data == nil {
		data = struct{}{}
	}
	bts, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest(way, URL, bytes.NewBuffer(bts))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return resp.StatusCode, body, nil
}
