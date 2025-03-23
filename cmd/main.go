package main

import (
	"eventapp/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// create logfile here
	logFile, err := os.OpenFile("../log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Error Opening log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Log File intilization successfully done.")

	// initilize gin app
	router := gin.Default()
	routes.RegisterRouters(router)
	err = router.Run(":8000")
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
