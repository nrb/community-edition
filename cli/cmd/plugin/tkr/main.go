// Copyright 2021 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"os"

	cliv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/cli/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/command/plugin"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/log"
)

func main() {
	var descriptor = cliv1alpha1.PluginDescriptor{
		Name:        "tkr",
		Description: "Generate TKR files",
		Group:       cliv1alpha1.RunCmdGroup,
		Version:     "v0.0.1",
	}
	// plugin!
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "unable to initilize new plugin")
	}

	p.AddCommands(
		PrintCmd,
	)

	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
