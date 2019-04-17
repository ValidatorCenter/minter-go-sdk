#!/bin/sh

# install: EasyJSON
#go get -u github.com/mailru/easyjson/...

easyjson -all apiEstimateCoin.go
easyjson -all apiGetAddress.go
easyjson -all apiGetBlock.go
easyjson -all apiGetCandidate.go
easyjson -all apiGetCandidates.go
easyjson -all apiGetCoinInfo.go
easyjson -all apiGetEvents.go
easyjson -all apiGetMinGas.go
easyjson -all apiGetStatus.go
easyjson -all apiGetTransaction.go
easyjson -all apiGetValidators.go
easyjson -all apiSetTransaction.go