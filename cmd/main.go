package main

import (
	"fmt"

	"github.com/thslopes/bff/apicall"
	"github.com/thslopes/bff/setup"
)

func main() {
	mappings, err := setup.LoadApiCallFromFile("./mappings/query.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	apicaller := apicall.Caller{
		Getter: apicall.NewHttpGetter(),
	}

	for _, v := range mappings {
		data, err := apicaller.Get(v)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data))
	}
}
