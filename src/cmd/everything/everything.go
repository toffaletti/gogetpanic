package main

import (
	"fmt"

	"github.com/golang/glog"

	// this import causes a panic in go get
	"cmd/everything/widget"
)

func main() {
	fmt.Println("I am everything")
	glog.Infoln("I am everything")
}
