package bpi

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "errors"
    "fmt"
)

var (
    ProtoHttp = "http"
    ProtoHttps = "https"
    ApiUrl = "api.coindesk.com/v1/bpi"
)

type ApiClient struct {
    proto       string
    url         string
}

type Currency struct {
    Currency    string      `json:"currency"`
    Country     string      `json:"country"`
}

func New() *ApiClient {
    return NewWithOptions(ProtoHttps, ApiUrl)
}

func NewWithOptions(proto string, url string) *ApiClient {
    return &ApiClient{proto: proto, url: url}
}

func (c *ApiClient) SupportedCurrencies() ([]Currency, error) {
    data, err := c.apiCall("supported-currencies", nil)
    if err != nil { return nil, err }

    var sc []Currency
    err = json.Unmarshal(data, &sc)
    if err != nil { return nil, err }

    return sc, nil
}

func (c *ApiClient) apiCall(endpoint string, args map[string]string) ([]byte, error) {
    // Build URL
    argstring := ""
    for k := range args {
        argstring = fmt.Sprintf("%v&%v=%v", argstring, k, args[k])
    }
    if argstring != "" { argstring = argstring[1:len(argstring)] }

    url := fmt.Sprintf("%v://%v/%v.json%v", c.proto, c.url, endpoint, argstring)

    // Make request
    resp, err := http.Get(url)
    if err != nil { return nil, err }

    // Retrieve raw JSON response
    var body []byte
    body, err = ioutil.ReadAll(resp.Body)
    if err != nil { return nil, err }
    defer resp.Body.Close()

    // Process API-level error conditions
    if resp.StatusCode != 200 {
        return nil, errors.New(string(body))
    }

    return body, nil
}
