package main

import (
	gateway "github.com/ivanbatutin921/Anti-bruteforce/gateway"
	grpc "github.com/ivanbatutin921/Anti-bruteforce/mk/service/app"
	//"google.golang.org/genproto/googleapis/cloud/gkeconnect/gateway/v1"
)

func main() {
	go grpc.RunGRPCApp()
	gateway.RunHTTPApp()

}
