package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shivgitcode/qv/storage"
	"github.com/AlecAivazis/survey/v2"
)


func List(values []string){
	listCmd:=flag.NewFlagSet("list",flag.ExitOnError)
	visualFlag:=listCmd.Bool("visual",false,"use to see a option menu of present variables and select from those")
	f:=storage.ReadFile()
	
	err:=listCmd.Parse(values)
	check(err)

	var5:=make(map[string]string)
	var arr1 []string
	

	if *visualFlag{
		var option string
		err:=json.Unmarshal(f,&var5)
		check(err)
		
		for k,_:=range var5{
			arr1=append(arr1, k)
		}
		prompt:=&survey.Select{
			Message: "Select from the list of Present Variables",
			Options: arr1,
		}
		survey.AskOne(prompt,&option)
		fmt.Println(var5[option])
		return
	}
	check(err)
	err=json.Unmarshal(f,&var5)
	check(err)
	for k,_:=range var5{
		fmt.Println(k)
	}
	

}