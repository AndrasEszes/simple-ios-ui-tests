package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Copy the UI test file with a counter many times what the script gets as argument
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <number_of_files> <number_of_test_functions>")
		os.Exit(1)
	}

	numFiles, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number for number of files\n", os.Args[1])
		os.Exit(1)
	}

	numFunctions, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number for number of test functions\n", os.Args[2])
		os.Exit(1)
	}

	if numFiles <= 0 {
		fmt.Println("Error: Number of files must be greater than 0")
		os.Exit(1)
	}

	if numFunctions <= 0 {
		fmt.Println("Error: Number of test functions must be greater than 0")
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
		// Create new filename
		newFileName := fmt.Sprintf("benchmarkUITestsLaunchTests%d.swift", i)
		newFilePath := filepath.Join(targetDir, newFileName)

		// Create new class name
		newClassName := fmt.Sprintf("benchmarkUITestsLaunchTests%d", i)

		// Replace class name in content
		newContent := strings.Replace(originalContent, "benchmarkUITestsLaunchTests", newClassName, -1)

		// Replace filename in header comment
		newContent = strings.Replace(newContent, "benchmarkUITestsLaunchTests.swift", newFileName, 1)

		// Extract the test function to duplicate it
		testFunctionPattern := `(?s)@MainActor\s+func testButtonAndAlert\(\) throws \{.*?\n    \}`
		re := regexp.MustCompile(testFunctionPattern)
		matches := re.FindAllString(newContent, -1)

		if len(matches) == 0 {
			fmt.Printf("Warning: Could not find testButtonAndAlert function in %s\n", newFileName)
			continue
		}

		originalTestFunction := matches[0]

		// Generate multiple test functions
		var testFunctions strings.Builder
		for j := 1; j <= numFunctions; j++ {
			indexedTestFunction := strings.Replace(originalTestFunction, "testButtonAndAlert", fmt.Sprintf("testButtonAndAlert%d", j), 1)
			testFunctions.WriteString(indexedTestFunction)
			if j < numFunctions {
				testFunctions.WriteString("\n\n    ")
			}
		}

		// Replace the original test function with multiple indexed ones
		newContent = strings.Replace(newContent, originalTestFunction, testFunctions.String(), 1)

		// Write the new file
		err := os.WriteFile(newFilePath, []byte(newContent), 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", newFilePath, err)
			continue
		}

		fmt.Printf("Created: %s with class %s and %d test functions\n", newFilePath, newClassName, numFunctions)
	}

	fmt.Printf("Successfully created %d test file copies, each with %d test functions\n", numFiles, numFunctions)
}
