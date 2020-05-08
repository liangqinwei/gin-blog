package utils

import "encoding/json"

func StrMap(data map[string]interface{}) string{
	b,err:=json.Marshal(data)
	if err!=nil{
		return ""
	}
	return string(b)
}

