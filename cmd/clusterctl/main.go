/*
Copyright 2018 The Kubernetes Authors.

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

package main

import (
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-digitalocean/pkg/cloud/digitalocean/actuators/machine"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
	"sigs.k8s.io/cluster-api/pkg/apis/cluster/common"
)

func main() {
	var err error
	machine.MachineActuator, err = machine.NewMachineActuator(machine.ActuatorParams{})
	if err != nil {
		klog.Fatalf("Error creating cluster provisioner for google : %v", err)
	}
	common.RegisterClusterProvisioner("digitalocean", machine.MachineActuator)
	cmd.Execute()
}
