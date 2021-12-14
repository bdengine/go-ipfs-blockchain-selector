package selector

import (
	"github.com/ipfs/go-ipfs-auth/standard/model"
)

func GetPeerList(num int) ([]model.CorePeer, error) {
	// todo 缓存？
	/*get, exist := api.goCache.Get(peerListKey)
	if !exist {
		list, err := api.GetPeerList(num)
		if err != nil {
			return nil, err
		}
		err = api.goCache.Add(peerListKey, list, cacheExpire)
		return list, err
	}*/
	return api.GetPeerList(num)
}

func AddFile(info model.IpfsFileInfo) error {
	return api.AddFile(info)
}

func DeleteFile(info string) error {
	return api.DeleteFile(info)
}

func RechargeFile(cid string, days int64) error {
	return api.RechargeFile(cid, days)
}

func InitPeer(peer model.CorePeer) error {
	return api.InitPeer(peer)
}

func ReportContribute(num int64) error {
	return nil
}

func GetUserCode() (string, error) {
	return api.GetUserCode()
}

func GetPeer(id string) (model.CorePeer, error) {
	return api.GetPeer(id)
}

func GetChallenge() (string, error) {
	return api.GetChallenge()
}

func Mining(m []model.IpfsMining) error {
	return api.Mining(m)
}

func GetBootStrap() ([]string, error) {
	boot, exist := api.goCache.Get(bootstrapListKey)
	if !exist {
		list, err := GetPeerList(10)
		if err != nil {
			return nil, err
		}
		var blist []string
		for _, peer := range list {
			blist = append(blist, peer.Addresses...)
		}
		err = api.goCache.Add(bootstrapListKey, blist, cacheExpire)
		return blist, nil
	}
	return boot.([]string), nil
}
