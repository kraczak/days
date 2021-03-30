package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var weekCmd = &cobra.Command{
	Use: "weekday [date: dd-mm-yyyy]",
	Short: "prints the name of the day at given date",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args{
			date := parseDate(arg)
			var answer string
			today := todayDate()
			if date.Before(today) {
				answer = fmt.Sprintf("%s was %s", formatDate(date), date.Weekday())
			} else if date.After(today){
				answer = fmt.Sprintf("%s will be %s", formatDate(date), date.Weekday())
			} else{
				answer = fmt.Sprintf("Today, %s is %s", formatDate(date), date.Weekday())
			}
			fmt.Println(answer)
		}
	},
}

func init()  {
	RootCmd.AddCommand(weekCmd)
}