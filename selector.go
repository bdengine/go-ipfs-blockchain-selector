package selector

import (
	"github.com/ipfs/go-ipfs-auth/auth-source-fabric/api"
	standard "github.com/ipfs/go-ipfs-auth/standard/interface"
)

const (
	AuthWayFabric = "fabric"
)

const (
	RoleCore = "core"
	RoleLite = "lite"
)

var DefaultAuthWay, DefaultRole = AuthWayFabric, RoleLite

var authApi standard.AuthAPI

func Initialization(url string, authWay string, role string) error {
	if authWay == "" {
		authWay = DefaultAuthWay
	} else {
		DefaultAuthWay = authWay
	}
	if role == "" {
		role = DefaultRole
	} else {
		DefaultRole = role
	}

	switch authWay {
	case AuthWayFabric:
		switch role {
		case RoleCore:
			authApi = api.CoreAPI{}
		case RoleLite:
			authApi = api.LiteAPI{}
		default:
			return errUnknownRole
		}
		return api.Initialization(url)
	default:
		return errUnknownAuthWay
	}
}
