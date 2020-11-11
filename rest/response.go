package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Response represents an HTTP response
type Response struct {
	Status      int
	Headers     []Header
	Body        string
	DecodedBody interface{}
}

func (response *Response) setDecodedBody() {
	var decodedBody map[string]interface{}

	for _, header := range response.Headers {
		if header.Name != "Content-Type" {
			continue
		}

		if strings.Contains(header.Value, "application/json") {
			json.Unmarshal([]byte(response.Body), &decodedBody)
		}
	}

	response.DecodedBody = decodedBody
}

func (response *Response) setHeaders(httpResponse *http.Response) {
	httpHeaders := []Header{}

	for name, value := range httpResponse.Header {
		httpHeader := Header{
			Name:  name,
			Value: strings.Join(value, ";"),
		}

		httpHeaders = append(httpHeaders, httpHeader)
	}

	response.Headers = httpHeaders
}

func (response *Response) setBody(httpResponse *http.Response) {
	defer httpResponse.Body.Close()
	body, _ := ioutil.ReadAll(httpResponse.Body)

	response.Body = string(body)
}
