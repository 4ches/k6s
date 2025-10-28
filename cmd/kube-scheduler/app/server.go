package app

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

func NewSchedulerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "kube-scheduler",
		Long: `The Kubernetes scheduler is a control plane process which assigns
Pods to Nodes. The scheduler determines which Nodes are valid placements for
each Pod in the scheduling queue according to constraints and available
resources. The scheduler then ranks each valid Node and binds the Pod to a
suitable Node. Multiple different schedulers may be used within a cluster;
kube-scheduler is the reference implementation.
See [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/)
for more information about scheduling and the kube-scheduler component.`,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Info("starting kube-scheduler")
			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// TODO: add flags.
	return cmd
}
