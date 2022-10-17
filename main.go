package main

import (
	"fmt"
	"github.com/mackerelio/golib/logging"
)

var logger = logging.GetLogger("main")

func main() {
	fmt.Println("fmt.Println hello world")
	logger.Infof("here is main")
}
