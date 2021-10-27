package selector

import (
	"github.com/ipfs/go-ipfs-auth/auth-source-fabric/api"
	"github.com/ipfs/go-ipfs-auth/standard"
)

const (
	AuthWayFabric = "fabric"
)

const (
	RoleCore = "core"
	RoleLite = "lite"
)

var defaultAuthWay, defaultRole = AuthWayFabric, RoleLite

var authApi standard.AuthAPI

func Initialization(url string, authWay string, role string) error {
	if authWay == "" {
		authWay = defaultAuthWay
	} else {
		defaultAuthWay = authWay
	}
	if role == "" {
		role = defaultRole
	} else {
		defaultRole = role
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
