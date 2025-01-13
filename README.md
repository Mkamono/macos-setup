# Setup Mac OS

データのバックアップはこちらから

[backup.md](./docs/backup.md)

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
brew bundle --file=~/macos-setup/Brewfile
task -g brew:bundle
```

## Setup Zsh, etc... (via task)

```bash
cd ~/macos-setup
task init
```
