export LOAD_ALIAS=true
export MY_CONFIG_DIR=$HOME/macos-setup

alias ll="ls -l"
alias la="ls -la"
alias dc="docker compose"
alias tf="terraform"
alias gd="task -g git:gd"
alias gpr="task -g git:gpr"
alias gu="git add . && git commit -m 'update' && git push"
alias zshconfig="code ~/.zshrc"
alias myconfig="code $MY_CONFIG_DIR"
alias relogin="exec $SHELL -l && clear"
