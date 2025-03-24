#!/bin/zsh

# alias

# git
alias g="git"
alias gd="go run ~/macos-setup/pkgs/gd/main.go"
alias gu="git add . && git commit -m 'update' && git push"

# terraform
alias terraform="docker run -it --rm -v $(pwd):/workspace -w /workspace hashicorp/terraform:light"
alias tf="terraform"
alias tfv="terraform -v"
alias tfi="terraform init"
alias tfa="terraform apply"
alias tfaa="terraform apply -auto-approve"
alias tfd="terraform destroy"
alias tff="terraform fmt"
alias tfp="terraform plan"
alias tfs="terraform state"

# ls
alias ll="ls -l"
alias la="ls -la"
alias l="ls -lA"

# shell
alias c="clear"
alias reload="clear; exec $SHELL -l"

# config
alias cfg="code ~/macos-setup"

eval "$(mise activate zsh)"
eval "$(sheldon source)"

# To customize prompt, run `p10k configure` or edit ~/.p10k.zsh.
[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh
