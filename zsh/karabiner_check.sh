#!/bin/zsh

# Karabinerの設定差分チェック
cd $REPO_DIR && \
    task karabiner:diff && \
    cd - > /dev/null
