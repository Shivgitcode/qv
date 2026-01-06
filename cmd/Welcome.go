package cmd

import (
    "fmt"
    figure "github.com/common-nighthawk/go-figure"
)


func Welcome() {
    fig := figure.NewFigure("qv", "", true)
    fig.Print()

    fmt.Println("QuickVariables (qv)")
    fmt.Println("A tiny CLI to store and fetch named variables.")
    fmt.Println()
    fmt.Println("Quick init:")
    fmt.Println("  2) Run: qv init")
    fmt.Println()
    fmt.Println("Common commands:")
    fmt.Println("  - qv set --name <key> --var <value>")
    fmt.Println("  - qv get --name <key>")
    fmt.Println("  - qv update --name <key> --var <value>")
    fmt.Println("  - qv delete --name <key>")
    fmt.Println("  - qv list ")
}
