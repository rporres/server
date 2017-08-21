package main

import (
	"os"

	"github.com/bblfsh/server/runtime"

	"github.com/sirupsen/logrus"
	"github.com/jessevdk/go-flags"
)

var version string
var build string

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	runtime.Bootstrap()
}

func main() {
	parser := flags.NewNamedParser("bblfsh", flags.Default)
	parser.AddCommand("server", "", "Run server", &serverCmd{})
	parser.AddCommand("client", "", "Run client", &clientCmd{})

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				os.Exit(0)
			} else {
				parser.WriteHelp(os.Stderr)
				os.Exit(1)
			}
		}

		logrus.Errorf("exiting with error: %s", err)
		os.Exit(1)
	}

	logrus.Info("exiting without error")
}
