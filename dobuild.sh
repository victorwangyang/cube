#!/bin/bash


echo "do building......."
go build -o ./do/do ./do/*.go

echo "master building......."
go build -o ./cluster/master/master ./cluster/master/master.go

echo "node building......."
go build -o ./cluster/node/node ./cluster/node/node.go

echo "doserver building......."
go build -o ./doserver/doserver ./doserver/doserver.go


echo "ending......."