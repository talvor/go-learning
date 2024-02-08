package main

import (
    "fmt"

    "talvor.github.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Phillip")
    fmt.Println(message)
}