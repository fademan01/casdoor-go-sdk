package main

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func main() {
	casdoorsdk.InitConfig("http://localhost:8000", "casdoor", "casdoor", "/Users/khaidao/.casdoor/cert.pem", "khaidao", "casdoor-go-sdk")
	var orgate *casdoorsdk.Organization
	orgate, _ = casdoorsdk.GetOrganization("casbin")
	GetOrganizationNames()
	println(orgate.Name)
}
