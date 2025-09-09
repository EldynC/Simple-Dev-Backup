package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func copyFile(dst, src string) error {
	// Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	// Close the source file
	defer sourceFile.Close()
	// Create the destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	// Close the destination File
	defer destFile.Close()

	_, err = destFile.ReadFrom(sourceFile)
	return err

}
func copyDir(src string, dst string) error {

	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}
		fileName := filepath.Base(path)
		fmt.Println("The file name is ", fileName)
		var excludedPaths = map[string]bool{
			"node_modules":  true, // Node.js dependencies
			"target":        true, // Your target directory
			".venv":         true, // Python virtual environment
			".kiro":         true, // Your custom directory
			".git":          true, // Git repository data
			".DS_Store":     true, // macOS system files
			"Thumbs.db":     true, // Windows thumbnail cache
			".vscode":       true, // VS Code settings
			".idea":         true, // IntelliJ/WebStorm settings
			".env":          true, // Environment files
			".env.local":    true, // Local environment files
			"dist":          true, // Build output directories
			"build":         true, // Build output directories
			"out":           true, // Build output directories
			"coverage":      true, // Test coverage reports
			".nyc_output":   true, // Node.js coverage
			"tmp":           true, // Temporary files
			"temp":          true, // Temporary files
			".cache":        true, // Cache directories
			"vendor":        true, // Go dependencies (if using vendor)
			"__pycache__":   true, // Python cache
			".pytest_cache": true, // Python test cache
			".next":         true, // Next.js build
			".nuxt":         true, // Nuxt.js build
			".svelte-kit":   true, // SvelteKit build
			".tmp":          true, // Temporary files
			".temp":         true, // Temporary files
		}

		if excludedPaths[fileName] {
			fmt.Printf("\n We are skipping %v", fileName)
			if info.IsDir() {
				return filepath.SkipDir // Skip this directory and all its contents
			}
			return nil // Skip this file but continue walking
		}
		outpath := filepath.Join(dst, strings.TrimPrefix(path, src))
		// fmt.Printf("\nThe outpath is %v and the type is %T", outpath, outpath)
		// fmt.Println("\noutpath is ", outpath)
		// fmt.Println("\nInfo is ", info.Name())
		if info.IsDir() {
			os.MkdirAll(outpath, info.Mode())
			return nil
		} else {
			copyFile(outpath, path)
			return nil
		}
	})
}
func getSourceDirectories(source string, target string) {
	err := copyDir(source, target)
	check(err)

}

func getPath(source string) string {
	var absolutePath string = ""
	cmd := exec.Command("pwd")
	// err := cmd.Run()
	out, cmdErr := cmd.Output()

	if cmdErr != nil {
		log.Fatal(cmdErr)
	}
	if source == "." {
		fmt.Println("\nSource is string a dot")
		absolutePath = string(out)
		absolutePath = strings.TrimSpace(absolutePath)
		fmt.Println("\nCommand output:", absolutePath)

	} else if strings.HasPrefix(source, "./") {

		var after []string = strings.Split(source, "./")
		absolutePath = strings.TrimSpace(string(out))
		fmt.Printf("\n%v <-- out | after is %v", strings.TrimSpace(string(out)), after[1])
		absolutePath = absolutePath + "/" + after[1]

	} else {
		absolutePath = source
	}
	fmt.Println(absolutePath)
	return absolutePath

}

func main() {
	source := os.Args[1]
	target := os.Args[2]
	fmt.Printf("\nWe want to copy from %v to %v", source, target)
	var sourceAbsolutePath string = getPath(source)
	var targetAbsolutePath string = getPath(target)
	fmt.Printf("\nWe will copy from %v to %v", sourceAbsolutePath, targetAbsolutePath)
	getSourceDirectories(sourceAbsolutePath, targetAbsolutePath)

}
