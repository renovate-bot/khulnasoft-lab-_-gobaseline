package dbmodule

import (
	"errors"

	"github.com/khulnasoft-lab/gobaseline/database"
	"github.com/khulnasoft-lab/gobaseline/dataroot"
	"github.com/khulnasoft-lab/gobaseline/modules"
	"github.com/khulnasoft-lab/gobaseline/utils"
)

var (
	databaseStructureRoot *utils.DirStructure

	module *modules.Module
)

func init() {
	module = modules.Register("database", prep, start, stop)
}

// SetDatabaseLocation sets the location of the database for initialization. Supply either a path or dir structure.
func SetDatabaseLocation(dirStructureRoot *utils.DirStructure) {
	if databaseStructureRoot == nil {
		databaseStructureRoot = dirStructureRoot
	}
}

func prep() error {
	SetDatabaseLocation(dataroot.Root())
	if databaseStructureRoot == nil {
		return errors.New("database location not specified")
	}

	return nil
}

func start() error {
	err := database.Initialize(databaseStructureRoot)
	if err != nil {
		return err
	}

	startMaintenanceTasks()
	return nil
}

func stop() error {
	return database.Shutdown()
}
