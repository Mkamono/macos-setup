#!/bin/zsh

REPO_DIR=$HOME/macos-setup

# 基本設定の読み込み
source $REPO_DIR/zsh/alias.sh

# 機能別スクリプトの実行
source "$REPO_DIR/zsh/path_config.sh"

# プラグインと補完の設定
source /opt/homebrew/share/zsh-autocomplete/zsh-autocomplete.plugin.zsh
source /opt/homebrew/share/google-cloud-sdk/completion.zsh.inc

# sync
task -g --parallel --silent pull karabiner:sync vscode:sync
