module github.com/ipfs/go-ipfs-auth/selector

go 1.15

require (
	github.com/ipfs/go-ipfs-auth/standard v0.0.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/ipfs/go-ipfs-auth/auth-source-eth v0.0.0
)

replace (
	github.com/ipfs/go-ipfs-auth/standard => ../standard
	github.com/ipfs/go-ipfs-auth/auth-source-eth => ../auth-source-eth
)
