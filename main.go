package main

import (
	gateway "github.com/ivanbatutin921/Anti-bruteforce/gateway"
	app "github.com/ivanbatutin921/Anti-bruteforce/mk/cmd/app"
	//"google.golang.org/genproto/googleapis/cloud/gkeconnect/gateway/v1"
)

func main() {
	app.App()
	gateway.App()
}