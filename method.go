package selector

import (
	"github.com/ipfs/go-ipfs-auth/standard"
)

func NewPeer(pid string) (*standard.Peer, error) {
	return authApi.NewPeer(pid)
}

func ResponseApply(cid string, pid string) (*standard.Apply, error) {

	return authApi.ResponseApply(cid, pid)
}

func GetServerList() ([]standard.CorePeer, error) {
	return authApi.GetServerList()
}

func ApplyJoinServerNet(addresses []string) error {
	return authApi.ApplyJoinServerNet(addresses)
}

func AuthJoinApply(pid string, i int) (standard.JoinServerListApply, error) {
	return authApi.AuthJoinApply(pid, i)
}

func JoinServerNet() ([]standard.Peer, error) {
	return authApi.JoinServerNet()
}

func GetApply(peerId string) ([]standard.JoinServerListApply, error) {
	return authApi.GetApply(peerId)
}

// 获取文件权限信息
func FindFileAuth(cid string) (*standard.Apply, error) {
	return authApi.FindFileAuth(cid)
}

// 获取已拥有文件列表
func FindFileList(pid string) ([]string, error) {
	return authApi.FindFileList(pid)
}

// 获取文件阅读者列表
func FindFileReaderList(cid string) ([]string, error) {
	return authApi.FindFileReaderList(cid)
}

// 获取节点阅读历史
func FindPeerReadingHistory() ([]string, error) {
	return authApi.FindPeerReadingHistory()
}

// NewIpfs 新增ipfs文件权限信息
func NewIpfsFilePermission(cid string, state int) (*standard.IpfsFileInfo, error) {
	return authApi.NewIpfsFilePermission(cid, state)
}

// Share 分享文件
func Share(cid string, pid string, limit standard.Limit) (*standard.IpfsFileInfo, error) {
	return authApi.Share(cid, pid, limit)
}

func Change(cid string, ipfs standard.IpfsFileInfo) (*standard.IpfsFileInfo, error) {
	return authApi.Change(cid, ipfs)
}

func Delete(cid string) error {
	return authApi.Delete(cid)
}

func ApplyIpfsRemote(cid string) (*standard.Apply, error) {
	return authApi.ApplyIpfsRemote(cid)
}

func ApplyIpfsLocal(cid string) (string, error) {
	return authApi.ApplyIpfsLocal(cid)
}

// ParseState 解析状态
func ParseState(state uint8) (readType uint8, shareType uint8, noticeType uint8, authType uint8, err error) {
	authType = state & 3
	if authType == 3 {
		err = errWrongAuthType
		return
	}
	state >>= 2
	noticeType = state & 3
	state >>= 2
	shareType = state & 3
	state >>= 2
	readType = state & 3
	return
}

func GetState(readType uint8, shareType uint8, noticeType uint8, authType uint8) (uint8, error) {
	if readType > 3 || shareType > 3 || noticeType > 3 || authType > 2 {
		return 0, errWrongState
	}
	return authType + noticeType<<2 + shareType<<4 + readType<<6, nil
}
