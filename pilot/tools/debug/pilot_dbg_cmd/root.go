package pilot_dbg_cmd

import (
	"github.com/spf13/cobra"
)

var (
	// kubeConfig is the path to the kubeconfig file
	kubeConfig string

	// pilotURL is the pilot/istiod URL.
	pilotURL string

	streaming bool

	// Pod name or app label or istio label to identify the proxy.
	proxyTag string

	// Either sidecar, ingress or router
	proxyType string

	// Path to output file. Leave blank to output to stdout.
	outputFile string

	// If set, output the whole XDS response.
	outputAll bool
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&kubeConfig, "kubeconfig", "k", "~/.kube/config", "path to the kubeconfig file. Default is ~/.kube/config")
	RootCmd.PersistentFlags().StringVarP(&pilotURL, "pilot", "p", "", "pilot address. Will try port forward if not provided.")
	RootCmd.PersistentFlags().BoolVarP(&streaming, "streaming", "s", false, "If set, waiting on streaming gRPC until terminated.")
	RootCmd.PersistentFlags().StringVarP(&proxyTag, "proxytag", "t", "", "Pod name or app label or istio label to identify the proxy.")
	RootCmd.PersistentFlags().StringVarP(&proxyType, "proxytype", "", "sidecar", "sidecar, ingress, router. Default 'sidecar'.")
	RootCmd.PersistentFlags().StringVarP(&outputFile, "out", "o", "", "output file. Leave blank to go to stdout")
	RootCmd.PersistentFlags().BoolVarP(&outputAll, "all", "a", false, "If set, output the whole XDS response.")

	RootCmd.AddCommand(lds())
	RootCmd.AddCommand(cds())
	RootCmd.AddCommand(eds())
	RootCmd.AddCommand(rds())
}

// RootCmd is the root command line.
var RootCmd = &cobra.Command{}
