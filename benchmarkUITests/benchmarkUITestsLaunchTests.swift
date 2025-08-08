//
//  benchmarkUITestsLaunchTests.swift
//  benchmarkUITests
//
//  Created by Andras Eszes on 2025. 08. 07..
//

import XCTest

final class benchmarkUITestsLaunchTests: XCTestCase {

    override class var runsForEachTargetApplicationUIConfiguration: Bool {
        false
    }

    override func setUpWithError() throws {
        continueAfterFailure = false
    }

    @MainActor
    func testButtonAndAlert() throws {
        let app = XCUIApplication()
        app.launch()
        
        for _ in 1...1 {
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
    }
}
