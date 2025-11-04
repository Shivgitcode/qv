package cmd

import (
	// "bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/shivgitcode/qv/internal"
	"github.com/shivgitcode/qv/storage"

	"github.com/joho/godotenv"
)

func check(e error){
	if e!=nil{
		panic(e)
	}

}

func Save(values []string){
	err:=godotenv.Load()
	check(err)
	setCmd:=flag.NewFlagSet("set",flag.ExitOnError)
	nameFlag:=setCmd.String("name","","used to give name to variable")
	varFlag:=setCmd.String("var","","used to give name to variable")

	var2:=make(map[string]string)

	err=setCmd.Parse(values)
	check(err)


	if _,err:=os.Stat(internal.Path);err==nil{
		data:=storage.ReadFile()
		check(err)
		err=json.Unmarshal(data,&var2)
		check(err)
		if var2[*nameFlag]==""{
			var2[*nameFlag]=*varFlag
		}else{
			fmt.Println("this variable already exist , do you want to update it")
			return

		}
		f,_:=json.Marshal(var2)
		storage.WriteFile(f)
		check(err)
		fmt.Println("variable saved!")
		return
	}else{
		fmt.Println("initialize the file please use qv init")
	}

	
}