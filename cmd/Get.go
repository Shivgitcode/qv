package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shivgitcode/qv/storage"

)



func Get(values []string){
	getCmd:=flag.NewFlagSet("get",flag.ExitOnError)
	nameFlag:=getCmd.String("name","","name for getting the flag")
	err:=getCmd.Parse(values)
	check(err)

	var var1 map[string]string

	dat:=storage.ReadFile()
	check(err)
	err=json.Unmarshal(dat,&var1)



	check(err)
	fmt.Println(var1[*nameFlag])

	

}