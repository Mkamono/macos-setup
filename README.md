# Setup Mac OS

## Install Homebrew

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

その後のインストラクションに従ってPATHを通す

## Install Brewfile

```bash
brew bundle
```


remote brewfile

local brewfile

差分を検知
remoteにないがlocalにあるもの→アンスコ
依存として使われていないかを確認?
すべてのrmが失敗するまでループ

brew list -1 --formulaeとかでいろいろできそう

remoteにあるがlocalにないもの→インスコ