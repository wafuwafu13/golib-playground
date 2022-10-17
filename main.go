package main

import (
	"fmt"
	"github.com/mackerelio/golib/logging"
	"github.com/wafuwafu13/golib-sample/metric"
	"github.com/wafuwafu13/golib-sample/check"
)

var logger = logging.GetLogger("main")

func main() {
	logging.SetLogLevel(logging.INFO)
	fmt.Println("fmt.Println hello world")
	logger.Infof("here is main")
	metric.Generate()
	check.Generate()
}
