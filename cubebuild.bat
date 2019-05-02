
go build -o ./bin/do.exe ./do/
go build -o ./bin/master.exe ./cluster/master/master.go
go build -o ./bin/node.exe ./cluster/node/node.go
go build -o ./bin/doserver.exe ./doserver/doserver.go

