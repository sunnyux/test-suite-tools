# CS246 Test Suite Tools

This is a command line tool that facilitates test creation, submission, and usage for CS 246 assignments. It can be used in combination with the bash scripts `produceOutputs` and `runSuite` that you have written in Assugnment #1. 

## Installation

Download the binary file from the release page (builds are available for Linux, and Mac): [link]

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```
### `produceOutputs` and `runSuite`

Add the bash scripts to your `$PATH`

## Usage

**Generate Test files (`--generate-tests` and `-g`):**

Generate the suite file (`suite_file.txt`), .in, and .args files from a specially formatted file (`genfile.txt`)
```
$ tst -g genfile.txt suite_file.txt
```

An example `genfile.txt` is formatted as the following:
``` text
test1
---
content of the .args file
---
content of the .in file
more content of the .in file
===
test2
---
content of the .args file
...
```
Note that the default delimiter for different files is `"==="` and the default delimiter for the content of the file is `"---"`.

You can use `--separator-a ` or `-a` to change the delimiter for different files, and use `--separator-b ` or `-b` to change the delimiter for the content of files.
These two options must be used with `--generate-tests` or `-g`. For example:
```
$ tst -g genfile.txt -a="^=^" -b "^-^" suite_file.txt
```
This will change the file delimiter from "===" to "^=^", and the delimiter for the content of .in and .args files from "---" to "^-^". All the files will be generated same as before.


**Produce Outputs (`--produce-outputs` and `-p`):**
This option uses the bash script `produceOutputs` included in your `$PATH`
```
$ tst -p ./myprogram suite_file.txt
```
The produceOutputs script runs program on each test in the test suite and, for each test, creates a `.out` file that contains the output produced for that test.
See CS246 Assignment #1 for detail.

**Run suite (`--run-suite ` and `-r`):**
Similar to `produceOutputs`, this option uses the bash script `runSuite` included in your `$PATH`
```
$ tst -r ./myprogram suite_file.txt
```
The runSuite script runs program on each test in the test suite and reports on any tests whose output does not match the expected output.
See CS246 Assignment #1 for detail.

**Zip files for Due Date 1 submission (`--zip` and `-z`):**
```
$ tst -z a0q0.zip suite_file.txt
```
This zips your suite file `suite_file.txt`, together with the associated `.in`, `.out` and `.args` files, into the file `a0q0.zip`.

**Unzip files (`--unzip` and `-u`):**
```
$ tst -u a0q0.zip
```
This unzips `a0q0.zip`.

**Remove the test suite (`--remove`, `--delete`, `--rm`, and `-d`):**
```
$ tst -d suite_file.txt
```
This removes your suite file `suite_file.txt`, together with the associated `.in`, `.out` and `.args` files.


## Authors

* **Matthew Froggatt** 
* **Sunny Xie**

## License

This project is licensed under the GPL License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
