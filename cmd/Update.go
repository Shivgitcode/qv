package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/shivgitcode/qv/internal"
	"github.com/shivgitcode/qv/storage"
)



func Update(values []string){
	updateCmd:=flag.NewFlagSet("update",flag.ExitOnError)
	nameFlag:=updateCmd.String("name","","finds the variable by name to update")
	varFlag:=updateCmd.String("var","","the new variables which will replace the old")

	err:=updateCmd.Parse(values)
	check(err)

	var4:=make(map[string]string)

	if _,err:=os.Stat(internal.GetFilePath());err==nil{
		f:=storage.ReadFile()
		err:=json.Unmarshal(f,&var4)
		check(err)
		if var4[*nameFlag]==""{
			fmt.Println("it seems you want to create new variable use set command")
			return
		}
		var4[*nameFlag]=*varFlag
		data,err:=json.Marshal(var4)
		check(err)
		storage.WriteFile(data)
		fmt.Println("variable updated !")
	}else{
		fmt.Println("first initialize the file please !")
	}


}