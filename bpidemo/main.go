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
    current(client)
    currentForCurrency(client)
}

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

func current(client *bpi.ApiClient) {
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

func currentForCurrency(client *bpi.ApiClient) {
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
