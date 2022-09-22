coverage:
	echo "coverage starting"
	go test -coverprofile coverage.out ./
	go tool cover -html=coverage.out -o coverage.html
	echo "coverage completed"

update:
	echo "update go mod init"
	rm -f go.*
	go mod init github.com/jeffotoni/gmelhorenvio
	go mod tidy
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gmelhorenvio main.go
	echo "buid complete"

build:
	echo "building"
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o gmelhorenvio main.go
	echo "buid complete"

docker-build:
	echo "building container"
	docker build --no-cache -f Dockerfile -t jeffotoni/gmelhorenvio .
	echo "buid complete"

docker-run:
	echo "running container"
	docker run -p 8080:8080 \
	-e MELHORENVIO_CLIENT_ID=$CLIENT_ID \
	-e MELHORENVIO_CLIENT_SECRET=$CLIENT_SECRET \
	-e MELHORENVIO_REDIRECT_URI=$REDIRECT_URI \
	-e API_STATIC_TOKEN=$API_STATIC_TOKEN \
	jeffotoni/gmelhorenvio

copy:
	aws s3 cp ./cmd/credentials/credentials.json s3://${AWS_BUCKET_PRD}
