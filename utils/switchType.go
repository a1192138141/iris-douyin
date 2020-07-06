package utils

import (
	"encoding/json"
	"github.com/goinggo/mapstructure"
)


func MapToSturct(rows interface{} ,strus interface{}) error {
	if err :=mapstructure.Decode(rows,&strus) ;err != nil{
		return err
	}
	return  nil
}

func BytesToStruct(bytes []byte , row interface{})  bool {
	if err  := json.Unmarshal(bytes,&row) ;err !=nil{
		return false
	}
	return true
}