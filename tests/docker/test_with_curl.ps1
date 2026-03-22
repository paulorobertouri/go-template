$ErrorActionPreference = "Stop"

$image = "go-template-e2e"
$container = "go-template-e2e-container"

try {
  docker build -f docker/build.Dockerfile -t $image .
  $containerId = docker run -d --name $container -p 19000:8000 $image
  Start-Sleep -Seconds 2

  Invoke-WebRequest -UseBasicParsing http://127.0.0.1:19000/v1/public | Out-Null
  Invoke-WebRequest -UseBasicParsing http://127.0.0.1:19000/v1/customer | Out-Null

  $login = Invoke-WebRequest -UseBasicParsing http://127.0.0.1:19000/v1/auth/login
  $headerToken = $login.Headers["X-JWT-Token"]
  if (-not $headerToken) { throw "X-JWT-Token header was not returned" }

  $token = ($login.Content | ConvertFrom-Json).token
  if (-not $token) { throw "JWT token not found in response body" }

  Invoke-WebRequest -UseBasicParsing -Headers @{ Authorization = "Bearer $token" } http://127.0.0.1:19000/v1/private | Out-Null

  Write-Host "All Docker+curl tests passed!"
}
finally {
  docker rm -f $container 2>$null | Out-Null
}
