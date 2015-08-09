package node

import (
    "github.com/lonnc/golang-nw"
)

func Run() {

    // Create a link back to node-webkit using the environment variable
    // populated by golang-nw's node-webkit code
    nodeWebkit, err := nw.New()
    if err != nil {
        panic(err)
    }

    // Pick a random localhost port, start listening for http requests using default handler
    // and send a message back to node-webkit to redirect
    if err := nodeWebkit.ListenAndServe(nil); err != nil {
        panic(err)
    }
}