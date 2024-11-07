package main

import "fmt"

func main() {
    counter := createCounter()
    fmt.Println(counter()) // Output: 1
    fmt.Println(counter()) // Output: 2
    fmt.Println(counter()) // Output: 3
}

// Closure yang mengembalikan fungsi yang akan menghitung setiap kali dipanggil
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
