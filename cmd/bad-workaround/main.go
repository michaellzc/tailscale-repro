package main

import (
	"context"
	"log"
	"time"

	"tailscale.com/ipn/store/mem"
	"tailscale.com/tsnet"
)

// upon shutdown, the ephemeral device should be removed from control plane
func main() {
	s := new(tsnet.Server)
	s.Hostname = "test-bad-workaround-ephemeral-device"
	s.Ephemeral = true
	s.Store = new(mem.Store)

	ln, err := s.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	ln.Close()

	// manually logout works
	if client, err := s.LocalClient(); err != nil {
		log.Fatal(err)
	} else {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		if err := client.Logout(ctx); err != nil {
			log.Fatal(err)
		}
	}

	s.Close()
}
