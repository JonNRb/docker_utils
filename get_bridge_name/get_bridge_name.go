package main

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
)

func main() {

	if len(os.Args) != 2 {
		panic("no network name specified")
	}

	// silently fail on default networks
	if s := os.Args[1]; s == "bridge" || s == "host" || s == "none" {
		os.Exit(1)
		return
	}

	// TODO: make environment variable for this
	client, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		panic(err)
	}

	nets, err := client.ListNetworks()
	if err != nil {
		panic(err)
	}

	for _, net := range nets {
		if net.Name == os.Args[1] {
			fmt.Printf("br-%s\n", net.ID[:12])
			return
		}
	}

	os.Exit(1)
}
