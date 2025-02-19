.PHONY: check-coverage
check-coverage:
	go test ./... -coverprofile=coverage.out
	@echo "==================== Running Coverage Report ===================="
	@go tool cover -func=coverage.out | tee coverage-summary.txt
	@go tool cover -html=coverage.out -o coverage.html
	@echo ""
	@echo "✅ Coverage report generated successfully!"
	@printf "📊 Overall Coverage: \033[1;32m%s\033[0m\n" "$$(grep 'total:' coverage-summary.txt | awk '{print $$3}')"
	@printf "🌐 Open the report in your browser: \033[1;34mfile://$(PWD)/coverage.html\033[0m\n"
	@echo "==============================================================="
	@rm coverage-summary.txt