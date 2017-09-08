/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clusterregistry

import (
	"io"

	"k8s.io/apiserver/pkg/util/flag"
	"k8s.io/client-go/tools/clientcmd"
	kubefedinit "k8s.io/cluster-registry/pkg/kubefed/init"

	"github.com/spf13/cobra"
)

// NewClusterregistryCommand creates the `clusterregistry` command.
func NewClusterregistryCommand(out io.Writer, defaultServerImage, defaultEtcdImage string) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "clusterregistry",
		Short: "clusterregistry runs a cluster registry",
		Long:  "clusterregistry runs a cluster registry.",
		Run:   runHelp,
	}

	// From this point and forward we get warnings on flags that contain "_" separators
	cmds.SetGlobalNormalizationFunc(flag.WarnWordSepNormalizeFunc)

	cmds.AddCommand(kubefedinit.NewCmdInit(out, clientcmd.NewDefaultPathOptions(), defaultServerImage, defaultEtcdImage))

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
