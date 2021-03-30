package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var toCmd = &cobra.Command{
	Use:   "to [date: yyyy-mm-dd]",
	Short: "days between now and future date",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			date := parseDate(arg)
			today := todayDate()
			daysDifference := date.Sub(today).Hours() / 24
			if daysDifference >= 1 {
				dayNum := int(date.Sub(today).Hours() / 24)
				fmt.Printf("%s is %d days after now.\n", date.Format("Mon 2 Jan 2006"), dayNum)
			} else if daysDifference <= -1 {
				dayNum := int(today.Sub(date).Hours() / 24)
				fmt.Printf("%s is %d days before now.\n", date.Format("Mon 2 Jan 2006"), dayNum)
			} else {
				fmt.Println("Don't make me fool. You provided today's date, you troll!")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(toCmd)
}
