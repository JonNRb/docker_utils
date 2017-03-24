package main

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
)

func main() {

	if len(os.Args) != 3 {
		panic("must specify container and network name")
	}

	targetContainer, targetNetwork := os.Args[1], os.Args[2]

	client, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		panic(err)
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{All:true})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {

		var found = false

		for _, containerName := range container.Names {
			if containerName == "/" + targetContainer {
				found = true
				break
			}
		}

		if !found {
			continue
		}

		for networkName, network := range container.Networks.Networks {
			if networkName != targetNetwork {
				continue
			}
			fmt.Printf("%s\n", network.IPAddress)
			return
		}

		// if network not found for this container
		panic("doesn't seem as if container \"" + targetContainer + "\" is connected to network \"" + targetNetwork + "\"")

	}

	panic("no container with name \"" + targetContainer + "\" found")
}
