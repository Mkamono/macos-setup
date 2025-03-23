SHELL=/bin/zsh

.PHONY: init
init: /opt/homebrew/bin/brew
	@source ~/.zprofile && cd ./config && brew bundle && task init

/opt/homebrew/bin/brew:
	@./brew.sh
