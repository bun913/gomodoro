// Package cmd has startCmd defined
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/hatappi/gomodoro/config"
	"github.com/hatappi/gomodoro/net/unix"
)

// remainCmd represents the remain command
var remainCmd = &cobra.Command{
	Use:   "remain",
	Short: "get remain time",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.GetConfig()
		if err != nil {
			return err
		}

		c, err := unix.NewClient(config.UnixDomainScoketPath)
		if err != nil {
			return err
		}

		r, err := c.Get()
		if err != nil {
			return err
		}

		fmt.Printf("%s", r.GetRemain())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(remainCmd)
}