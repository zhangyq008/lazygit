package main

import "github.com/fatih/color"

// File : A staged/unstaged file
// TODO: decide whether to give all of these the Git prefix
type File struct {
	Name               string
	HasStagedChanges   bool
	HasUnstagedChanges bool
	Tracked            bool
	Deleted            bool
	HasMergeConflicts  bool
	DisplayString      string
}

func (f *File) getDisplayString() string {
	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	if !f.Tracked && !f.HasStagedChanges {
		return red.Sprint(f.DisplayString)
	}
	output := green.Sprint(f.DisplayString[0:1])
	output += red.Sprint(f.DisplayString[1:3])

	if f.HasUnstagedChanges {
		output += red.Sprint(f.Name)
	} else {
		output += green.Sprint(f.Name)
	}
	return output
}
