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
    supportedCurrencies(client)
    currentPrice(client)
    currentPriceForCurrency(client)
    historical(client)
    historicalForYesterday(client)
    historicalForDates(client)
}

func version() {
    fmt.Println("BPI Interface Version:", bpi.Version)
    fmt.Println("BPI Interface Author:", bpi.Author)
}

func supportedCurrencies(client *bpi.ApiClient) {
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

func currentPrice(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("Get current BPI.")
    fmt.Println("=======================================")

    r, err := client.CurrentPrice()
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("Time / Updated:", r.Time.Updated)
    fmt.Println("Time / UpdatedISO:", r.Time.UpdatedISO)
    fmt.Println("Time / UpdatedUK:", r.Time.UpdatedUK)
    fmt.Println("BPI:")

    for k := range r.BPI {
        fmt.Println("  " + k + ":")
        fmt.Println("    Code", r.BPI[k].Code)
        fmt.Println("    Symbol", r.BPI[k].Symbol)
        fmt.Println("    Rate", r.BPI[k].Rate)
        fmt.Println("    Description", r.BPI[k].Description)
        fmt.Println("    RateFloat", r.BPI[k].RateFloat)
    }
}

func currentPriceForCurrency(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("Get current BPI for symbol UYU.")
    fmt.Println("=======================================")

    r, err := client.CurrentPriceForCurrency("UYU")
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("Time / Updated:", r.Time.Updated)
    fmt.Println("Time / UpdatedISO:", r.Time.UpdatedISO)
    fmt.Println("Time / UpdatedUK:", r.Time.UpdatedUK)
    fmt.Println("BPI:")

    for k := range r.BPI {
        fmt.Println("  " + k + ":")
        fmt.Println("    Code", r.BPI[k].Code)
        fmt.Println("    Symbol", r.BPI[k].Symbol)
        fmt.Println("    Rate", r.BPI[k].Rate)
        fmt.Println("    Description", r.BPI[k].Description)
        fmt.Println("    RateFloat", r.BPI[k].RateFloat)
    }
}

func historical(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("Get historical BPI (last 31 days).")
    fmt.Println("=======================================")

    r, err := client.Historical()
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("Time / Updated:", r.Time.Updated)
    fmt.Println("Time / UpdatedISO:", r.Time.UpdatedISO)
    fmt.Println("BPI:")

    for k := range r.BPI {
        fmt.Println("  " + k, r.BPI[k])
    }
}

func historicalForYesterday(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("Get historical BPI for yesterday.")
    fmt.Println("=======================================")

    r, err := client.HistoricalForYesterday()
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("Time / Updated:", r.Time.Updated)
    fmt.Println("Time / UpdatedISO:", r.Time.UpdatedISO)
    fmt.Println("BPI:")

    for k := range r.BPI {
        fmt.Println("  " + k, r.BPI[k])
    }
}

func historicalForDates(client *bpi.ApiClient) {
    fmt.Println()
    fmt.Println("=======================================")
    fmt.Println("Get historical BPI from 2014-01-01 to")
    fmt.Println("2014-02-28.")
    fmt.Println("=======================================")

    r, err := client.HistoricalForDates("2014-01-01", "2014-02-28")
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }

    fmt.Println("Disclaimer:", r.Disclaimer)
    fmt.Println("Time / Updated:", r.Time.Updated)
    fmt.Println("Time / UpdatedISO:", r.Time.UpdatedISO)
    fmt.Println("BPI:")

    for k := range r.BPI {
        fmt.Println("  " + k, r.BPI[k])
    }
}
