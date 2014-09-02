# install Go

安装可以去[官网][1]下载对应的 package 或者使用 brew 直接在 command line 里进行安装配置，这里介绍 brew 安装方式。

## Step 1: ``vi .zshrc``

```
# for golang
# mkdir $HOME/go
# mkdir -p $GOPATH/src/github.com/user
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

记得 ``source .zshrc``！

## Step 2: install & configuration

```
brew install go
mkdir $HOME/go
mkdir -p $GOPATH/src/github.com/user
```

### 参考

[install go (in OS X)][2]

[1]: http://golang.org/doc/install
[2]: https://gist.github.com/fyears/5607418