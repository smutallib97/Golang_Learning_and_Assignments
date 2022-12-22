
package main

import "fmt"

func main() {
    // arr := [6]int{10, 20, 30, 40, 50, 60}
    // sliceCreated := arr[2:4]
    
    // fmt.Printf("Slice created = %v\n", sliceCreated)
    
    // fmt.Printf("From array = %v\n", arr)
  
    // fmt.Printf("length = %d, capacity = %d \n", len(sliceCreated), cap(sliceCreated) )
    
    
    strngArr := [...]string{"a","b","c","d","e","f","g"}
    sl1 :=  strngArr[1:3]
    sl2 :=  strngArr[2:5]
    fmt.Println(sl1, sl2)
}
