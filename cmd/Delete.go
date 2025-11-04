package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"github.com/shivgitcode/qv/internal"
	"github.com/shivgitcode/qv/storage"
)


func Delete(values []string){
	delCmd:=flag.NewFlagSet("delete",flag.ExitOnError)
	nameFlag:=delCmd.String("name","","to delete an existing variable")

	err:=delCmd.Parse(values)
	check(err)

	var3:=make(map[string]string)

	f:=storage.ReadFile()
	err=json.Unmarshal(f,&var3)
	delete(var3,*nameFlag)
	data,err:=json.Marshal(var3)
	check(err)
	err=os.WriteFile(internal.Path,data,0644)
	check(err)
	fmt.Println("variable removed !")
}