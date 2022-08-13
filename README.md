# ssh-tools

## 使用方法
### 0. init
`ssh-tools init`

### 1. add machine
`ssh-tools add --name xx --ip 1.1.1.1 --description yyyy --user user --passwd passwd`

### 2. delete machine
`ssh-tools delete --name xx`
`ssh-tools delete --id 1`

### 3. update machine
`ssh-tools update --name`

### 4. list machine
`ssh-tools list`

## 编译
`go build -ldflags "-X 'github.com/zhanglianx111/ssh-tools/cmd.Author=username' 
    -X 'github.com/zhanglianx111/ssh-tools/cmd.GoVersion=`go version`' 
    -X 'github.com/zhanglianx111/ssh-tools/cmd.CommitID=`git log --pretty=format:%h -1`'" main.go
`

