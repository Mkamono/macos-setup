export LOAD_SCRIPT=true
export MY_CONFIG_DIR=$HOME/macos-setup
export XDG_CONFIG_HOME=$HOME/.config # ghostty config path

cd $MY_CONFIG_DIR && \
    git pull > /dev/null & \
    git --no-pager diff && \
    git --no-pager status -s && \
    cd - > /dev/null
