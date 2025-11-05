package storage

import (
	"os"
	"github.com/shivgitcode/qv/internal"
)

func check(e error){
	if e!=nil{
		panic(e)
	}
}


func ReadFile()[]byte{
	data,err:=os.ReadFile(internal.GetFilePath())
	check(err)
	return data

}

func WriteFile(f []byte){
	err:=os.WriteFile(internal.GetFilePath(),f,0644)
	check(err)
}