module github.com/ipfs/go-ipfs-auth/selector

go 1.15

require (
	github.com/bdengine/go-ipfs-blockchain-standard v0.0.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/bdengine/go-ipfs-blockchain-eth v0.0.2
)

replace (
	github.com/bdengine/go-ipfs-blockchain-standard => ../standard
	github.com/bdengine/go-ipfs-blockchain-eth => ../auth-source-eth
)
