// swift-tools-version: 5.7
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "benchmark",
    dependencies: [
        // Dependencies declare other packages that this package depends on.
        .package(name: "km", path: "km"),
        .package(name: "karmem", path: "../swift"),
    ],
    targets: [
        // Targets are the basic building blocks of a package. A target can define a module or a test suite.
        // Targets can depend on other targets in this package, and on products in packages this package depends on.
        .executableTarget(name: "benchmark", dependencies: ["km", "karmem"], path: "swift"),
    ]
)
