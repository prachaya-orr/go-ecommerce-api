package utils

import (
	"encoding/json"
	"fmt"
)

func Debug(data any) {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(bytes))
}
