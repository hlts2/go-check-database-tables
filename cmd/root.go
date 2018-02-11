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

var rootCmd = &cobra.Command{
	Use:   "go-check-database-tables",
	Short: "A CLI Tool for checking database tables",
	Run: func(cmd *cobra.Command, args []string) {
		if err := root(cmd, args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var host string
var port int
var user string
var password string
var dbms string
var database string
var databaseTable string

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "Host", "H", "localhost", "Host Name（localhost）")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 3306, "Port（3306）")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "User Name（root）")
	rootCmd.PersistentFlags().StringVarP(&password, "Password", "P", "", "Password")
	rootCmd.PersistentFlags().StringVarP(&dbms, "type", "t", "mysql", "Database Management System（mysql）")
	rootCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "Database Name")
	rootCmd.Flags().StringVarP(&databaseTable, "Table", "T", "", "Database Table")
}

func root(cmd *cobra.Command, args []string) error {
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
		return errors.New("Invalid Dabase Management System")
	}
	describeTables, err := dao.GetTableDescribe(databaseTable)
	if err != nil {
		progress.Stop()
		return err
	}

	writer := tablewriter.NewWriter(os.Stdout)
	writer.SetHeader([]string{"Field", "Type", "Null", "Key", "Default", "Extra"})

	for _, describeTable := range describeTables {
		writer.Append(describeTable.FieldValue())
	}
	progress.Stop()
	writer.Render()

	return nil
}

//Execute run command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
