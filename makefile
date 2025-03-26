default:
	@/bin/bash -c "$$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
	@cd ./.config && /opt/homebrew/bin/brew bundle || true

	@mkdir -p ~/.config/fish
	@cp ./.config/fish/config.fish ~/.config/fish/config.fish
	@cp ./.config/fish/fish_variables ~/.config/fish/fish_variables

	@echo /opt/homebrew/bin/fish | sudo tee -a /etc/shells
	@chsh -s /opt/homebrew/bin/fish
	@echo "\n\nNext: task init"
	@/opt/homebrew/bin/fish
