SHELL=/bin/zsh

.PHONY: init
init: /opt/homebrew/bin/brew
	@source ~/.zprofile && brew bundle && task init

/opt/homebrew/bin/brew:
	@./brew.sh
