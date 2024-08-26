package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{"name": "John", "age": 30, "city": "New York"}`

	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result["name"]) // John
	fmt.Println(result["age"])  // 30
	fmt.Println(result["city"]) // New York
}
