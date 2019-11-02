package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func init() {
}

func main() {
	// Variables to Store Flag Values
	//==================================================================================================================
	var help, rm bool
	//==================================================================================================================
	
	// App Creation - This is all that is needed
	//==================================================================================================================
	app := cli.NewApp()
	//==================================================================================================================

	// App Help Menu Setup - A lot of stuff needs to go here
	//==================================================================================================================
	app.Name = "Test Suite Tools"
	app.Usage = "A set of tools to make CS 246 test suite creation and usage easier."
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Matthew Froggatt",
			Email: "",
		},
		cli.Author{
			Name:  "Sunny Xie",
			Email: "",
		},
	}
	app.HideHelp = true
	app.HideVersion = true
	cli.HelpFlag = cli.BoolFlag{
		Name:        "help, h",
		Usage:       "Show help",
		Destination: &help,
	}
	//==================================================================================================================

	// Argument Definitions - Need to add flags and positional arguments
	//==================================================================================================================
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "produce-outputs, p",
			Usage: "Produce test outputs based on `EXEC`",
		},
		cli.StringFlag{
			Name:  "run-suite, r",
			Usage: "Run test suite on `EXEC`",
		},
		cli.StringFlag{
			Name:  "zip, z",
			Usage: "Zip test suite to `ZIPFILE`",
		},
		cli.StringFlag{
			Name:        "unzip, u",
			Usage:       "Unzip test suite from `ZIPFILE`",
		},
		cli.BoolFlag{
			Name:  "remove, rm",
			Usage: "Remove test suite",
			Destination: &rm,
		},
		cli.HelpFlag,
	}
	//==================================================================================================================

	// App Action - Add basic app behaviour
	//==================================================================================================================
	app.Action = func(c *cli.Context) {
		switch {
		case rm:
			fmt.Println("Removing...")
		case help:
			cli.ShowAppHelpAndExit(c, 0)
		default:
			cli.ShowAppHelpAndExit(c, 1)
		}
		fmt.Println("Done!")
	}
	//==================================================================================================================

	// App Run - This is it
	//==================================================================================================================
	err := app.Run(os.Args)
	// Only run if something real weird happens
	if err != nil {
		//log.Fatalln(err)
	}
	//==================================================================================================================
}
