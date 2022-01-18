package selector

import (
	"fmt"
	"github.com/bdengine/go-ipfs-blockchain-eth/implement"
	_interface "github.com/bdengine/go-ipfs-blockchain-standard/interface"
	"github.com/ipfs/go-ipfs-auth/selector/mock"
	"github.com/patrickmn/go-cache"
	"time"
)

var errInitFail error = fmt.Errorf("权限模块初始化失败")

type BlockchainAPI struct {
	_interface.Peer
	goCache *cache.Cache
}

var api *BlockchainAPI

const (
	cacheExpire = 2 * time.Hour
	cacheClean  = 1 * time.Hour
)

const (
	sourceFabric = "fabric"
	sourceCosmos = "cosmos"
	sourceMock   = "mock"
)

func Daemon(configRoot string, source string, peerId string) (*BlockchainAPI, []string, error) {
	// todo 依赖注入
	switch source {
	case sourceFabric:
		return nil, nil, fmt.Errorf("不支持的source：%s", source)
	case sourceCosmos:
		apiImplement, err := implement.NewApi(configRoot, peerId)
		if err != nil {
			return nil, nil, err
		}
		api = &BlockchainAPI{
			Peer:    apiImplement,
			goCache: cache.New(cacheExpire, cacheClean),
		}
	case sourceMock:
		api = &BlockchainAPI{
			Peer:    &mock.MockPeerApi{},
			goCache: cache.New(cacheExpire, cacheClean),
		}
	default:
		return nil, nil, fmt.Errorf("不支持的source：%s", source)
	}

	err := api.DaemonPeer()
	if err != nil {
		return nil, nil, err
	}
	peerAddrList, err := GetBootStrap()
	if err != nil {
		return nil, nil, err
	}
	return api, peerAddrList, nil
}

func GetApiInfo() (*BlockchainAPI, error) {
	if api == nil {
		return nil, errInitFail
	}
	return api, errInitFail
}
