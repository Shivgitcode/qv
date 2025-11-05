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
	data,err:=os.ReadFile(internal.Path)
	check(err)
	return data

}

func WriteFile(f []byte){
	err:=os.WriteFile(internal.Path,f,0644)
	check(err)
}