# learningGo

# Dowload Go (For linux)  https://golang.org/dl/

Unzip the tar.gz file in the / usr / local directory

# In home modify the .bashrc/.zshrc file

export GOPATH=/home/andrea/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go

export PATH=$PATH:$GOBIN:$GOROOT/bin

# verify that it is correct go

go version
