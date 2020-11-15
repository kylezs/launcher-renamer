package main

// Creates a launcher package for mac containing an email

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// first one is the mac one
var fileNames = [2]string{"launcher.darwin-launchd-pkg.pkg", "launcher.windows-service-msi.msi"}

// PackageLocation contains the location of the base launcher packages
var PackageLocation = filepath.Join(appPath, "./../../launcher-pkgs")

// Check if the mac, prebuilt package exists
func packagesExists() {
	for _, fileName := range fileNames {
		relPath := fmt.Sprintf("%s/%s", PackageLocation, fileName)
		// relPath := filepath.Join(PackageLocation, fileName)
		xPlatformRelPath := filepath.FromSlash(relPath)
		if _, err := os.Stat(xPlatformRelPath); err == nil {
			fmt.Printf("File '%s' exists\n", fileName)
		} else {
			err := fmt.Errorf("File '%s' does not exist", xPlatformRelPath)
			panic(err)
		}
	}
}

func mkNewPackage(emailClean string, lpOs string) {

	var fileName string
	if lpOs == "Windows" {
		fileName = fileNames[1]
	} else if lpOs == "MacOS" {
		fileName = fileNames[0]
	}

	emailBytes := []byte(emailClean)
	emailB64Str := base64.StdEncoding.EncodeToString(emailBytes)

	newPkgName := fmt.Sprintf("%s-%s", emailB64Str, fileName)
	from := fmt.Sprintf("%s/%s", PackageLocation, fileName)
	xFrom := filepath.FromSlash(from)

	emailFolder := createOutputFolder(emailClean)

	to := fmt.Sprintf("%s/%s", emailFolder, newPkgName)
	xTo := filepath.FromSlash(to)
	fmt.Printf("Copying from: %s to %s...\n", xFrom, xTo)
	_, err := copy(xFrom, xTo)
	if err != nil {
		fmt.Println("There was an error copying the file")
		fmt.Println(err)
		os.Exit(1)
	}
}

// Code from: https://opensource.com/article/18/6/copying-files-go
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func renamePackage(formInfo *FormInfo) {

	// First make sure the packages exist
	packagesExists()

	mkNewPackage(formInfo.email, formInfo.buildOs)
}
