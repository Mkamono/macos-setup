export MY_CONFIG_DIR=$HOME/macos-setup

cd $MY_CONFIG_DIR && git --no-pager diff && git --no-pager status -s && cd -
