package main

import (
	"fmt"
	"os"

	"github.com/Com1Software/go-dbase/dbase"
	"golang.org/x/text/encoding/charmap"
)

type Tags struct {
	Tag     string `dbase:"TAG"`
	Testing string `dbase:"TESTING"`
}

func main() {
	tt := "TAGS.DBF"
	if _, err := os.Stat(tt); err == nil {
		table, err := dbase.OpenTable(&dbase.Config{
			Filename:   "TAGS.DBF",
			TrimSpaces: true,
			WriteLock:  true,
		})
		if err != nil {
			panic(err)
		}
		defer table.Close()
		row, err := table.Row()
		if err != nil {
			panic(err)
		}
		p := Tags{
			Testing: "test123",
		}

		row, err = table.RowFromStruct(p)
		if err != nil {
			panic(err)
		}
		fmt.Println(row)
		err = row.FieldByName("TAG").SetValue("TAG_VALUE")
		if err != nil {
			panic(err)
		}
		err = row.FieldByName("TESTING").SetValue("TEST_VALUE")
		if err != nil {
			panic(err)
		}

		err = row.Write()
		if err != nil {
			panic(err)
		}

		fmt.Printf(
			"Last modified: %v Columns count: %v Record count: %v File size: %v \n",
			table.Header().Modified(0),
			table.Header().ColumnsCount(),
			table.Header().RecordsCount(),
			table.Header().FileSize(),
		)

	} else {
		file, err := dbase.NewTable(
			dbase.FoxProVar,
			&dbase.Config{
				Filename:   tt,
				Converter:  dbase.NewDefaultConverter(charmap.Windows1250),
				TrimSpaces: true,
			},
			tcolumns(),
			64,
			nil,
		)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		row, err := file.RowFromStruct(&Tags{
			Tag: "TAG",
		})
		if err != nil {
			panic(err)
		}

		err = row.Add()
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"Last modified: %v Columns count: %v Record count: %v File size: %v \n",
			file.Header().Modified(0),
			file.Header().ColumnsCount(),
			file.Header().RecordsCount(),
			file.Header().FileSize(),
		)

	}

}

func tcolumns() []*dbase.Column {

	tagCol, err := dbase.NewColumn("Tag", dbase.Varchar, 80, 0, false)
	testCol, err := dbase.NewColumn("Testing", dbase.Character, 80, 0, true)
	if err != nil {
		panic(err)
	}
	return []*dbase.Column{
		tagCol,
		testCol,
	}
}
