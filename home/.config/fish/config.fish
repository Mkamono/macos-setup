if status is-interactive
    eval (/opt/homebrew/bin/brew shellenv)
    eval (mise activate fish)
end

# alias

# git
alias g="git"
alias gd="go run ~/macos-setup/pkgs/gd/main.go"
alias gu="git add . && git commit -m 'update' && git push"

# terraform
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
