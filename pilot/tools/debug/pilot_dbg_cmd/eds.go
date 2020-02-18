package pilot_dbg_cmd

import (
	// "fmt"

	"github.com/spf13/cobra"
	// xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	// "github.com/golang/protobuf/ptypes"
	// "istio.io/pkg/log"
	// "github.com/golang/protobuf/jsonpb"
)

func init() {
	RootCmd.AddCommand(eds())
}

func eds() *cobra.Command {
	var resources []string
	localCmd := &cobra.Command{
		Use:   "eds",
		Short: "Show ClusterLoadAssignment (EDS) for given resources",
		Long:  `
Show ClusterLoadAssignment (EDS) for given resources. Example:
pilot_cli --proxytag=httpbin eds -r "outbound|8000||httpbin.default.svc.cluster.local"
`,
		Run: func(cmd *cobra.Command, args []string) {
			showEDS(resources)
		},
	}
	localCmd.Flags().StringArrayVarP(&resources, "resources", "r", nil, "Resources to show")

	return localCmd
}

func showEDS(resources []string) {
	pilotClient := NewPilotClient(pilotURL, kubeConfig)

	defer func() {
		pilotClient.Close()
	}()

	pod := NewPodInfo(proxyTag, resolveKubeConfigPath(kubeConfig), proxyType)

	req := pod.appendResources(pod.makeRequest("eds"), resources)
	resp := pilotClient.GetXdsResponse(req)
	Output(resp)
}
