package internal

import (
	"encoding/json"
	"os"
)

var Path string="/Users/admin/.config/quickvariable.json"

func IsValidjson(filepath string) bool{
	data,_:=os.ReadFile(filepath)
	var temp map[string]string
	err:=json.Unmarshal(data,temp)

	if err==nil{
		return true;

	}
	return false
}