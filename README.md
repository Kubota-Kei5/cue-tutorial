--Setup for Golang--<br>
OS：Ubuntu22.04.5（WSL2）<br>
install command：
```
sudo snap install --classic go
```
<br>
--Setup for cuelang--<br>

1.Goを使ってcueをインストール
```
go install cuelang.org/go/cmd/cue@latest
```

2.pathの追加
```
export PATH=$PATH:$(go env GOPATH)/bin
```
