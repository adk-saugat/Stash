build:
	cd client && go build -o ../stash

run:
	cd server/cmd/server && go run main.go