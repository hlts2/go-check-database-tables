package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hlts2/go-check-database-tables/dao/databases/config"
	"github.com/hlts2/go-check-database-tables/dao/factories"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List for database tables",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ls(cmd, args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func ls(cmd *cobra.Command, args []string) error {
	c := config.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}

	progress := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	progress.Start()

	dao := factories.FactoryTableDao(dbms, c)
	if dao == nil {
		return errors.New("Invaild Dabase Management System")
	}
	tables, err := dao.GetTables()
	if err != nil {
		progress.Stop()
		return err
	}

	writer := tablewriter.NewWriter(os.Stdout)
	writer.SetHeader([]string{"Table Name"})

	for _, table := range tables {
		writer.Append([]string{table.Name})
	}

	progress.Stop()
	writer.Render()

	return nil
}
