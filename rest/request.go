package rest

import (
	"bytes"
	"net/http"
)

// Request represents an HTTP request
type Request struct {
	Method  string
	Target  string
	Body    string
	Cookies []Cookie
	Headers []Header
}

// AddHeader add a new header in HTTP request
func (request *Request) AddHeader(name string, value string) {
	request.Headers =
		append(request.Headers, Header{
			Name:  name,
			Value: value,
		})
}

// AddCookie add a new cookie in HTTP request
func (request *Request) AddCookie(name string, value string) {
	request.Cookies =
		append(request.Cookies, Cookie{
			Name:  name,
			Value: value,
		})
}

// SetRequestBody set request body
func (request *Request) SetRequestBody(body string) {
	request.Body = body
}

// Send send a request
func (request *Request) Send() (Response, error) {
	response := Response{}

	// prepare and send http request
	httpRequest := prepareNetHTTPRequest(request)
	httpResponse, err := sendNetHTTPRequest(httpRequest)

	// set all response entities
	response.setBody(httpResponse)
	response.setHeaders(httpResponse)
	response.setDecodedBody()

	return response, err
}

func prepareNetHTTPRequest(request *Request) *http.Request {
	requestBody := bytes.NewBufferString(request.Body)
	httpRequest, _ := http.NewRequest(
		request.Method,
		request.Target,
		requestBody,
	)

	for _, cookie := range request.Cookies {
		httpRequest.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}

	for _, header := range request.Headers {
		httpRequest.Header.Add(header.Name, header.Value)
	}

	return httpRequest
}

func sendNetHTTPRequest(httpRequest *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}

	return httpClient.Do(httpRequest)
}
