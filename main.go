package main

import (
	"fmt"
	"os"

	"github.com/kevincobain2000/cover-totalizer/pkg"
)

var version = "dev"

const ACTION_COVERAGE = "coverage"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: cover-totalizer <action>\n")
		return
	}
	action := os.Args[1]

	if action == ACTION_COVERAGE {
		coverage()
		os.Exit(0)
	}
}

func coverage() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: cover-totalizer %s <coverage.xml>\n", os.Args[1])
		return
	}
	xmlFilePath := os.Args[2]
	total, err := pkg.NewCoverageService().ParseCoveragePercentage(xmlFilePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("%.2f\n", total)
}
