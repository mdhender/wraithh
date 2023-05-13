// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package cli

import "github.com/spf13/cobra"

// cmdGenerate runs the generate command
var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "generate things",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cmdRoot.AddCommand(cmdGenerate)
}
