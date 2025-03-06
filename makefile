# MacをセットアップするためのMakefile
SHELL := /bin/zsh
DOTFILE_DIR := ~/macos-setup

# mise envでも指定しているので注意
XDG_CONFIG_HOME := $(HOME)/macos-setup/.config

.PHONY: install
install: /opt/homebrew/bin/brew symlink /usr/local/bin/battery
	@source ~/.zshrc
	@echo "Installed"

/opt/homebrew/bin/brew: # Homebrew
	@/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
	@echo "Installed Homebrew"

.PHONY: symlink
symlink: # シンボリックリンクを作成
	@ln -sf $(DOTFILE_DIR)/.zshrc ~/.zshrc
	@mkdir -p ~/.config/mise && ln -sf $(XDG_CONFIG_HOME)/mise/config.toml ~/.config/mise/config.toml
	@mkdir -p ~/.config/ghostty && ln -sf $(XDG_CONFIG_HOME)/ghostty/config ~/.config/ghostty/config
	@echo "Symlinked"

/usr/local/bin/battery: # battery
	@curl -s https://raw.githubusercontent.com/actuallymentor/battery/main/setup.sh | bash
	@battery maintain 80
	@echo "Installed battery"
