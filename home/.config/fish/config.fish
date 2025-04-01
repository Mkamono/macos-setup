if status is-interactive
    eval (/opt/homebrew/bin/brew shellenv)
    eval (mise activate fish)
    git -C ~/macos-setup --no-pager diff
    git -C ~/macos-setup ls-files --others --exclude-standard
end

# abbr

# git
abbr -a -- g git
abbr -a -- gd 'go run ~/macos-setup/pkgs/gd/main.go'
abbr -a -- gu 'git add . && git commit -m "update" && git push'

# terraform
abbr -a -- tf terraform
abbr -a -- tfv 'terraform -v'
abbr -a -- tfi 'terraform init'
abbr -a -- tfa 'terraform apply'
abbr -a -- tfaa 'terraform apply -auto-approve'
abbr -a -- tfd 'terraform destroy'
abbr -a -- tff 'terraform fmt'
abbr -a -- tfp 'terraform plan'
abbr -a -- tfs 'terraform state'

# ls
abbr -a -- ll 'ls -l'
abbr -a -- la 'ls -la'
abbr -a -- l 'ls -lA'

# shell
abbr -a -- c clear
abbr -a -- reload 'clear; exec $SHELL -l'

# config
abbr -a -- cfg 'code ~/macos-setup'

# kubernetes
abbr -a -- k kubectl
abbr -a -- kg 'kubectl get'
abbr -a -- kga 'kubectl get --all-namespaces'

# mise
abbr -a -- m mise
abbr -a -- mr 'mise run'

# other
abbr -a -- h history
