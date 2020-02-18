package main

import (
	// "fmt"
	// go_runtime "runtime"
	// "strings"	
	"istio.io/pkg/log"
	cmd "istio.io/istio/pilot/tools/debug/pilot_dbg_cmd"
)

// func init() {
// 	log.SetFormatter(&log.TextFormatter{
// 		ForceColors:true,
// 		FullTimestamp: true,
// 		CallerPrettyfier: func(f *go_runtime.Frame) (string, string) {
// 				s := strings.Split(f.Function, ".")
// 				funcname := s[len(s)-1]

// 				return funcname, fmt.Sprintf("%s:%d", f.File, f.Line)
// 		},
// 	})
// 	log.SetReportCaller(true)
// }

func main() {
	log.Infof("Starting pilot_cli ...")

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Errorf("%v", err)
	}	
}