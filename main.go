package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Proxies struct {
	Proxies []Proxy `yaml:"proxies"`
}

type Proxy struct {
	Endpoint    string            `yaml:"endpoint"`
	Api_version string            `yaml:"api_version"`
	Methods     []string          `yaml:"methods"`
	Headers     map[string]string `yaml:"headers"`
	Http_result int               `yaml:"http_result"`
	Body        string            `yaml:"body"`
}

func main() {
	// read file - done
	// convert file to []Proxies - done
	yamlFile, err := ioutil.ReadFile("data.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	fmt.Println("YAML File Content")
	fmt.Println(string(yamlFile))
	proxies := &Proxies{}

	err = yaml.Unmarshal(yamlFile, proxies)

	if err != nil {
		log.Printf("Unmarshal err #%v", err)
	}

	// define endpoint
	// if endpoint exists - return result
	// if not exists - 404 - error endpoint not set
	// start server

}
