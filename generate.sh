#!/bin/bash
goxc -prefix github.com/openyard/ebics/2.5 -schema xsd/h004/ebics_H004.xsd
goxc -prefix github.com/openyard/ebics/2.5 -schema xsd/h004/ebics_hev.xsd
goxc -prefix github.com/openyard/ebics/3.0 -schema xsd/h005/ebics_H005.xsd
goxc -prefix github.com/openyard/ebics/3.0 -schema xsd/h005/ebics_hev.xsd

rm -rf 2.5 3.0
mv github.com/openyard/ebics/* .
rm -rf github.com

rm 2.5/h004/TimeStamp.go
go fmt ./...
go build ./...
go fmt ./...
