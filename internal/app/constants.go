package app

import "os"

// GoBuilderDataDir is the name of GoBuilder's Data Directory
const (
	GoBuilderDataDir                           = "gobuilder"
	GoBuilderExecutableDir                     = "gobuilder/.executable"
	GoBuilderExecutableDirKeepFile             = "gobuilder/.executable/.keep"
	GoBuilderTasksDir                          = "gobuilder/tasks"
	GoBuilderTasksDirKeepFile                  = "gobuilder/tasks/.keep"
	GoBuilderCommandFile                       = "gobuilder/cmd.gb"
	DefaultProjectPath                         = "./"
	DefaultPerm                    os.FileMode = 777
)
