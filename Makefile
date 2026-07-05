build:
	go build -o 5e-cli ./cmd/5e-cli

run:
	go run ./cmd/5e-cli $(ARGS)

session:
	open -a "Google Chrome" https://5e.tools/classes.html
	open -a "Google Chrome" https://5e.tools/bestiary.html
	open -a "Google Chrome" https://5e.tools/spells.html

lint:
	golangci-lint run