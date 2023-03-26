package main

import (
	"fmt"
	"os"

	server "github.com/WildEgor/checker/pkg"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Set logging settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	server, _ := server.NewServer()
	log.Fatal(server.Listen(fmt.Sprintf(":%v", "8888")))
}
