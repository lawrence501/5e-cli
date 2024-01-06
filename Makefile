run:
	go run cmd/5e-cli/*.go

session:
	open -a "Google Chrome" https://5e.tools/classes.html
	open -a "Google Chrome" https://5e.tools/bestiary.html
	open -a "Google Chrome" https://5e.tools/spells.html

lint:
	golangci-lint run