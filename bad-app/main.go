package main

import (
	"fmt"
	"bad-app/service"
)

func main() {
	service, err := service.NewUserService()
	if err != nil {
		fmt.Println("=======================[ ERROR ]")
		fmt.Println(err)
		return
	}
	fmt.Println("=======================[ OK ]")
	fmt.Println(service)
	fmt.Println("serviceを使ってごにょごにょ")
}
