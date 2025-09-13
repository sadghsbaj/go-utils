package main

import (
	"fmt"
	"github.com/sadghsbaj/go-utils/errorutils"
	"github.com/sadghsbaj/go-utils/network"
)

func main() {
	fmt.Println("Starte 'validServerUrls'")
	validServerUrls()

	hyphen := "=============================================================="
	fmt.Printf("%s\n\n", hyphen)

	fmt.Println("Starte 'unvalidServerUrls'")
	unvalidServerUrls()
}

func validServerUrls() {
	port := ":8000"
	fmt.Printf("Test1, port: %s", port)
	e := network.PrintServerUrl(port)
	if errorutils.Handler(e, "error") {}

	port = "8000"
	fmt.Printf("Test2, port: %s", port)
	e = network.PrintServerUrl(port)
	if errorutils.Handler(e, "error") {}

	port = " : 8 0   0 0"
	fmt.Printf("Test3, port: %s", port)
	e = network.PrintServerUrl(port)
	if errorutils.Handler(e, "error") {}
}

func unvalidServerUrls() {
	port := ";8000"
	fmt.Printf("\nTest1, port: %s\n", port)
	e := network.PrintServerUrl(port)
	if errorutils.Handler(e, "error") {}

	port = "80000"
	fmt.Printf("\nTest2, port: %s\n", port)
	e = network.PrintServerUrl(port)
	if errorutils.Handler(e, "error") {}

	port = "-8000"
	fmt.Printf("\nTest3, port: %s\n", port)
	e = network.PrintServerUrl(port)
	if errorutils.Handler(e, "error") {}
}
