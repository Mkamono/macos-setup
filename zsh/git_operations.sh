#!/bin/zsh

# Git操作（6時間に1回のみpullを実行）
LAST_PULL_FILE="$HOME/.last_git_pull"
CURRENT_TIME=$(date +%s)
SHOULD_PULL=false

if [ ! -f "$LAST_PULL_FILE" ]; then
    SHOULD_PULL=true
else
    LAST_PULL_TIME=$(cat "$LAST_PULL_FILE")
    TIME_DIFF=$((CURRENT_TIME - LAST_PULL_TIME))
    # 6時間（21600秒）経過しているかチェック
    if [ $TIME_DIFF -ge 21600 ]; then
        SHOULD_PULL=true
    fi
fi

cd $REPO_DIR
if [ "$SHOULD_PULL" = true ]; then
    git pull
    echo $CURRENT_TIME > "$LAST_PULL_FILE"
fi
git --no-pager diff
git --no-pager status -s
cd - > /dev/null
