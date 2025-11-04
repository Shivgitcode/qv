package storage

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/shivgitcode/qv/internal"
)

func check(e error){
	if e!=nil{
		panic(e)
	}
}


func ReadFile()[]byte{
	err:=godotenv.Load()
	check(err)
	data,err:=os.ReadFile(internal.Path)
	check(err)
	return data

}

func WriteFile(f []byte){
	err:=godotenv.Load()
	check(err)
	err=os.WriteFile(internal.Path,f,0644)
	check(err)
}