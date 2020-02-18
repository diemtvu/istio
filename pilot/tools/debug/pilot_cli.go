package main

import (
	"istio.io/pkg/log"
	cmd "istio.io/istio/pilot/tools/debug/pilot_dbg_cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Errorf("%v", err)
	}	
}