compile:
	echo "Compiling for Linux, Mac and Windows"
	GOOS=linux GOARCH=amd64 go build -o bin/redis-linux-64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/redis-mac-64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/redis-win-64.exe main.go
	cp .env.example bin/.env