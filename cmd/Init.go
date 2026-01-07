package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/shivgitcode/qv/v2/internal"
)
var path string="/Users/admin/.config/quickvariable.json"

func writeFile(filepath string){
	if _,err:=os.Stat(filepath);err==nil && !internal.IsValidjson(filepath){
		fmt.Println("File already initialised at",path)
		return
	}

	vars:=make(map[string]string)
	
	f,err:=os.Create(filepath)
	if err!=nil{
		panic(err)
	}
	data,_:=json.Marshal(vars)
	f.Write(data)
	defer f.Close()

	fmt.Println("qv initialized at",filepath)
}


func Init(values []string){
	initCmd:=flag.NewFlagSet("init",flag.ExitOnError)

	if err:=initCmd.Parse(values); err!=nil{
		fmt.Println("atleast one argument required")
		return
	}



	writeFile(internal.GetFilePath())

	
}