package main

import "fmt"

func main() {
    // Function literal yang dicetak langsung
    func() {
        fmt.Println("Hello from a function literal!")
    }()

    // Function literal disimpan dalam variabel
    add := func(a, b int) int {
        return a + b
    }

    fmt.Println("Hasil dari add(3, 4):", add(3, 4)) // Output: 7
}
