package main

import (
	"log"
	"time"

	"tailscale.com/ipn/store/mem"
	"tailscale.com/tsnet"
)

// upon shutdown, the ephemeral device should still present in the control plane
// due to ungraceful shutdowns
//
// you should see below logs from tsnet that indicate the problem
//
// 	TryLogout control response: mustRegen=false, newURL=, err=no nodekey to log out
func main() {
	s := new(tsnet.Server)
	s.Hostname = "test-bad-ephemeral-device"
	s.Ephemeral = true
	s.Store = new(mem.Store)

	ln, err := s.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	ln.Close()
	s.Close()
}
