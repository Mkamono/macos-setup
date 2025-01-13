# Setup Mac OS

## Install Homebrew

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

その後のインストラクションに従ってPATHを通す

## Clone Repository

```bash
cd ~
git config --global user.name <username>
git config --global user.email <email>
git clone https://github.com/Mkamono/macos-setup.git
```

## Install Brew Packages

```bash
cd ~
cp macos-setup/Brewfile ~/Brewfile
brew bundle --global
rm ~/Brewfile
```

## Setup Zsh, etc... (via task)

```bash
cd ~/macos-setup
task init
```
