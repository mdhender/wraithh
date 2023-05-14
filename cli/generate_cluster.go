// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package cli

import (
	"fmt"
	"github.com/mdhender/wraithh/ec"
	"github.com/spf13/cobra"
	"log"
)

// cmdGenerateCluster runs the cluster generator command
var cmdGenerateCluster = &cobra.Command{
	Use:   "cluster",
	Short: "generate a new cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch argsGenerateCluster.kind {
		case "cluster", "sphere", "uniform":
		// okay
		default:
			return fmt.Errorf("invalid kind")
		}
		if argsGenerateCluster.radius < 15 {
			return fmt.Errorf("radius must be at least 15")
		} else if argsGenerateCluster.radius > 45 {
			return fmt.Errorf("radius must be at most 45")
		}
		e := &ec.Engine{}
		if err := e.GenerateCluster(argsGenerateCluster.kind, argsGenerateCluster.radius, 128); err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

var argsGenerateCluster struct {
	kind   string // uniform, cluster, surface
	radius float64
}

func init() {
	cmdGenerate.AddCommand(cmdGenerateCluster)

	// inputs
	cmdGenerateCluster.Flags().StringVar(&argsGenerateCluster.kind, "kind", "uniform", "point distribution (uniform, cluster, sphere)")
	cmdGenerateCluster.Flags().Float64Var(&argsGenerateCluster.radius, "radius", 15.0, "cluster radius")

	// outputs
}
