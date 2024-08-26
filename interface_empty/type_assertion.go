package main

import "fmt"

func main(){
	var value interface{} = "A string"

str, ok := value.(string)
if ok {
    fmt.Println("The value is a string:", str)
} else {
    fmt.Println("The value is not a string")
}


}