package main

import (
    "fmt"
)

func getOdd(numbers []int, oddChan chan<- int) {
    for _, num := range numbers {
        if num%2 != 0 {
            oddChan <- num
        }
    }
    close(oddChan)
}

func getEven(numbers []int, evenChan chan<- int) {
    for _, num := range numbers {
        if num%2 == 0 {
            evenChan <- num
        }
    }
    close(evenChan)
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    oddChan := make(chan int)
    evenChan := make(chan int)

    go getOdd(numbers, oddChan)
    go getEven(numbers, evenChan)

    fmt.Println("Odd numbers:")
    for num := range oddChan {
        fmt.Print(num, " ")
    }

    fmt.Println("\nEven numbers:")
    for num := range evenChan {
        fmt.Print(num, " ")
    }
}