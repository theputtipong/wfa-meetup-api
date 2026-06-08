<#
.SYNOPSIS
WFA Meetup API - Windows Setup Script
This script automates the setup process for Windows developers.
#>

$ErrorActionPreference = 'Stop'

function Invoke-WindowsSetup {
    Write-Host 'WFA Meetup API - Windows Setup' -ForegroundColor Cyan
    Write-Host '======================================' -ForegroundColor Cyan
    Write-Host ''

    # Check if running on Windows
    $isWindowsOS = $IsWindows -or ($PSVersionTable.PSEdition -eq 'Desktop') -or ([Environment]::OSVersion.Platform -match 'Win32NT')
    if (-not $isWindowsOS) {
        Write-Host 'This script is intended for Windows.' -ForegroundColor Yellow
        exit 1
    }

    Write-Host 'Running on Windows' -ForegroundColor Green
    Write-Host ''

    # Check for winget
    if (Get-Command winget -ErrorAction SilentlyContinue) {
        Write-Host 'winget found' -ForegroundColor Green
    }
    else {
        Write-Host 'winget not found. Please install App Installer from Microsoft Store.' -ForegroundColor Yellow
    }

    # Check for Go
    if (Get-Command go -ErrorAction SilentlyContinue) {
        $GoVersion = (go version).Split(' ')[2]
        Write-Host "Go found: $GoVersion" -ForegroundColor Green
    }
    else {
        Write-Host 'Installing Go...' -ForegroundColor Cyan
        winget install GoLang.Go --source winget --silent
        Write-Host 'Go installed. You may need to restart your terminal for path changes.' -ForegroundColor Green
    }

    # Check for Make
    if (Get-Command make -ErrorAction SilentlyContinue) {
        Write-Host 'make found' -ForegroundColor Green
    }
    else {
        Write-Host 'Installing make...' -ForegroundColor Cyan
        winget install ezwinports.make --source winget --silent
        Write-Host 'make installed' -ForegroundColor Green
    }

    # Check for Docker
    if (Get-Command docker -ErrorAction SilentlyContinue) {
        Write-Host 'Docker found' -ForegroundColor Green
    }
    else {
        Write-Host 'Docker CLI not found.' -ForegroundColor Yellow
        Write-Host 'Please install Docker Desktop: https://www.docker.com/products/docker-desktop' -ForegroundColor Yellow
        Write-Host 'Or use winget: winget install Docker.DockerDesktop' -ForegroundColor Yellow
    }

    Write-Host ''
    Write-Host 'Installing Go dependencies...' -ForegroundColor Cyan
    go mod download
    go mod tidy
    Write-Host 'Dependencies installed' -ForegroundColor Green

    Write-Host ''
    Write-Host 'Installing development tools...' -ForegroundColor Cyan

    if (Get-Command golangci-lint -ErrorAction SilentlyContinue) {
        Write-Host 'golangci-lint found' -ForegroundColor Green
    }
    else {
        Write-Host 'Installing golangci-lint...' -ForegroundColor Cyan
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
        Write-Host 'golangci-lint installed' -ForegroundColor Green
    }

    if (Get-Command air -ErrorAction SilentlyContinue) {
        Write-Host 'air found' -ForegroundColor Green
    }
    else {
        Write-Host 'Installing air...' -ForegroundColor Cyan
        go install github.com/cosmtrek/air@latest
        Write-Host 'air installed' -ForegroundColor Green
    }

    if (Get-Command goimports -ErrorAction SilentlyContinue) {
        Write-Host 'goimports found' -ForegroundColor Green
    }
    else {
        Write-Host 'Installing goimports...' -ForegroundColor Cyan
        go install golang.org/x/tools/cmd/goimports@latest
        Write-Host 'goimports installed' -ForegroundColor Green
    }

    Write-Host ''
    Write-Host 'Setup complete!' -ForegroundColor Green
    Write-Host ''
    Write-Host 'Next steps:'
    Write-Host '1. Start Docker Desktop'
    Write-Host '2. Run: make docker-up'
    Write-Host '3. Run: make run'
    Write-Host ''
    Write-Host 'For development with hot reload:'
    Write-Host '  make dev'
}

Invoke-WindowsSetup
