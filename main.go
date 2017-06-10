package main

import (
	"fmt"

	"github.com/bijukunjummen/multiple-interfaces-listener/listener"
	"github.com/voxelbrain/goptions"
	"sync"
)

func main() {
	options := struct {
		Server []string      `goptions:"-i, --ip, obligatory, description='IP of the interface'"`
		Port   []int         `goptions:"-p, --port, obligatory, description='Port on the interface'"`
		Help   goptions.Help `goptions:"-h, --help, description='Show this help'"`
	}{ // Default values go here

	}
	goptions.ParseAndFail(&options)

	if len(options.Server) != len(options.Port) {
		fmt.Println("Wrong command line option - the number of ip's and ports should match")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < len(options.Server); i++ {
		l, _ := listener.NewPortListener(options.Server[i], options.Port[i])
		l.ListenAndProvideStockResponses()
	}

	wg.Wait()
}
