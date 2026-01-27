package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	username := os.Args[1]
	// args:= os.Args[2:]

	switch username {
	case "-h", "--help":
		printUsage()
	default:
		event, _ := fetchEvents(username)
		printEvents(event, username)
	}
}

func printUsage() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "GitHub User Activity")
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "\t-h, --help\tShow this help message")
	fmt.Fprintln(w, "\t./github-activity <username>\tFetch and display recent GitHub activity for the specified username")

	w.Flush()
}
