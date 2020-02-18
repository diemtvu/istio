package pilot_dbg_cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(rds())
}

func rds() *cobra.Command {
	var resources []string
	localCmd := &cobra.Command{
		Use:   "rds",
		Short: "Show CDS config for resources",
		Long: "Show RDS config for resources",
		Run: func(cmd *cobra.Command, args []string) {
			showRDS(resources)
		},
	}
	localCmd.Flags().StringArrayVarP(&resources, "resources", "r", nil, "Resources to show")

	return localCmd
}

func showRDS(resources []string) {
	pilotClient := NewPilotClient(pilotURL, kubeConfig)

	defer func() {
		pilotClient.Close()
	}()

	pod := NewPodInfo(proxyTag, resolveKubeConfigPath(kubeConfig), proxyType)
	req := pod.appendResources(pod.makeRequest("rds"), resources)
	resp := pilotClient.GetXdsResponse(req)
	Output(resp)
}
