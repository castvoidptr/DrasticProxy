package main

import (
	"drasticproxy/proxy"
)

func main() {
	proxy := proxy.CreateProxy(25565)

	successful := proxy.Init()

	if !successful {
		return
	}

	proxy.Start()
}
