
package main

import "fmt"


func main() {
 testString := "Golang"
    
    runeSlice := []rune(testString)
    
    fmt.Println(runeSlice)
    
    byteSlice := []byte(testString)
    
    fmt.Println(byteSlice)
}