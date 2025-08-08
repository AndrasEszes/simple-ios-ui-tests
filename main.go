package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Copy the UI test file with a counter many times what the script gets as argument
func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <number_of_files> <min_loop_iterations> <max_loop_iterations>")
		os.Exit(1)
	}

	numFiles, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number for number of files\n", os.Args[1])
		os.Exit(1)
	}

	minLoopIterations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number for minimum loop iterations\n", os.Args[2])
		os.Exit(1)
	}

	maxLoopIterations, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number for maximum loop iterations\n", os.Args[3])
		os.Exit(1)
	}

	if numFiles <= 0 {
		fmt.Println("Error: Number of files must be greater than 0")
		os.Exit(1)
	}

	if minLoopIterations <= 0 {
		fmt.Println("Error: Minimum loop iterations must be greater than 0")
		os.Exit(1)
	}

	if maxLoopIterations <= 0 {
		fmt.Println("Error: Maximum loop iterations must be greater than 0")
		os.Exit(1)
	}

	if minLoopIterations > maxLoopIterations {
		fmt.Println("Error: Minimum loop iterations cannot be greater than maximum loop iterations")
		os.Exit(1)
	}

	sourceFile := "benchmarkUITests/benchmarkUITestsLaunchTests.swift"
	targetDir := "benchmarkUITests"

	// Read the original file
	content, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("Error reading source file: %v\n", err)
		os.Exit(1)
	}

	originalContent := string(content)

	for i := 1; i <= numFiles; i++ {
		// Generate random loop iterations for this file
		loopIterations := minLoopIterations + rand.Intn(maxLoopIterations-minLoopIterations+1)

		// Create new filename
		newFileName := fmt.Sprintf("benchmarkUITestsLaunchTests%d.swift", i)
		newFilePath := filepath.Join(targetDir, newFileName)

		// Create new class name
		newClassName := fmt.Sprintf("benchmarkUITestsLaunchTests%d", i)

		// Replace class name in content
		newContent := strings.Replace(originalContent, "benchmarkUITestsLaunchTests", newClassName, -1)

		// Replace filename in header comment
		newContent = strings.Replace(newContent, "benchmarkUITestsLaunchTests.swift", newFileName, 1)

		// Replace the loop counter in the test function
		loopPattern := `for _ in 1\.\.\.1 \{`
		newLoopStatement := fmt.Sprintf("for _ in 1...%d {", loopIterations)
		re := regexp.MustCompile(loopPattern)
		newContent = re.ReplaceAllString(newContent, newLoopStatement)

		// Write the new file
		err := os.WriteFile(newFilePath, []byte(newContent), 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", newFilePath, err)
			continue
		}

		fmt.Printf("Created: %s with class %s and loop iterations set to %d\n", newFilePath, newClassName, loopIterations)
	}

	fmt.Printf("Successfully created %d test file copies with random loop iterations between %d and %d\n", numFiles, minLoopIterations, maxLoopIterations)
}
