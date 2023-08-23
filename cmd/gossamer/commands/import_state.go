// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package commands

import (
	"fmt"

	"github.com/ChainSafe/gossamer/dot"
	"github.com/ChainSafe/gossamer/lib/trie"
	"github.com/ChainSafe/gossamer/lib/utils"
	"github.com/spf13/cobra"
)

func init() {
	ImportStateCmd.Flags().String("state-file", "", "Path to JSON file consisting of key-value pairs")
	ImportStateCmd.Flags().String("state-version", trie.DefaultStateVersion.String(), "State version v0 or v1")
	ImportStateCmd.Flags().String("header-file", "", "Path to JSON file of block header corresponding to the given state")
	ImportStateCmd.Flags().Uint64("first-slot", 0, "The first BABE slot of the network")
}

// ImportStateCmd is the command to import a state from a JSON file
var ImportStateCmd = &cobra.Command{
	Use:   "import-state",
	Short: "Import state from a JSON file and set it as the chain head state",
	Long: `The import-state command allows a JSON file containing a given state
in the form of key-value pairs to be imported.
Input can be generated by using the RPC function state_getPairs.
Example: 
	gossamer import-state --state-file state.json --state-version v1 --header-file header.json --first-slot <first slot of network>`, //nolint:lll
	RunE: func(cmd *cobra.Command, args []string) error {
		return execImportState(cmd)
	},
}

func execImportState(cmd *cobra.Command) error {
	if basePath == "" {
		basePath = config.BasePath
	}

	if basePath == "" {
		return fmt.Errorf("basepath must be specified")
	}

	firstSlot, err := cmd.Flags().GetUint64("first-slot")
	if err != nil {
		return fmt.Errorf("failed to get first-slot: %s", err)
	}

	stateFile, err := cmd.Flags().GetString("state-file")
	if err != nil {
		return fmt.Errorf("failed to get state-file: %s", err)
	}
	if stateFile == "" {
		return fmt.Errorf("state-file must be specified")
	}

	headerFile, err := cmd.Flags().GetString("header-file")
	if err != nil {
		return fmt.Errorf("failed to get header-file: %s", err)
	}
	if headerFile == "" {
		return fmt.Errorf("header-file must be specified")
	}

	stateVersionFlag, err := cmd.Flags().GetString("state-version")
	if err != nil {
		return fmt.Errorf("failed to get state-version: %s", err)
	}
	stateVersion, err := trie.ParseVersion(stateVersionFlag)
	if err != nil {
		return fmt.Errorf("failed to parse state-version: %s", err)
	}

	basePath = utils.ExpandDir(basePath)

	return dot.ImportState(basePath, stateFile, headerFile, firstSlot, stateVersion)
}
