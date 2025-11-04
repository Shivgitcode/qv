package cmd

import (
    "fmt"
    figure "github.com/common-nighthawk/go-figure"
)

// Welcome prints a short help banner and quick initialization steps
// for the QuickVariables (qv) CLI, using go-figure for the banner.
func Welcome() {
    // ASCII banner using go-figure (nighthawk repo)
    fig := figure.NewFigure("qv", "", true)
    fig.Print()

    fmt.Println("QuickVariables (qv)")
    fmt.Println("A tiny CLI to store and fetch named variables.")
    fmt.Println()
    fmt.Println("Quick init:")
    fmt.Println("  1) Create .env with QV_PATH=~/.config/quickvariable.json")
    fmt.Println("  2) Run: qv init")
    fmt.Println()
    fmt.Println("Common commands:")
    fmt.Println("  - qv set --name <key> --var <value>")
    fmt.Println("  - qv get --name <key>")
    fmt.Println("  - qv update --name <key> --var <value>")
    fmt.Println("  - qv delete --name <key>")
    fmt.Println("  - qv list")
}
