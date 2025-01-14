cd $REPO_DIR && \
    git pull && \
    git --no-pager diff && \
    git --no-pager status -s && \
    cd - > /dev/null
