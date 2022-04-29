package script_cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var changeShopPremiumRate = &cobra.Command{
	Use:   "change_shop_premium_rate",
	Short: "change_shop_premium_rate",
	Run: func(cmd *cobra.Command, args []string) {
		if err := changePremium(); err != nil {
			panic(errors.Wrap(err, "change shop premium rate"))
		}
	},
}

func changePremium() error {
	fmt.Println("test success")
	return nil
}
