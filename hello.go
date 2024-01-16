package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var hoursPerDay, durationMonths, daysPerWeek int
	var err error

	for {
		fmt.Print("Enter hours per day of ratting: ")
		hoursPerDayInput, _ := reader.ReadString('\n')
		hoursPerDay, err = strconv.Atoi(strings.TrimSpace(hoursPerDayInput))
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please enter a number.")
	}

	for {
		fmt.Print("Select duration (1, 3, 6, 12, or 24 months): ")
		durationMonthsInput, _ := reader.ReadString('\n')
		durationMonths, err = strconv.Atoi(strings.TrimSpace(durationMonthsInput))
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please enter a number.")
	}

	for {
		fmt.Print("Enter days per week of ratting: ")
		daysPerWeekInput, _ := reader.ReadString('\n')
		daysPerWeek, err = strconv.Atoi(strings.TrimSpace(daysPerWeekInput))
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please enter a number.")
	}

	omegaCostPerYear := 18500000000
	tickAmount := 4000000
	ticksPerHour := 3

	totalHours := float64(hoursPerDay*daysPerWeek) * 4.34524
	totalHours *= float64(durationMonths)

	totalIsk := float64(tickAmount*ticksPerHour) * totalHours

	omegaCost := float64(omegaCostPerYear) * (float64(durationMonths) / 12)

	fmt.Printf("\nEstimated ISK earned from ratting: %s\n", humanize.Comma(int64(totalIsk)))
	fmt.Printf("Estimated Omega cost for %d months: %s\n", durationMonths, humanize.Comma(int64(omegaCost)))
	fmt.Printf("Net ISK after Omega payment: %s\n", humanize.Comma(int64(totalIsk-omegaCost)))

	if totalIsk >= omegaCost {
		fmt.Println("Congratulations! Your ratting income is enough to cover Omega.")
	} else {
		fmt.Println("Unfortunately, your ratting income is not enough to cover Omega.")
	}
}