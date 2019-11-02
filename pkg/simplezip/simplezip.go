package simplezip

import (
	"archive/zip"
	"io"
	"os"
)

// Adds file filename to zip that is accessed through zipWriter
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

// Creates a new zip file with zipName and zips files into it.
// ONLY WORKS IF zipName DOES NOT ALREADY EXIST
func ZipFiles(zipName string, files []string) error {
	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}

	writer := zip.NewWriter(zipFile)
	defer writer.Close()

	for _, file := range files{
		if err := addFileToZip(writer, file); err != nil {
			return err
		}
	}
	return nil
}


func copyFileFromZip(zipReader *zip.ReadCloser, f *zip.File) error {
	// Make File
	if err := os.MkdirAll(f.Name, os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	defer outFile.Close()
	if err != nil {
		return err
	}

	rc, err := f.Open()
	defer rc.Close()
	if err != nil {
		return err
	}

	if _, err = io.Copy(outFile, rc); err != nil {
		return err
	}

	return nil
}

func UnzipHere(zipFile string) error {
	reader, err := zip.OpenReader(zipFile)
	defer reader.Close()
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		if err = copyFileFromZip(reader, f); err != nil {
			return err
		}
	}

	return nil
}
