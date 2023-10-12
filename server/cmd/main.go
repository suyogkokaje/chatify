package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main(){
	dbConn,err:=db.NewDatabase()
	if err!=nil{
		log.Fatalf("Failed to initialize database connection: %s",err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHanlder(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}