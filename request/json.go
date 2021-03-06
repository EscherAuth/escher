package request

import (
	"encoding/json"
)

type jsonMapper struct {
	URL     string      `json:"url"`
	Body    string      `json:"body"`
	Method  string      `json:"method"`
	Expires int         `json:"expires"`
	Headers [][2]string `json:"headers"`
}

func ParseJSON(data []byte) (*Request, error) {
	r := &Request{}
	err := mapJSONContentToRequest(r, data)
	return r, err
}

func (r *Request) UnmarshalJSON(data []byte) error {
	return mapJSONContentToRequest(r, data)
}

func (r *Request) MarshalJSON() ([]byte, error) {
	j := &jsonMapper{
		URL:     r.url,
		Body:    r.body,
		Method:  r.method,
		Expires: r.expires,
		Headers: r.headers,
	}
	return json.Marshal(j)
}

func mapJSONContentToRequest(r *Request, data []byte) error {
	var j jsonMapper
	err := json.Unmarshal(data, &j)

	if err != nil {
		return err
	}

	r.url = j.URL
	r.body = j.Body
	r.method = j.Method
	r.expires = j.Expires
	r.headers = j.Headers

	// uri, err := url.Parse(j.URL)
	// if err != nil {
	// 	return err
	// }
	// r.UniversalResourceLocator = uri

	return nil
}
