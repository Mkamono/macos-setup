SHELL=/bin/zsh

.PHONY: init
init: /opt/homebrew/bin/brew
	@source ~/.zprofile && cd ./.config && brew bundle || true && task init

/opt/homebrew/bin/brew:
	@./brew.sh
