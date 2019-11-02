package util

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

//CreateTable creates a pre-specified table style to be written in std-out
func CreateTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowSeparator(".")
	table.SetRowLine(true)
	return table
}
