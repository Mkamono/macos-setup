# # echo $0
# # echo $SHELL
# # readlink "/proc/$$/exe"

alias ll="ls -l"
alias la="ls -la"
alias dc="docker-compose"
alias tf="terraform"
alias gd="task -g git:gd"
alias gpr="task -g git:gpr"
alias gr="task -g git:gr"
alias gu="git add . && git commit -m 'update' && git push"
alias zshconfig="code ~/.zshrc"
alias relogin="clear && zsh -c 'source ~/.zshrc'"

alias myconfig="code ~/macos-setup"

export ZSH="$HOME/.local/share/sheldon/repos/github.com/ohmyzsh/ohmyzsh"

# See https://github.com/ohmyzsh/ohmyzsh/wiki/Themes
ZSH_THEME="jonathan"

eval "$(mise activate zsh)"
eval $(sheldon source)

eval "$(/opt/homebrew/bin/brew shellenv)"

source <(minikube completion zsh)
eval "$(task --completion zsh)"

source "/Users/makotokamono/.local/share/sheldon/repos/github.com/ohmyzsh/ohmyzsh/oh-my-zsh.sh"
source "/Users/makotokamono/.local/share/sheldon/repos/github.com/nnao45/zsh-kubectl-completion/zsh-kubectl-completion.plugin.zsh"
source "/Users/makotokamono/.local/share/sheldon/repos/github.com/marlonrichert/zsh-autocomplete/zsh-autocomplete.plugin.zsh"
source "/Users/makotokamono/.local/share/sheldon/repos/github.com/zsh-users/zsh-syntax-highlighting/zsh-syntax-highlighting.plugin.zsh"
