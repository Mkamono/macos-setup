#!/bin/zsh

# プラグインと補完の設定
# source /opt/homebrew/share/zsh-autocomplete/zsh-autocomplete.plugin.zsh
# source /opt/homebrew/share/google-cloud-sdk/completion.zsh.inc
source <(kubectl completion zsh)
source <(minikube completion zsh)

# sync
task -g --parallel --silent pull karabiner:sync vscode:sync
