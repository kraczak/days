package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

var RootCmd = &cobra.Command{
	Use:   "days",
	Short: "operations on dates",
}

var inputDateLayout string = "02-01-2006"

func parseDate(date string) time.Time {
	parsedDate, err := time.Parse(inputDateLayout, date)
	if err != nil {
		log.Fatalf("Date must be in format dd-mm-yyyy, but got: \"%s\"", date)
	}
	return parsedDate
}

func parseDay(day string) int {
	days, err := strconv.Atoi(day)
	if err != nil {
		log.Fatalf("Amount of days must be int, but got typ: [%T] and value: [%s]", day, day)
	}
	return days
}

func todayDate() time.Time {
	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	return today
}

func formatDateWithName(date time.Time) string {
	return date.Format("Mon 2 Jan 2006")
}

func formatDate(date time.Time) string {
	return date.Format("2 Jan 2006")
}

