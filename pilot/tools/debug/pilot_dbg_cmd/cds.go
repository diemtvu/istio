package pilot_dbg_cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/protobuf/ptypes"
	"istio.io/pkg/log"
)

func cds() *cobra.Command {
	handler := &cdsHandler{}
	localCmd := makeXDSCmd("cds", handler)
	localCmd.Flags().StringVarP(&handler.filter, "name", "n", "", "Show only cluster with this name")
	return localCmd
}

type cdsHandler struct {
	filter string
}

func (c *cdsHandler) makeRequest(pod *PodInfo) *xdsapi.DiscoveryRequest {
	return pod.makeRequest("cds")
}

func (c *cdsHandler) onXDSResponse(resp *xdsapi.DiscoveryResponse) error {
	if outputAll {
		outputJSON(resp)
		return nil
	}
	seenClusters := make([]string, 0, len(resp.Resources))
	for _, res := range resp.Resources {
		cluster := &xdsapi.Cluster{}
		if err := ptypes.UnmarshalAny(res, cluster); err != nil {
			log.Errorf("Cannot unmarshal any proto to cluster: %v", err)
			continue
		}
		seenClusters = append(seenClusters, cluster.Name)

		if c.filter == cluster.Name {
			outputJSON(cluster)
			return nil
		}
	}
	msg := fmt.Sprintf("Cannot find any listener with name %q. Seen:\n", c.filter)
	for _, c := range seenClusters {
		msg += fmt.Sprintf("  %s\n", c)
	}
	return fmt.Errorf("%s", msg)
}
