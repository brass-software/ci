package main

import "github.com/mikerybka/util"

func main() {
	ci := &util.CI{
		PeriodMinutes: 15,
		SourceDir:     "/root/server",
		OutFile:       "/usr/local/bin/server",
		ServiceName:   "server",
	}
	ci.Start()
}
