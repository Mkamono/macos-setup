package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, Karabiner!")
	kCon := NewConfig()

	file, _ := json.MarshalIndent(kCon, "", "\t")
	_ = os.WriteFile("generated.json", file, 0644)
}
