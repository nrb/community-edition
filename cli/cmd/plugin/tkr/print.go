package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/community-edition/cli/cmd/plugin/tkr/convert"
	tkrtypes "github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkr/pkg/types"
)

var PrintCmd = &cobra.Command{
	Use:   "print --file <path to TKR>",
	Short: "print a TKR as a Go struct",
	RunE:  printTkr,
}

type printTkrOpts struct {
	filePath string
}

var co = printTkrOpts{}

func init() {
	PrintCmd.Flags().StringVarP(&co.filePath, "file", "f", "", "TKR file to print")
}

func printTkr(cmd *cobra.Command, args []string) error {
	tkr, err := convert.ReadTKR(co.filePath)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", tkr)

	return nil
}

// components -> kubernetes-sigs_kind -> images[0] -> kindNodeImage -> imagePath
func printAddon(a tkrtypes.Addon, tkr tkrtypes.Bom) {
	// TODO: group/sort by repo?
	// The TKR BOM defines exactly 1 repo :(
	fmt.Printf("----\n")
	fmt.Printf("Name: %s\n", a.PackageName)
	// This is not currently populated, might have to get it from the
	// Components struct
	fmt.Printf("Version: %s\n", a.Version)
	fmt.Printf("Category: %s\n", a.Category)
	fmt.Printf("Templates Image Path: %s\n", a.TemplatesImagePath)
	fmt.Printf("Templates Image Tag: %s\n", a.TemplatesImageTag)
	fmt.Printf("Cluster Types: %s\n", a.ClusterTypes)
	fmt.Printf("Container Images:\n")
	for _, aci := range a.AddonContainerImages {
		fmt.Printf("\t%s\n", aci)
	}
	fmt.Printf("Template Images:\n")
	for _, ati := range a.AddonContainerImages {
		fmt.Printf("\t%s\n", ati)
	}

	fmt.Printf("\n=== Component ===\n")
	c, err := tkr.GetComponent(a.PackageName)
	if err != nil {
		fmt.Printf("Error for %s: %s\n\n", a.PackageName, err)
		return
	}
	fmt.Printf("%s\n", c)

	fmt.Println()
}
