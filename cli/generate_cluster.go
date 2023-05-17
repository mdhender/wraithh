// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package cli

import (
	"encoding/json"
	"github.com/mdhender/wraithh/generators/clusters"
	"github.com/spf13/cobra"
	"log"
	"os"
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

		c, sy, st, err := clusters.Generate(optCluster...)
		if err != nil {
			log.Fatal(err)
		}
		// adapt c to json
		if data, err := json.MarshalIndent(c, "", "  "); err != nil {
			log.Fatal(err)
		} else if err = os.WriteFile("g1/out/cluster.json", data, 0660); err != nil {
			log.Fatal(err)
		}
		log.Printf("cluster: created g1/out/cluster.json")
		// adapt sy to json
		if data, err := json.MarshalIndent(sy, "", "  "); err != nil {
			log.Fatal(err)
		} else if err = os.WriteFile("g1/out/systems.json", data, 0660); err != nil {
			log.Fatal(err)
		}
		log.Printf("cluster: created g1/out/stars.json")
		// adapt st to json
		if data, err := json.MarshalIndent(st, "", "  "); err != nil {
			log.Fatal(err)
		} else if err = os.WriteFile("g1/out/stars.json", data, 0660); err != nil {
			log.Fatal(err)
		}
		log.Printf("cluster: created g1/out/stars.json")

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
