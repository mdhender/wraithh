// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package cli

import (
	"github.com/mdhender/wraithh/ec"
	"github.com/spf13/cobra"
	"log"
)

// cmdGenerateCluster runs the cluster generator command
var cmdGenerateCluster = &cobra.Command{
	Use:   "cluster",
	Short: "generate a new cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		e := &ec.Engine{}
		if err := e.GenerateCluster(); err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

var argsGenerateCluster struct {
}

func init() {
	cmdGenerate.AddCommand(cmdGenerateCluster)

	// inputs

	// outputs
}
