rm ./sauced

env GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p builds/sauced-windows
mv sauced builds/sauced-windows

env GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p builds/sauced-linux
mv sauced builds/sauced-linux


go build -ldflags "-X github.com/mdsauce/sauced/cmd.CurVersion=`git rev-parse --short HEAD`" -o sauced
mkdir -p builds/sauced-mac
mv sauced builds/sauced-mac
