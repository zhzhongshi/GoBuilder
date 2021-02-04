package app

import "os"

// GoBuilderDataDir is the name of GoBuilder's Data Directory
const GoBuilderDataDir = "gobuilder"
const GoBuilderExecutableDir = "gobuilder/.executable"
const GoBuilderTasksDir = "gobuilder/tasks"
const GoBuilderCommandFile = "gobuilder/cmd.gb"
const DefaultProjectPath = "./"
const DefaultPerm os.FileMode = 777
