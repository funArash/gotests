#!/bin/sh 
GO=${GO:-go}
${GO} test -bench=. ./types
