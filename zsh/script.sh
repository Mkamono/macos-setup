cd $REPO_DIR && \
    task karabiner:diff && \
    cd - > /dev/null

cd $REPO_DIR && \
    git pull && \
    git --no-pager diff && \
    git --no-pager status -s && \
    cd - > /dev/null

export GOPATH="$HOME/go"
export PATH="$HOME/go/bin:$PATH"
