
package main

import "fmt"


func main() {
 stringOfOneChar := 'a'
    
    s_byte := byte(stringOfOneChar)

    fmt.Printf("for char %c -  byte value is %d  and its type is %T \n",
                stringOfOneChar,s_byte,s_byte)
}