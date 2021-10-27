module github.com/ipfs/go-ipfs-auth/selector

go 1.15

require (
	github.com/ipfs/go-ipfs-auth/standard v0.0.0
	github.com/ipfs/go-ipfs-auth/auth-source-fabric v0.0.0
)

replace (
	github.com/ipfs/go-ipfs-auth/standard v0.0.0 => ../standard
	github.com/ipfs/go-ipfs-auth/auth-source-fabric v0.0.0 => ../auth-source-fabric
)
