package pilot_dbg_cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/protobuf/ptypes"
	"istio.io/pkg/log"	
)

func init() {
	RootCmd.AddCommand(lds())
}

func lds() *cobra.Command {
	var filter string
	localCmd := &cobra.Command{
		Use:   "lds",
		Short: "Show LDS config for addresses",
		Long:  "Show LDS config for addresses",
		Run: func(cmd *cobra.Command, args []string) {
			showLDS(filter)
		},
	}
	localCmd.Flags().StringVarP(&filter, "filter", "f", "", "Show only listener with name matching the filter")

	return localCmd
}

func showLDS(filter string) {
	pilotClient := NewPilotClient(pilotURL, kubeConfig)

	defer func() {
		pilotClient.Close()
	}()

	pod := NewPodInfo(proxyTag, resolveKubeConfigPath(kubeConfig), proxyType)

	req := pod.makeRequest("lds")
	resp := pilotClient.GetXdsResponse(req)
	
	if outputAsRaw {
		Output(resp)
		return
	}
	seenListener := make([]string, 0, len(resp.Resources))
	for _, res := range resp.Resources {
		listener := &xdsapi.Listener{}
		if err := ptypes.UnmarshalAny(res, listener); err != nil {
			log.Errorf("Cannot unmarshal any proto to listener: %v", err)
			continue
		}

		seenListener = append(seenListener, listener.Name)
		if filter == listener.Name {
			Output(listener)
			return
		}
	}
	fmt.Printf("Cannot find any listener with name %q. Seen:\n", filter)
	for _, c := range seenListener {
		fmt.Printf("  %s\n", c)
	}
}
