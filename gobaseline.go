package main

import (
	"os"

	_ "github.com/khulnasoft-lab/gobaseline/api"
	"github.com/khulnasoft-lab/gobaseline/info"
	"github.com/khulnasoft-lab/gobaseline/run"
)

func main() {
	// Set Info
	info.Set("Portbase", "0.0.1", "GPLv3", false)

	// Run
	os.Exit(run.Run())
}
