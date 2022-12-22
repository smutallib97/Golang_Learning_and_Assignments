
package main

import "fmt"


type Test struct {
    x string
    y int
} 
func main() {
  b := Test{"x", -1}
  arr := [...]Test{
        Test{"a", 10}, 
        Test{"b", 2},
        Test{"c", 3},
    }
 for _, z := range arr {
    if z.y > b.y {
      b = z
    }
 }
 fmt.Println(b.x)
}