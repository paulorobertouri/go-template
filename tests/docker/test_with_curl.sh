set -euo pipefail

# Build production Docker image
docker build -f docker/build.Dockerfile -t go-template:latest .

# Start container in background
container_id=$(docker run -d -p 19000:8000 go-template:latest)
trap "docker rm -f $container_id" EXIT

# Wait for server to be ready
sleep 2

# Test public endpoint
curl -fsS http://127.0.0.1:19000/v1/public > /dev/null

# Test customer endpoint
curl -fsS http://127.0.0.1:19000/v1/customer > /dev/null

# Test login and extract JWT token from response body
raw_headers=$(curl -isS http://127.0.0.1:19000/v1/auth/login)
if ! echo "$raw_headers" | grep -iq '^x-jwt-token:'; then
  echo "X-JWT-Token header was not returned"
  exit 1
fi

token=$(curl -fsS http://127.0.0.1:19000/v1/auth/login | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
if [ -z "$token" ]; then
  echo "JWT token not found in response"
  exit 1
fi

# Test private endpoint with JWT
curl -fsS -H "Authorization: Bearer ${token}" http://127.0.0.1:19000/v1/private > /dev/null

echo "All Docker+curl tests passed!"
