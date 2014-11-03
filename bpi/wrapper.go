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

type BPI struct {
    Time        BPITime                 `json:"time"`
    Disclaimer  string                  `json:"disclaimer"`
    BPI         map[string]BPICurrency  `json:"bpi"`
}

type BPITime struct {
    Updated     string      `json:"updated"`
    UpdatedISO  string      `json:"updatedISO"`
    UpdatedUK   string      `json:"updateduk"`
}

type BPICurrency struct {
    Code        string      `json:"code"`
    Symbol      string      `json:"symbol"`
    Rate        json.Number `json:"rate"`
    Description string      `json:"description"`
    RateFloat   json.Number `json:"rate_float"`
}

type HistoricalBPI struct {
    Time        HistoricalBPITime       `json:"time"`
    Disclaimer  string                  `json:"disclaimer"`
    BPI         map[string]json.Number  `json:"bpi"`
}

type HistoricalBPITime struct {
    Updated     string      `json:"updated"`
    UpdatedISO  string      `json:"updatedISO"`
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

func (c *ApiClient) CurrentPrice() (*BPI, error) {
    return c.current("currentprice")
}

func (c *ApiClient) CurrentPriceForCurrency(symbol string) (*BPI, error) {
    return c.current("currentprice/" + symbol)
}

func (c *ApiClient) current(endpoint string) (*BPI, error) {
    data, err := c.apiCall(endpoint, nil)
    if err != nil { return nil, err }

    var b BPI
    err = json.Unmarshal(data, &b)
    if err != nil { return nil, err }

    return &b, nil
}

func (c *ApiClient) Historical() (*HistoricalBPI, error) {
    return c.HistoricalWithOptions(false, "", "")
}

func (c *ApiClient) HistoricalForYesterday() (*HistoricalBPI, error) {
    return c.HistoricalWithOptions(true, "", "")
}

func (c *ApiClient) HistoricalForDates(start string, end string) (*HistoricalBPI, error) {
    return c.HistoricalWithOptions(false, start, end)
}

func (c *ApiClient) HistoricalWithOptions(yesterday bool, start string, end string) (*HistoricalBPI, error) {
    args := make(map[string]string)

    if yesterday {
        args["for"] = "yesterday"
    } else if start != "" && end != "" {
        args["start"] = start
        args["end"] = end
    }

    data, err := c.apiCall("historical/close", args)
    if err != nil { return nil, err }

    var hb HistoricalBPI
    err = json.Unmarshal(data, &hb)
    if err != nil { return nil, err }

    return &hb, nil
}

func (c *ApiClient) apiCall(endpoint string, args map[string]string) ([]byte, error) {
    // Build URL
    argstring := ""
    for k := range args {
        argstring = fmt.Sprintf("%v&%v=%v", argstring, k, args[k])
    }
    if argstring != "" { argstring = "?" + argstring[1:len(argstring)] }

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
