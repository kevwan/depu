package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func main() {
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spin.Suffix = " Checking for updates..."
	spin.Start()

	directs, err := getDirectModules()
	if err != nil {
		spin.Stop()
		processError("parseGoMod", err)
		return
	}

	modules, err := getDepPackages()
	if err != nil {
		spin.Stop()
		processError("parseModules", err)
		return
	}

	spin.Stop()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Package", "Current", "Latest", "GoVersion"})
	table.SetBorder(false)
	for _, mod := range modules {
		if mod.Update == nil {
			continue
		}

		if !contains(directs, mod.Path) {
			continue
		}

		table.Append([]string{
			mod.Path,
			mod.Version,
			mod.Update.Version,
			mod.GoVersion,
		})
	}
	table.Render()
}

// contains checks if str is in list.
func contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}

	return false
}

func processError(action string, err error) {
	writer := color.New(color.FgRed)
	if e, ok := err.(*exec.ExitError); ok {
		writer.Fprintf(os.Stderr, "%s: %s", action, string(e.Stderr))
		os.Exit(e.ExitCode())
	} else {
		writer.Fprintf(os.Stderr, "%s: %s\n", action, err.Error())
	}
}
