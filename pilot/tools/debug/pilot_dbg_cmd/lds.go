package pilot_dbg_cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/protobuf/ptypes"
	"istio.io/pkg/log"
)

func lds() *cobra.Command {
	handler := &ldsHandler{}
	localCmd := makeXDSCmd("lds", handler)
	localCmd.Flags().StringVarP(&handler.filter, "name", "n", "", "Show only listener with this name")
	return localCmd
}

type ldsHandler struct {
	filter string
}

func (c *ldsHandler) makeRequest(pod *PodInfo) *xdsapi.DiscoveryRequest {
	return pod.makeRequest("lds")
}

func (c *ldsHandler) onXDSResponse(resp *xdsapi.DiscoveryResponse) error {
	if outputAll {
		outputJSON(resp)
		return nil
	}
	seenListener := make([]string, 0, len(resp.Resources))
	for _, res := range resp.Resources {
		listener := &xdsapi.Listener{}
		if err := ptypes.UnmarshalAny(res, listener); err != nil {
			log.Errorf("Cannot unmarshal any proto to listener: %v", err)
			continue
		}

		seenListener = append(seenListener, listener.Name)
		if c.filter == listener.Name {
			outputJSON(listener)
			return nil
		}
	}
	msg := fmt.Sprintf("Cannot find any listener with name %q. Seen:\n", c.filter)
	for _, c := range seenListener {
		msg += fmt.Sprintf("  %s\n", c)
	}
	return fmt.Errorf("%s", msg)
}
