param(
    [switch]$Build,
    [switch]$Restart
)

$ErrorActionPreference = "Stop"

$Root = Resolve-Path (Join-Path $PSScriptRoot "..")
$Bin = Join-Path $Root "bin\db-api.exe"
$LogDir = Join-Path $Root "logs"
$Stamp = Get-Date -Format "yyyyMMdd_HHmmss"

New-Item -ItemType Directory -Force -Path (Split-Path $Bin) | Out-Null
New-Item -ItemType Directory -Force -Path $LogDir | Out-Null

if ($Restart) {
    Get-Process -Name "db-api" -ErrorAction SilentlyContinue | Stop-Process -Force
}

if ($Build -or -not (Test-Path $Bin)) {
    Push-Location $Root
    try {
        go build -o $Bin ./cmd/db-api
    } finally {
        Pop-Location
    }
}

$OutLog = Join-Path $LogDir "db-api-$Stamp.out.log"
$ErrLog = Join-Path $LogDir "db-api-$Stamp.err.log"

Start-Process -FilePath $Bin -WorkingDirectory $Root -RedirectStandardOutput $OutLog -RedirectStandardError $ErrLog -WindowStyle Hidden
Write-Host "db-api started"
Write-Host "stdout: $OutLog"
Write-Host "stderr: $ErrLog"
