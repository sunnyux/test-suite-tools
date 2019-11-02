package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"log"
	"os"
	"os/exec"
)

func addFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

func getFiles(suiteFName string) ([]string, error) {
	var testFiles []string
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
			Name:        "generate-tests, g",
			Usage:       "Generates test files based on `GENFILE`",
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
			Name:        "remove, rm",
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

		// TODO allow for unzip then any of runSuite and remove
		valid := help || (c.NArg() == 1 && (g || p || r || rm || z)) || (c.NArg() == 0 && u)

		if !valid {
			cli.ShowAppHelpAndExit(c, 1)
		}

		if c.NArg() == 1 {
			suiteFile = c.Args()[0]
			testFiles, _ = getFiles(suiteFile)
		}

		if help {
			cli.ShowAppHelpAndExit(c, 0)
		}
		if g {
		}
		if p {
			out, err := exec.Command("produceOutputs", refExec, suiteFile).CombinedOutput()
			fmt.Print(out)
			if err != nil {
				log.Fatalln(err)
			}
		}
		if r {
			out, err := exec.Command("runSuite", testExec, suiteFile).CombinedOutput()
			fmt.Print(out)
			if err != nil {
				log.Fatalln(err)
			}
		}
		if z {
			_ = os.Remove(zipFile)
			newFile, err := os.Create(zipFile)
			if err != nil {
				log.Fatalln(err)
			}
			writer := zip.NewWriter(newFile)

			for _, file := range testFiles {
				if err = addFileToZip(writer, file); err != nil {
					log.Fatalln(err)
				}
			}

			err = writer.Close()
			if err != nil {
				log.Println(writer)
			}
		}
		if u {

		}
		if rm {
			fmt.Println("Removing...")
			for _, file := range testFiles {
				_ = os.Remove(file)
			}
		}
	}
	fmt.Println("Done!")

	//==================================================================================================================

	// App Run - This is it
	//==================================================================================================================
	err := app.Run(os.Args)
	// Only run if something real weird happens
	if err != nil {
		log.Fatalln(err)
	}
	//==================================================================================================================
}
