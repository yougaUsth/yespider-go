package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"yespider-go/settings"
	"yespider-go/www/controller"
)


func init(){
	// Init config
	settings.Setup()
}


func main() {
	gin.SetMode(settings.ServerSettings.RunMode)
	// Init handler
	routersHandler := controller.InitRouter()

	HttpPort := fmt.Sprintf(":%v", settings.ServerSettings.HttpPort)
	server := &http.Server{
		Addr: HttpPort,
		Handler:routersHandler,
		ReadTimeout: settings.ServerSettings.ReadTimeout * time.Second,
		WriteTimeout: settings.ServerSettings.WriteTimeout * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil{
		log.Fatalf("Server error : %v", err)
	}
	fmt.Print("Server listen to ", HttpPort)
}
