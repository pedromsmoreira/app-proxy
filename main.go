package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pedromsmoreira/app-proxy/cmd"
	"github.com/pedromsmoreira/app-proxy/proxies"
	"gopkg.in/yaml.v3"
)

func main() {
	// read file - done
	// convert file to []Proxies - done
	yamlFile, err := ioutil.ReadFile("data.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	fmt.Println("YAML File Content")
	fmt.Println(string(yamlFile))
	proxies := &proxies.Proxies{}

	err = yaml.Unmarshal(yamlFile, proxies)

	if err != nil {
		log.Printf("Unmarshal err #%v", err)
	}

	// TODO: Convert proxies into a map for faster and easier access
	// Option 2: use redis to store these structures
	code := mainWithReturnCode(proxies)
	if code != 0 {
		os.Exit(code)
	}

	// define endpoint
	// if endpoint exists - return result
	// if not exists - 404 - error endpoint not set
	// start server
}

func mainWithReturnCode(prxies *proxies.Proxies) int {
	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "3000"
	}

	address, exists := os.LookupEnv("ADDRESS")

	if !exists {
		address = "localhost"
	}

	sc := make(chan os.Signal, 1)

	server := cmd.NewServer()

	err := server.Start(address, port, prxies)

	if err != nil {
		return 1
	}

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = server.Shutdown()

	if err == nil {
		return 0
	}

	return 1
}
