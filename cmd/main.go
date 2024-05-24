package main

import (
	"fmt"

	"github.com/thslopes/bff/apicall"
	"github.com/thslopes/bff/setup"
)

func main() {
	api_call, err := setup.LoadApiCallFromFile("./queries/query.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	apicaller := apicall.Caller{
		Getter: apicall.NewHttpGetter(),
	}

	data, err := apicaller.Get(api_call)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
