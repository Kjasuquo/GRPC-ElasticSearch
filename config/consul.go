package config

import (
	"encoding/json"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type IP struct {
	Query string
}

func getip2() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func ServiceRegistryWithConsul(serviceId string, serviceName string, thePort string, configAddress string, tags []string) {
	config := consulapi.DefaultConfig()
	config.Address = configAddress
	// Get a new client
	discoveryClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("Error creating client: ", err)
		panic(err)
	}

	var port int
	address := getip2()
	if thePort[0] == ':' {
		port, _ = strconv.Atoi(thePort[1:len(thePort)])
	} else {
		port, _ = strconv.Atoi(thePort)
	}
	client := discoveryClient
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = serviceId
	registration.Name = serviceName
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	if tags[0] == "GRPC" {
		registration.Check = &consulapi.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", address, port),
			Interval:                       "5s",
			Timeout:                        "5s",
			Notes:                          "Check if the service is alive",
			GRPCUseTLS:                     false,
			DeregisterCriticalServiceAfter: "5s",
		}
	} else {
		registration.Check = &consulapi.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d", address, port),
			Interval:                       "5s",
			Timeout:                        "5s",
			Notes:                          "Check if the service is alive",
			DeregisterCriticalServiceAfter: "5s",
		}
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}
}
