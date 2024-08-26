package main

import "fmt"

func main()  {
	var values []interface{}

values = append(values, "A string")
values = append(values, 123)
values = append(values, 45.67)

for _, v := range values {
    fmt.Println(v)
}

}