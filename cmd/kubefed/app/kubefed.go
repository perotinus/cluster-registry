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

package app

import (
	"os"

	"k8s.io/apiserver/pkg/util/logs"
	clusterregistry "k8s.io/cluster-registry/pkg/kubefed"
)

const (
	clusterregistryImageName = "clusterregistry:dev"
	DefaultEtcdImage         = "gcr.io/google_containers/etcd:3.0.17"
)

func Run() error {
	logs.InitLogs()
	defer logs.FlushLogs()

	return clusterregistry.NewClusterregistryCommand(os.Stdout, clusterregistryImageName, DefaultEtcdImage).Execute()
}
