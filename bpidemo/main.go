package main

import (
    "github.com/mvillalba/go-coindesk/bpi"
    "fmt"
)

func main() {
    // Init ApiClient
    client := bpi.New()

    // Fun stuff
    version()
    currencies(client)
/*    latest(client)
    historical(client)
    latestWithOptions(client)
    historicalWithOptions(client)
    timeSeries(client)
    timeSeriesWithOptions(client)
    convert(client)
*/}

func version() {
    fmt.Println("BPI Interface Version:", bpi.Version)
    fmt.Println("BPI Interface Author:", bpi.Author)
}

func currencies(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("List all supported currencies.")
    fmt.Println("=======================================")

    curr, err := client.SupportedCurrencies()
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    for k := range curr {
        fmt.Println(curr[k].Currency, curr[k].Country)
    }
}
/*
func latest(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("List latest exchange rates.")
    fmt.Println("=======================================")

    r, err := client.Latest()
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("License:", r.License)
    fmt.Println("Timestamp:", r.Timestamp)
    fmt.Println("Base:", r.Base)
    fmt.Println("")

    for k := range r.Rates {
        fmt.Println("USD/" + k, r.Rates[k])
    }
}

func latestWithOptions(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("List latest exchange rates for base")
    fmt.Println("smybol BTC and quote symbols EUR, NZD,")
    fmt.Println("USD, ARS, and JPY.")
    fmt.Println("=======================================")

    syms := []string{"EUR", "NZD", "USD", "ARS", "JPY"}
    r, err := client.LatestWithOptions("BTC", syms)
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("License:", r.License)
    fmt.Println("Timestamp:", r.Timestamp)
    fmt.Println("Base:", r.Base)
    fmt.Println("")

    for k := range r.Rates {
        fmt.Println("USD/" + k, r.Rates[k])
    }
}

func historical(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("List latest exchange rates as they were")
    fmt.Println("on 2014-01-01.")
    fmt.Println("=======================================")

    r, err := client.Historical("2014-01-01")
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    for k := range r.Rates {
        fmt.Println("USD/" + k, r.Rates[k])
    }
}

func historicalWithOptions(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("List latest exchange rates as they were")
    fmt.Println("on 2014-01-01 for base symbol BTC and")
    fmt.Println("quote symbols EUR, NZD, USD, ARS, and")
    fmt.Println("JPY.")
    fmt.Println("=======================================")

    syms := []string{"EUR", "NZD", "USD", "ARS", "JPY"}
    r, err := client.HistoricalWithOptions("2014-01-01", "BTC", syms)
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    for k := range r.Rates {
        fmt.Println("BTC/" + k, r.Rates[k])
    }
}

func timeSeries(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("List historical exchange rates in bulk")
    fmt.Println("for the first 7 days of 2014.")
    fmt.Println("=======================================")

    r, err := client.TimeSeries("2014-01-01", "2014-01-07")
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("License:", r.License)
    fmt.Println("Start:", r.StartDate)
    fmt.Println("End:", r.EndDate)
    fmt.Println("Base:", r.Base)
    fmt.Println("Rates:")

    for k := range r.Rates {
        fmt.Println("  " + k + ":")
        for kk := range r.Rates[k] {
            fmt.Println("    " + kk, r.Rates[k][kk])
        }
    }
}

func timeSeriesWithOptions(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("List historical exchange rates in bulk")
    fmt.Println("for the first 7 days of 2014 for base")
    fmt.Println("symbol BTC and quote symbols AUD, THB,")
    fmt.Println("and SEK.")
    fmt.Println("=======================================")

    syms := []string{"AUD", "THB", "SEK"}
    r, err := client.TimeSeriesWithOptions("2014-01-01", "2014-01-07", "BTC", syms)
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("License:", r.License)
    fmt.Println("Start:", r.StartDate)
    fmt.Println("End:", r.EndDate)
    fmt.Println("Base:", r.Base)
    fmt.Println("Rates:")

    for k := range r.Rates {
        fmt.Println("  " + k + ":")
        for kk := range r.Rates[k] {
            fmt.Println("    " + kk, r.Rates[k][kk])
        }
    }
}

func convert(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("Convert 10.123456789 BTC to UYU.")
    fmt.Println("=======================================")

    c, err := client.Convert("10.123456789", "BTC", "UYU")
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", c.Disclaimer)
    fmt.Println("License:", c.License)
    fmt.Println("Request / Query:", c.Request.Query)
    fmt.Println("Request / Amount:", c.Request.Amount)
    fmt.Println("Request / From:", c.Request.From)
    fmt.Println("Request / To:", c.Request.To)
    fmt.Println("Meta / Timestamp:", c.Meta.Timestamp)
    fmt.Println("Meta / Rate:", c.Meta.Rate)
    fmt.Println("Response:", c.Response)
}*/
