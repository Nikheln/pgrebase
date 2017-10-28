package main

import (
	"fmt"
)

// Sanity type encapsulates requirement checks.
type Sanity struct{}

// Check makes sure the fs is ready to be used.
func (sanity *Sanity) Check() (err error) {
	if err = sanity.directoryExists(); err != nil {
		return err
	}
	if err = sanity.typedDirExists(); err != nil {
		return err
	}
	Cfg.ScanFiles()
	if err = sanity.sqlFilesPresent(); err != nil {
		return err
	}

	return
}

// directoryExists checks the provided sql directory is indeed a directory.
func (sanity *Sanity) directoryExists() (err error) {
	if !IsDir(Cfg.SqlDirPath) {
		return fmt.Errorf("%s is not a directory", Cfg.SqlDirPath)
	}

	return
}

// typedDirExists makes sure that at least one of functions/, views/, triggers/, types/ exists.
func (sanity *Sanity) typedDirExists() (err error) {
	directories := make([]string, 0)

	for _, typedDir := range []string{"functions", "triggers", "types", "views"} {
		path := Cfg.SqlDirPath + "/" + typedDir
		if IsDir(path) {
			directories = append(directories, path)
		}
	}

	if len(directories) == 0 {
		return fmt.Errorf("No functions/, triggers/, types/ or views/ directory found in %s", Cfg.SqlDirPath)
	}

	return
}

// sqlFilesPresent checks there are source file.
//
// No need to process any further if there are no sql files to load.
func (sanity *Sanity) sqlFilesPresent() (err error) {
	if len(Cfg.FunctionFiles)+len(Cfg.TriggerFiles)+len(Cfg.ViewFiles) == 0 {
		return fmt.Errorf("Didn't find any sql file in %s", Cfg.SqlDirPath)
	}

	return
}
