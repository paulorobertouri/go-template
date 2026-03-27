# PowerShell script to rename the Go template project
# Usage: ./rename-template.ps1 NewProjectName
param(
    [Parameter(Mandatory = $true)]
    [string]$NewName
)

# Replace all occurrences of 'go-template' and 'Go Template' in all files
Get-ChildItem -Path . -Recurse -File | ForEach-Object {
    (Get-Content $_.FullName) -replace 'go-template', $NewName -replace 'Go Template', $NewName | Set-Content $_.FullName
}

Write-Host "Renamed project to $NewName. Please update go.mod manually if needed."
