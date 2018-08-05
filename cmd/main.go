package main

import (
	"cron"
	"os"
	"fmt"
	"log"
	"strings"
)

func main() {
	c := cron.Parse(os.Args[1:]...)

	if c == nil {
		log.Fatal("No cron expression detected in")
	}

	fmt.Printf("%-14s| %v\n", "minute", strings.Trim(strings.Join(strings.Fields(fmt.Sprint([]int(c.Minute))), " "), "[]"))
	fmt.Printf("%-14s| %v\n", "hour", strings.Trim(strings.Join(strings.Fields(fmt.Sprint([]int(c.Hour))), " "), "[]"))
	fmt.Printf("%-14s| %v\n", "day", strings.Trim(strings.Join(strings.Fields(fmt.Sprint([]int(c.Day))), " "), "[]"))
	fmt.Printf("%-14s| %v\n", "month", strings.Trim(strings.Join(strings.Fields(fmt.Sprint([]int(c.Month))), " "), "[]"))
	fmt.Printf("%-14s| %v\n", "weekday", strings.Trim(strings.Join(strings.Fields(fmt.Sprint([]int(c.WeekDay))), " "), "[]"))

	cmd := c.Command
	if cmd == "" {
		cmd = "Command not given"
	}
	fmt.Printf("%-14s| %v\n", "command", cmd)
}

