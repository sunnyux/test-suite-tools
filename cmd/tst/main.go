package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
	"test-suite-tools/pkg/tools"
)

func getFiles(suiteFName string) ([]string, error) {
	testFiles := []string{suiteFName}
	suiteFile, err := os.Open(suiteFName)
	if err != nil {
		return testFiles, err
	}

	scanner := bufio.NewScanner(suiteFile)
	for scanner.Scan() {
		test := scanner.Text()
		testFiles = append(testFiles, test+".in")
		testFiles = append(testFiles, test+".out")
		testFiles = append(testFiles, test+".args")
	}

	return testFiles, suiteFile.Close()
}

func main() {
	// Variables to Store Flag Values
	//==================================================================================================================
	genFile := ""
	refExec := ""
	testExec := ""
	zipFile := ""
	unzipFile := ""
	rm := false
	help := false
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
		{
			Name:  "Matthew Froggatt",
			Email: "",
		},
		{
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
		// TODO rewrite stuff
		cli.StringFlag{
			Name:  "generate-tests, g",
			Usage: "Generates test files based on `GENFILE`",
			/* the file will be formatted as
			===
			name_of_the_test_case
			---
			[argument1 argument2 ...]
			---
			[content for the .in file]
			===
			*/
			Destination: &genFile,
		},
		cli.StringFlag{
			Name:        "produce-outputs, p",
			Usage:       "Produce test outputs based on `EXEC`",
			Destination: &refExec,
		},
		cli.StringFlag{
			Name:        "run-suite, r",
			Usage:       "Run test suite on `EXEC`",
			Destination: &testExec,
		},
		cli.StringFlag{
			Name:        "zip, z",
			Usage:       "Zip test suite to `ZIPFILE`",
			Destination: &zipFile,
		},
		cli.StringFlag{
			Name:        "unzip, u",
			Usage:       "Unzip test suite from `ZIPFILE`",
			Destination: &unzipFile,
		},
		cli.BoolFlag{
			Name:        "remove, rm, delete, d",
			Usage:       "Remove test suite",
			Destination: &rm,
		},
		cli.HelpFlag,
	}
	//==================================================================================================================

	// App Action - Add basic app behaviour
	//==================================================================================================================
	// add "fallthrough" at bottom of case to fallthrough to next one
	app.Action = func(c *cli.Context) {
		g := genFile != ""
		p := refExec != ""
		r := testExec != ""
		z := zipFile != ""
		u := unzipFile != ""

		var suiteFile string
		var testFiles []string
		var firstDelim string
		var secDelim string

		// TODO allow for unzip then any of runSuite and remove
		valid := help || (c.NArg() == 1 && (g || p || r || rm || z)) || (c.NArg() == 0 && u) || (c.NArg() == 3 && g)

		if !valid {
			cli.ShowAppHelpAndExit(c, 1)
		}

		if c.NArg() == 1 || c.NArg() == 3 {
			suiteFile = c.Args()[0]
		}

		if c.NArg() == 1 && (z || rm) {
			testFiles, _ = getFiles(suiteFile)
		}

		if help {
			cli.ShowAppHelpAndExit(c, 0)
		}

		if g {
			fmt.Println("Generating...")
			if c.NArg() == 1 {
				firstDelim = "===\n"
				secDelim = "---\n"
			} else if c.NArg() == 3 {
				firstDelim = c.Args()[1] + "\n"
				secDelim = c.Args()[2] + "\n"
			}
			if err := tools.GenerateFiles(genFile, suiteFile, firstDelim, secDelim); err != nil {
				log.Fatalln(err)
			}
		}
		if p {
			fmt.Println("Running...")
			out, err := exec.Command("produceOutputs", suiteFile, refExec).CombinedOutput()
			fmt.Print(string(out))
			if err != nil {
				log.Fatalln(err)
			}
		}
		if r {
			fmt.Println("Testing...")
			out, err := exec.Command("runSuite", suiteFile, testExec).CombinedOutput()
			fmt.Print(string(out))
			if err != nil {
				log.Fatalln(err)
			}
		}
		if z {
			fmt.Println("Zipping...")
			if err := tools.ZipFiles(zipFile, testFiles); err != nil {
				log.Fatalln(err)
			}
		}
		if u {
			fmt.Println("Unzipping...")
			if err := tools.UnzipHere(unzipFile); err != nil {
				log.Fatalln(err)
			}
		}
		if rm {
			fmt.Println("Removing...")
			for _, file := range testFiles {
				_ = os.Remove(file)
			}
		}

	}

	//==================================================================================================================

	// App Run - This is it
	//==================================================================================================================
	err := app.Run(os.Args)
	fmt.Println("Done!")
	// Only run if something real weird happens
	if err != nil {
		log.Fatalln(err)
	}
	//==================================================================================================================
}
