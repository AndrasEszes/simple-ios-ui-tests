//
//  ContentView.swift
//  benchmark
//
//  Created by Andras Eszes on 2025. 08. 07..
//

import SwiftUI

struct ContentView: View {
    @State private var showAlert = false
    @Environment(\.colorScheme) var colorScheme
    var body: some View {
        ZStack {
            Color(UIColor { traitCollection in
                traitCollection.userInterfaceStyle == .dark ? .secondarySystemBackground : .systemBackground
            })
            .ignoresSafeArea()
            VStack(spacing: 24) {
                Button("Show Alert") {
                    showAlert = true
                }
                .buttonStyle(.borderedProminent)
                // Show current color scheme for debugging
                HStack {
                    Image(systemName: colorScheme == .dark ? "moon.fill" : "sun.max.fill")
                        .foregroundStyle(colorScheme == .dark ? .yellow : .orange)
                    Text(colorScheme == .dark ? "Dark Mode" : "Light Mode")
                        .font(.headline)
                        .foregroundStyle(.secondary)
                }
            }
            .alert("Alert", isPresented: $showAlert) {
                Button("OK", role: .cancel) { }
            }
        }
    }
}

#Preview {
    ContentView()
}
