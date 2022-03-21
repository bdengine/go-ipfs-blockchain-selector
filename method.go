package selector

import (
	"github.com/bdengine/go-ipfs-blockchain-standard/dto"
	"github.com/bdengine/go-ipfs-blockchain-standard/model"
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

func GetChallenge() (string, error) {
	return api.GetChallenge()
}

func GetUserCode() (string, error) {
	return api.GetUserCode()
}

func GetPeer(id string) (model.CorePeer, error) {
	return api.GetPeer(id)
}

func UpdateAddress(addrList []string) error {
	return api.UpdateAddress(addrList)
}

func Heartbeat() error {
	return api.Heartbeat()
}

func Mining(m dto.MiningDTO) error {
	return api.Mining(m)
}

func UpdateOrGen(m dto.StoreProofDTO) error {
	return api.UpdateOrGen(m)
}

func Prove(m dto.SVProofDTO) error {
	return api.Prove(m)
}

func GetChallengeStage() (int64,string,int64,error) {
	return api.GetChallengeStage()
}
func GetStoreChallenge() (string,error) {
	return api.GetStoreChallenge()
}

func GetFileList(n int64) ([]string, error) {
	return api.GetFileList(n)
}

func GetBootStrap() ([]string, error) {
	list, err := GetPeerList(10)
	if err != nil {
		return nil, err
	}
	var blist []string
	for _, peer := range list {
		blist = append(blist, peer.Addresses...)
	}
	return blist, nil
}
