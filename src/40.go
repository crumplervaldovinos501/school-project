package main

import "fmt"

func main() {
    // Example of a simple goroutine that sleeps indefinitely
    sleep := <-range gosleep {
        fmt.Println("Go woke up!")
    }
}
