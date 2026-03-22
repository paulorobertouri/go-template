# PowerShell script for Windows
$ScriptDirectory = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location "$ScriptDirectory\..\.."

& .\tests\docker\test_with_curl.ps1
