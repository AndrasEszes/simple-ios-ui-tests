//
//  benchmarkUITestsLaunchTests.swift
//  benchmarkUITests
//
//  Created by Andras Eszes on 2025. 08. 07..
//

import XCTest

final class benchmarkUITestsLaunchTests: XCTestCase {

    override class var runsForEachTargetApplicationUIConfiguration: Bool {
        true
    }

    override func setUpWithError() throws {
        continueAfterFailure = false
    }

    @MainActor
    func testButtonAndAlert() throws {
        let app = XCUIApplication()
        app.launch()
        let button = app.buttons["Show Alert"]
        XCTAssertTrue(button.waitForExistence(timeout: 2), "Show Alert button should exist")
        
        let attachment1 = XCTAttachment(screenshot: app.screenshot())
        attachment1.name = "Show Alert Button"
        attachment1.lifetime = .keepAlways
        add(attachment1)
        
        button.tap()
        let alert = app.alerts["Alert"]
        XCTAssertTrue(alert.waitForExistence(timeout: 2), "Alert should be presented")
        let okButton = alert.buttons["OK"]
        XCTAssertTrue(okButton.waitForExistence(timeout: 2), "OK button should exist on alert")
        
        let attachment2 = XCTAttachment(screenshot: app.screenshot())
        attachment2.name = "Alert"
        attachment2.lifetime = .keepAlways
        add(attachment2)
        
        okButton.tap()
        XCTAssertFalse(alert.exists, "Alert should be dismissed after tapping OK")
    }
}
