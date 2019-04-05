# Develop Podoloff Locally

## Running Podoloff for Development (without UI) [macOS]

1. Install MongoDB
```
    brew tap mongodb/brew
    brew install mongodb-community@4.0
```
2. Start MongoDB
    ```brew services start mongodb-community@4.0```
3. Run `go run podoloff.go start`
4. Do the things.
5. Stop podoloff `Ctrl+C`
6. Stop MongoDB
    ```brew services stop mongodb-community@4.0```