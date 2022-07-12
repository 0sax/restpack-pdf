package restpack_pdf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://restpack.io/api"

type client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func initClient(apiKey string) *client {
	return &client{
		baseURL:    baseUrl,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}
func (c *client) request(payload interface{}, endpoint string) (resp *http.Response, err error) {
	//marshal payload to json
	var b []byte
	b, err = json.Marshal(payload)
	if err != nil {
		return
	}

	//build request
	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%v%v", c.baseURL, endpoint),
		bytes.NewReader(b))
	if err != nil {
		return
	}
	req.Header.Set("x-access-token", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "github.com/0sax/restpack-pdf")

	//make request
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		bb, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("request to Restpack returned a %v error %+v\n", resp.Status, string(bb))
	}

	return
}
