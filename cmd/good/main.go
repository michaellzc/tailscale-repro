package main

import (
	"log"
	"time"

	"tailscale.com/tsnet"
)

// upon shutdown, the ephemeral device should be removed from control plane
func main() {
	s := new(tsnet.Server)
	s.Hostname = "test-good-ephemeral-device"
	s.Ephemeral = true

	ln, err := s.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	ln.Close()
	s.Close()
}
