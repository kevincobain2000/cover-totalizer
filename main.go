package main

import (
	"fmt"
	"os"

	"github.com/kevincobain2000/cover-totalizer/pkg"
)

var version = "dev"

const ACTION_VERSION = "version"

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("Usage: cover-totalizer <path/to/coverage.xml>\n")
		return
	}
	action := os.Args[1]

	if action == ACTION_VERSION {
		fmt.Printf("cover-totalizer version %s\n", version)
		os.Exit(0)
	}

	coverage(action)
}

func coverage(xmlPath string) {
	total, err := pkg.NewCoverageService().ParseCoveragePercentage(xmlPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("%.2f\n", total)
}
