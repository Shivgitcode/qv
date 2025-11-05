package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/shivgitcode/qv/internal"
)
var path string="/Users/admin/.config/quickvariable.json"

func writeFile(filepath string){
	if _,err:=os.Stat(filepath);err==nil{
		fmt.Println("File already initialised at",path)
		return
	}
	
	f,err:=os.Create(filepath)
	if err!=nil{
		panic(err)
	}
	defer f.Close()

	fmt.Println("qv initialized at",filepath)
}


func Init(values []string){
	initCmd:=flag.NewFlagSet("init",flag.ExitOnError)

	if err:=initCmd.Parse(values); err!=nil{
		fmt.Println("atleast one argument required")
		return
	}


	writeFile(internal.Path)

	
}