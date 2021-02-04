package app

import "os"

// GoBuilderDataDir is the name of GoBuilder's Data Directory
const GoBuilderDataDir = "gobuilder"
const GoBuilderExecutableDir = "gobuilder/.executable"
const GoBuilderExecutableDirKeepFile = "gobuilder/.executable/.keep"
const GoBuilderTasksDir = "gobuilder/tasks"
const GoBuilderTasksDirKeepFile = "gobuilder/tasks/.keep"
const GoBuilderCommandFile = "gobuilder/cmd.gb"
const DefaultProjectPath = "./"
const DefaultPerm os.FileMode = 777
