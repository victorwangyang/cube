#!/bin/bash


echo "do building......."
go build -o ./bin/do ./do/*.go

echo "master building......."
go build -o ./bin/master ./cluster/master/master.go

echo "node building......."
go build -o ./bin/node ./cluster/node/node.go

echo "doserver building......."
go build -o ./bin/doserver ./doserver/doserver.go


echo "ending......."