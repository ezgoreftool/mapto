package mapto

import (
	"encoding/json"
	"fmt"
)

func init() {
	fmt.Println(".")
}

func MapToJSON(data map[string]interface{}) string {
	jsonPayload, _ := json.Marshal(data)
	return string(jsonPayload)
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}