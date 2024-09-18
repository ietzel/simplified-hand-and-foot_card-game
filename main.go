package main

import (
    "os"
    "gioui.org/app"
    "gioui.org/op"
    "gioui.org/unit"
    "gioui.org/widget"
    "gioui.org/widget/material"
)

func main() {
    go func() {
        w := new(app.Window)
        for {
            w.Event()
        }
    }()
    app.Main()
}
