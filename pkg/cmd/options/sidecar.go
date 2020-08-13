package options

import (
	"os"

	"github.com/pkg/errors"
	"github.com/scylladb/scylla-operator/pkg/naming"
	"github.com/spf13/cobra"
)

// Singleton
var sidecarOpts = &sidecarOptions{
	commonOptions: GetCommonOptions(),
}

type sidecarOptions struct {
	*commonOptions
	CPU      string
	NodeName string
}

func GetSidecarOptions() *sidecarOptions {
	return sidecarOpts
}

func (o *sidecarOptions) AddFlags(cmd *cobra.Command) {
	o.commonOptions.AddFlags(cmd)
	cmd.Flags().StringVar(&o.CPU, "cpu", os.Getenv(naming.EnvVarCPU), "number of cpus to use")
	cmd.Flags().StringVar(&o.NodeName, "node-name", os.Getenv(naming.EnvVarEnvVarNodeName), "node name where the pod is running")
}

func (o *sidecarOptions) Validate() error {
	if err := o.commonOptions.Validate(); err != nil {
		return errors.WithStack(err)
	}
	if o.CPU == "" {
		return errors.New("cpu not set")
	}
	return nil
}
