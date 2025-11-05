package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func GetFilePath() string{
	home,_:=os.UserHomeDir()
	filePathName:=filepath.Join(home,".config","quickvariable.json")
	return filePathName
}


func IsValidjson(filepath string) bool{
	data,_:=os.ReadFile(filepath)
	var temp map[string]string
	err:=json.Unmarshal(data,temp)


	if err==nil && len(temp)>0{
		return true;

	}
	return false
}