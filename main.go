package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// generateTestMethods creates N test methods with M screenshots each
func generateTestMethods(numTestMethods int, screenshotsPerTest int) string {
	var testMethods strings.Builder
	
	for i := 1; i <= numTestMethods; i++ {
		if i > 1 {
			testMethods.WriteString("\n    ")
		}
		
		testMethods.WriteString(fmt.Sprintf(`@MainActor
    func testButtonAndAlert%d() throws {
        let app = XCUIApplication()
        app.launch()
        
        for i in 1...%d {
            let screenshot = app.screenshot()
            let attachment = XCTAttachment(screenshot: screenshot)
            attachment.name = "testButtonAndAlert%d_screenshot_\(i)"
            attachment.lifetime = .keepAlways
            add(attachment)
            
            let button = app.buttons["Show Alert"]
            XCTAssertTrue(button.waitForExistence(timeout: 2), "Show Alert button should exist")
            
            button.tap()
            let alert = app.alerts["Alert"]
            XCTAssertTrue(alert.waitForExistence(timeout: 2), "Alert should be presented")
            let okButton = alert.buttons["OK"]
            XCTAssertTrue(okButton.waitForExistence(timeout: 2), "OK button should exist on alert")
            
            okButton.tap()
            XCTAssertFalse(alert.exists, "Alert should be dismissed after tapping OK")
        }
    }`, i, screenshotsPerTest, i))
	}
	
	return testMethods.String()
}

// Copy the UI test file and generate N test methods with M screenshots each
func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <number_of_files> <number_of_test_methods> <screenshots_per_test>")
		os.Exit(1)
	}

	numFiles, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number for number of files\n", os.Args[1])
		os.Exit(1)
	}

	numTestMethods, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number for number of test methods\n", os.Args[2])
		os.Exit(1)
	}

	screenshotsPerTest, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number for screenshots per test\n", os.Args[3])
		os.Exit(1)
	}

	if numFiles <= 0 {
		fmt.Println("Error: Number of files must be greater than 0")
		os.Exit(1)
	}

	if numTestMethods <= 0 {
		fmt.Println("Error: Number of test methods must be greater than 0")
		os.Exit(1)
	}

	if screenshotsPerTest <= 0 {
		fmt.Println("Error: Screenshots per test must be greater than 0")
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

		// Generate multiple test methods with screenshots
		testMethods := generateTestMethods(numTestMethods, screenshotsPerTest)
		
		// Replace the original test method with multiple generated ones
		originalTestPattern := `(?s)@MainActor\s+func testButtonAndAlert\(\) throws \{.*?\n    \}`
		re := regexp.MustCompile(originalTestPattern)
		newContent = re.ReplaceAllString(newContent, testMethods)

		// Write the new file
		err := os.WriteFile(newFilePath, []byte(newContent), 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", newFilePath, err)
			continue
		}

		fmt.Printf("Created: %s with class %s containing %d test methods with %d screenshots each\n", newFilePath, newClassName, numTestMethods, screenshotsPerTest)
	}

	totalScreenshots := numFiles * numTestMethods * screenshotsPerTest
	fmt.Printf("Successfully created %d test files with %d test methods each, generating %d total screenshots\n", numFiles, numTestMethods, totalScreenshots)
}
