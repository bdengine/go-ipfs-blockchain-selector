module github.com/bdengine/go-ipfs-blockchain-eth

go 1.15

require (
	github.com/bdengine/go-ipfs-blockchain-standard v0.0.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/bdengine/go-ipfs-blockchain-eth v0.0.1
)

replace (
	github.com/bdengine/go-ipfs-blockchain-standard => ../go-ipfs-blockchain-standard
	github.com/bdengine/go-ipfs-blockchain-eth => ../go-ipfs-blockchain-eth
)
