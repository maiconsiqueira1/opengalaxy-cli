package rest

// Header http headers
type Header struct {
	Name  string
	Value string
}

// Cookie http cookies
type Cookie struct {
	Name   string
	Value  string
	Secure bool
}

// Start create a new Request
func Start(target string, method ...string) (Request, error) {
	return Request{Method: "GET", Target: target}, nil
}
