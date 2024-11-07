package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func copyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		log.Println("Error copying file", err)
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyEmbedDir(f embed.FS, origin, target string) error {
	if _, err := os.Stat(target); os.IsNotExist(err) {
		if err := os.MkdirAll(target, os.ModePerm); err != nil {
			err = fmt.Errorf("error creating directory: %v", err)
			return err
		}
	}

	files, err := f.ReadDir(origin)
	if err != nil {
		err = fmt.Errorf("error reading directory: %v", err)
		return err
	}

	for _, file := range files {
		sourceFileName := filepath.Join(origin, file.Name())
		destFileName := filepath.Join(target, file.Name())

		if file.IsDir() {
			if err := copyEmbedDir(f, sourceFileName, destFileName); err != nil {
				err = fmt.Errorf("error copying subdirectory: %v", err)
				return err
			}
			continue
		}

		fileContent, err := f.ReadFile(sourceFileName)
		if err != nil {
			err = fmt.Errorf("error reading file: %v", err)
			return err
		}

		if err := os.WriteFile(destFileName, fileContent, 0644); err != nil { // nolint: gosec
			log.Printf("error os.WriteFile error: %v", err)
			err = fmt.Errorf("error writing file: %w", err)
			return err
		}
	}

	return nil
}

func copyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = copyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = copyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func writetofile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func containsString(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}
