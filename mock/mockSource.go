package mock

import (
	"fmt"
	"github.com/bdengine/go-ipfs-blockchain-standard/model"
)

type MockPeerApi struct{}

func (p *MockPeerApi) ApplyLocal(cid string) error {
	return nil
}

func (p *MockPeerApi) ReportContribute(num int64) error {
	return nil
}

func (p *MockPeerApi) ApplyRemote(cid string) error {
	return nil
}

var cidMap = map[string]bool{}

func (p *MockPeerApi) AddFile(info model.IpfsFileInfo) error {
	f, _ := cidMap[info.Cid]
	if f {
		return fmt.Errorf("已存在该文件")
	} else {
		cidMap[info.Cid] = true
	}
	return nil
}

func (p *MockPeerApi) DeleteFile(cid string) error {
	delete(cidMap, cid)
	return nil
}

func (p *MockPeerApi) RechargeFile(cid string, days int64) error {
	return nil
}

func (p *MockPeerApi) GetPeerList(num int) ([]model.CorePeer, error) {
	peers := []model.CorePeer{{
		PeerId:    "12D3KooWLUTBUkLnfdcJbyV1C7ZRsdFcDoN89mmYknFH5ef9pTyM",
		Addresses: []string{"/ip4/172.19.17.12/tcp/4001/p2p/12D3KooWDaJ9JJezdx8vMJxxVEsysYBtcHdxUBoogGypQXfrqiEj"},
	}, {
		PeerId:    "12D3KooWLUTBUkLnfdcJbyV1C7ZRsdFcDoN89mmYknFH5ef9pTyM",
		Addresses: []string{"/ip4/192.168.70.25/tcp/4001/p2p/12D3KooWLUTBUkLnfdcJbyV1C7ZRsdFcDoN89mmYknFH5ef9pTyM"},
	}}
	return peers, nil
}

func (p *MockPeerApi) GetBootStrap() ([]string, error) {
	list, err := p.GetPeerList(10)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, peer := range list {
		res = append(res, peer.Addresses...)
	}
	return res, nil
}

func (p *MockPeerApi) GetUserCode() (string, error) {
	return "", nil
}

func (p *MockPeerApi) GetPeer(id string) (model.CorePeer, error) {
	return model.CorePeer{}, nil
}

func (p *MockPeerApi) InitPeer(peer model.CorePeer) error {
	return nil
}

func (p *MockPeerApi) DaemonPeer() error {
	return nil
}

func (p *MockPeerApi) GetChallenge() (string, error) {
	return "", nil
}

func (p *MockPeerApi) Mining(model.IpfsMining) error {
	return nil
}

func (p *MockPeerApi) UpdateAddress(addrList []string) error {
	return nil
}

func (p *MockPeerApi) Heartbeat() error {
	return nil
}

func (p *MockPeerApi) GetFileList(n int64) ([]string, error) {
	return nil, nil
}
