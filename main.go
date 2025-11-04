package main

import (
	"fmt"
	"os"
	"github.com/shivgitcode/qv/cmd"
)



func main(){
	if len(os.Args)<2 && (os.Args[0]=="./qv" || os.Args[0]=="qv"){
		cmd.Welcome()
		return
	}
	if (len(os.Args)<2){
		fmt.Println("Availaible commands are init , set , get , delete")
		return
	}

	

	command:=os.Args[1]
	args:=os.Args[2:]

	switch command{
	case "init":
		cmd.Init(args)
	case "set":
		cmd.Save(args)
	case "get":
		cmd.Get(args)
	case "delete":
		cmd.Delete(args)
	case "update":
		cmd.Update(args)
	case "list":
		cmd.List(args)
	default:
		fmt.Println("command do not exist here is the list of available command\ninit\nsave\ndelete\nupdate")
	}

	
}