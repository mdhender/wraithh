// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package cli

import (
	"github.com/spf13/cobra"
	"log"
	"time"
)

// cmdRoot represents the base command when called without any subcommands
var cmdRoot = &cobra.Command{
	Short:   "wraith: a game engine",
	Long:    `wraith is an engine inspired by better games.`,
	Version: "0.0.1",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		started := time.Now()

		if argsRoot.timeSelf {
			elapsed := time.Now().Sub(started)
			log.Printf("elapsed time: %v\n", elapsed)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the root Command.
func Execute() error {
	return cmdRoot.Execute()
}

var argsRoot struct {
	timeSelf bool
}

func init() {
	cmdRoot.PersistentFlags().BoolVar(&argsRoot.timeSelf, "time", false, "time commands")
}
