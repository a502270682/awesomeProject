package script_cmd

import "github.com/spf13/cobra"

func init() {
	Command.AddCommand(
		changeShopPremiumRate,
	)
}

var Command = &cobra.Command{
	Use:   "script",
	Short: "tools and scripts",
}
