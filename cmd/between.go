package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var betweenCmd = &cobra.Command{
	Use:   "between [date: yyyy-mm-dd] [date: yyyy-mm-dd]",
	Short: "amount of days between two dates",
	Run: func(cmd *cobra.Command, args []string) {
		date1 := parseDate(args[0])
		date2 := parseDate(args[1])
		daysDifference := date1.Sub(date2).Hours() / 24
		if daysDifference >= 1 {
			dayNum := int(date1.Sub(date2).Hours() / 24)
			fmt.Printf("%s is %d days after %s.\n", formatDateWithName(date1), dayNum, formatDateWithName(date2))
		} else if daysDifference <= -1 {
			dayNum := int(date2.Sub(date1).Hours() / 24)
			fmt.Printf("%s is %d days before %s.\n", formatDateWithName(date1), dayNum, formatDateWithName(date2))
		} else {
			fmt.Println("Don't make me fool. You provided date2's date1, you troll!")
		}
	},
}



func init() {
	RootCmd.AddCommand(betweenCmd)
}
