package cpptest

import (
    "gopkg.in/qml.v1"
    "fmt"
)

func main() {
    engine := qml.NewEngine()
    object := NewTestType(engine)
    fmt.Print(object)
}