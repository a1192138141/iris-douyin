package utils

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"ims/models"
)

func MapToSturct(rows interface{}, strus interface{}) error {
	if err := mapstructure.Decode(rows, &strus); err != nil {
		return err
	}

	fmt.Print("=========")
	fmt.Print(strus)
	return nil
}
func MapToStruct(rows interface{}, strus models.User) (res interface{}, err error) {
	if err := mapstructure.Decode(rows, strus); err != nil {
		return strus, err
	}

	return strus, nil
}

func BytesToStruct(bytes []byte, row interface{}) bool {
	if err := json.Unmarshal(bytes, &row); err != nil {
		return false
	}
	return true
}
