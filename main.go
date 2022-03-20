package main

import (
	"fmt"
	"github.com/fyxemmmm/k9s/client-go/k9s"
	"github.com/fyxemmmm/k9s/client-go/rest"
)

func main() {
	cfg, err := k9s.NewForConfig(&rest.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
}