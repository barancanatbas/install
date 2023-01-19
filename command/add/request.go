package add

import (
	"errors"
	"github.com/barancanatbas/install/command/create"
)

type editRequest struct {
	FileName       string
	PackageName    string
	AllowedPackage string
}

func (r *editRequest) BindRequest(args []string) error {
	for i, arg := range args {
		switch arg {
		case "-f":
			r.FileName = args[i+1]
		case "-p":
			r.PackageName = args[i+1]
		case "-a":
			r.AllowedPackage = args[i+1]
		}
	}

	err := r.validate()
	return err
}

func (r *editRequest) validate() error {
	if r.FileName == "" {
		r.FileName = create.INIT_FILE_NAME
	}

	if r.PackageName == "" {
		return errors.New("package name required")
	}
	if r.AllowedPackage == "" {
		return errors.New("allowed package required")
	}
	return nil
}
