# Install env for go
download package for linux : https://golang.org/dl/

```cd downloads
tar -C /usr/local -zxvf go1.11.1.linux-amd64.tar.gz
tar -cpvzf go1.11.1.linux-amd64.tar.gz
```

Add to .zshrc
```
#Add go path for zsh
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$(go env GOPATH)/bin
```

If don't have .zshrc
```
#add go path for bashrc
.bashrcinstall env for go
```

Create tree:
```
cd ( home/username)
mkdir -p go/{bin,doc,pkg,src}
cd go/src
```
Test: create main.go




```golang
package main

import "fmt"

func main() {
  fmt.Printf("hello, world\n")
}
```

```
go build main.go
```

Compile excecute file

```
./main
```

Only print result, not excecute file

```
go run main.go
```


# Install golint: 
```
https://github.com/golang/lint
```
Run
```
go get -u golang.org/x/lint/golint
```

# Notes for GIt
Up | Down
-----------|-------------
git add <filename> | git reset HEAD <filename>
git commit ... | git reset @~1 (git reset --soft/hard @~1)
git commit ... | git reset --hard @~1 (remove)

Somethings need to remember

Command | Meaning
-----------|-------------
git reset --hard origin/master | reset to the same remote
git commit --amend -m "xyz" | rename the message of the newest commit 
git commit --all --amend --no-edit | add all files edited into newest commit (=git add + git commit) 

If rebase cannot do (ex) => cherrypick

git rebase -i --root



