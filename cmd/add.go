package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var addCmd = &cobra.Command{
	Use:   "add [days: number] [date: yyyy-mm-dd]",
	Short: "Date in a given days after given date. You can omit date, today's date will be taken.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Provide how many days to add.")
		} else if len(args) > 2 {
			log.Fatal("You provided to many arguments.")
		} else if len(args) == 1 {
			args = append(args, todayDate().Format(inputDateLayout))
		}

		days := parseDay(args[0])
		date := parseDate(args[1])
		futureDate := date.Add(time.Duration(time.Hour) * 24 * time.Duration(days))
		fmt.Println(formatDateWithName(futureDate))
	},
}


func init() {
	RootCmd.AddCommand(addCmd)
}
