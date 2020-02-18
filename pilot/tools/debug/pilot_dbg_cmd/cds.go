package pilot_dbg_cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/protobuf/ptypes"
	"istio.io/pkg/log"
)

func init() {
	RootCmd.AddCommand(cds())
}

func cds() *cobra.Command {
	var filter string
	localCmd := &cobra.Command{
		Use:   "cds",
		Short: "Show CDS config for addresses",
		Long:  "Show CDS config for addresses",
		Run: func(cmd *cobra.Command, args []string) {
			showCDS(filter)
		},
	}
	localCmd.Flags().StringVarP(&filter, "filter", "f", "", "Show only cluster with name matching the filter")

	return localCmd
}

func showCDS(filter string) {
	pilotClient := NewPilotClient(pilotURL, kubeConfig)
	defer func() {
		pilotClient.Close()
	}()

	pod := NewPodInfo(proxyTag, resolveKubeConfigPath(kubeConfig), proxyType)

	req := pod.makeRequest("cds")
	resp := pilotClient.GetXdsResponse(req)

	if outputAsRaw {
		Output(resp)
		return
	}

	seenClusters := make([]string, 0, len(resp.Resources))
	for _, res := range resp.Resources {
		cluster := &xdsapi.Cluster{}
		if err := ptypes.UnmarshalAny(res, cluster); err != nil {
			log.Errorf("Cannot unmarshal any proto to cluster: %v", err)
			continue
		}
		seenClusters = append(seenClusters, cluster.Name)

		if filter == cluster.Name {
			Output(cluster)
			return
		}
	}

	fmt.Printf("Cannot find any cluster with name %q. Seen:\n", filter)
	for _, c := range seenClusters {
		fmt.Printf("  %s\n", c)
	}
}
