#!/bin/bash
goxc -prefix github.com/openyard/ebics/2.5 -schema xsd/h004/ebics_H004.xsd
goxc -prefix github.com/openyard/ebics/2.5 -schema xsd/h004/ebics_hev.xsd
goxc -prefix github.com/openyard/ebics/3.0 -schema xsd/h005/ebics_H005.xsd
goxc -prefix github.com/openyard/ebics/3.0 -schema xsd/h005/ebics_hev.xsd
