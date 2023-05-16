// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package cli

import (
	"github.com/mdhender/wraithh/generators/clusters"
	"github.com/spf13/cobra"
	"log"
)

// cmdGenerateCluster runs the cluster generator command
var cmdGenerateCluster = &cobra.Command{
	Use:   "cluster",
	Short: "generate a new cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		optCluster := []clusters.Option{}
		if opt, err := clusters.CreateHtmlMap(argsGenerateCluster.mapFile); err != nil {
			log.Fatal(err)
		} else {
			optCluster = append(optCluster, opt)
		}
		if opt, err := clusters.SetKind(argsGenerateCluster.kind); err != nil {
			log.Fatal(err)
		} else {
			optCluster = append(optCluster, opt)
		}
		if opt, err := clusters.SetSystems(128); err != nil {
			log.Fatal(err)
		} else {
			optCluster = append(optCluster, opt)
		}
		if opt, err := clusters.SetRadius(argsGenerateCluster.radius); err != nil {
			log.Fatal(err)
		} else {
			optCluster = append(optCluster, opt)
		}

		if _, err := clusters.Generate(optCluster...); err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

var argsGenerateCluster struct {
	kind    string // uniform, cluster, surface
	mapFile string
	radius  float64
}

func init() {
	cmdGenerate.AddCommand(cmdGenerateCluster)

	// inputs
	cmdGenerateCluster.Flags().StringVar(&argsGenerateCluster.kind, "kind", "uniform", "point distribution (uniform, clustered, sphere)")
	cmdGenerateCluster.Flags().StringVar(&argsGenerateCluster.mapFile, "html-map", "", "name of map file to create (optional)")
	cmdGenerateCluster.Flags().Float64Var(&argsGenerateCluster.radius, "radius", 15.0, "cluster radius")

	// outputs
}
