package pilot_dbg_cmd

import (
	// "fmt"

	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/spf13/cobra"
	// xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	// "github.com/golang/protobuf/ptypes"
	// "istio.io/pkg/log"
	// "github.com/golang/protobuf/jsonpb"
)

func eds() *cobra.Command {
	handler := &edsHandler{}
	localCmd := makeXDSCmd("eds", handler)
	localCmd.Flags().StringArrayVarP(&handler.resources, "resources", "r", nil, "Resources to show")
	return localCmd
}

type edsHandler struct {
	resources []string
}

func (c *edsHandler) makeRequest(pod *PodInfo) *xdsapi.DiscoveryRequest {
	return pod.appendResources(pod.makeRequest("eds"), c.resources)
}

func (c *edsHandler) onXDSResponse(resp *xdsapi.DiscoveryResponse) error {
	outputJSON(resp)
	return nil
}
