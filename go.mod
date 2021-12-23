module github.com/datachainlab/go-sdk-sample

go 1.16

require (
	github.com/ethereum/go-ethereum v1.9.25
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/harmony-one/go-sdk v1.2.8
	github.com/harmony-one/harmony v1.10.2-0.20210123081216-6993b9ad0ca1
	github.com/hyperledger-labs/yui-ibc-solidity v0.0.0-20211102033703-b1c507b339f0
	github.com/rs/zerolog v1.21.0 // indirect
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
)

replace (
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/hyperledger/fabric-sdk-go => github.com/datachainlab/fabric-sdk-go v0.0.0-20200730074927-282a61dcd92e
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
