@echo off
echo CLI Calculator - Go Application
echo ==============================
echo This script will run the calculator application.
echo Make sure Go is installed on your system.
echo.

REM Try to check if Go is installed
where go >nul 2>&1
if %errorlevel% neq 0 (
    echo Error: Go is not installed or not in your PATH.
    echo.
    echo Please install Go from https://golang.org/dl/
    echo After installation, make sure to add Go to your PATH.
    echo.
    pause
    exit /b
)

echo Go is installed. Running the calculator...
echo.

REM Run different operations as examples
echo Example 1: Adding 5 and 10
go run main.go add 5 10
echo.

echo Example 2: Subtracting 3 from 15
go run main.go subtract 15 3
echo.

echo Example 3: Multiplying 4 and 7
go run main.go multiply 4 7
echo.

echo Example 4: Dividing 20 by 4
go run main.go divide 20 4
echo.

echo Example 5: Division by zero error handling
go run main.go divide 10 0
echo.

echo Example 6: Invalid operation
go run main.go modulo 10 3
echo.

echo Test complete.
pause