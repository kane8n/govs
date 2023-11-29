.PHONY: install

install:
	@echo "Installing..."
	@mkdir -p ~/.govs/bin
	@go mod tidy
	@go build -ldflags="-s -w" -trimpath -o ~/.govs/bin/_govs_prompt .
	@cp -p ./run.sh ~/.govs/bin/_govs
	@echo "Done!"
	@echo "Add the following to your .bashrc or .zshrc config"
	@echo "    alias govs=\"source _govs\""
	@echo "    export PATH=\$$PATH:~/.govs/bin"
