package main

import (
	"log"
	"server/db"
)

func main(){
	_,err:=db.NewDatabase()
	if err!=nil{
		log.Fatalf("Failed to initialize database connection: %s",err)
	}
}