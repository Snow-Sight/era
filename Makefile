help:
	@echo "Available tasks:"
	@echo "    test              Run go tests"
	@echo "    cover             Run go tests and produce coverage.html"
	@echo "    cover-cleanup     Cleanup coverage files"

test:
	@echo "Running `go test ./...`"
	@go test ./...

cover:
	@echo "Generating go test coverage, file coverage.out and coverage.html will be created"
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Run \"make cover-cleanup\" to remove these files"

cover-cleanup:
	@echo "Cleaning up coverage files coverage.out and coverage.html"
	@rm ./coverage.out
	@rm ./coverage.html