/*
Copyright © 2024 knightchaser
*/
package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

// sha256Cmd represents the sha256 command
var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "Calculate SHA256 hash digest",
	Long:  "Calculate SHA256 hash digest of the given input as an argument",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		hash := CalculateSHA256(input)
		fmt.Println(hash)
	},
}

func CalculateSHA256(input string) string {
	sha256hasher := sha256.New()
	sha256hasher.Write([]byte(input))
	hash := sha256hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func init() {
	sha256Cmd.Flags().StringP("input", "i", "", "Input string for SHA256 calculation")

	rootCmd.AddCommand(sha256Cmd)
}
