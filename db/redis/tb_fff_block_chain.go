package redis

import (
	"encoding/json"
)

const TbBlockChainKey = "tb_fff_block_chain"

type FFFBlockChain struct {
	ChainId           string `json:"chainId"`           // 链ID
	Name              string `json:"name"`              // 链名称
	CurrentHeight     uint64 `json:"currentHeight"`     // 当前高度
	SynchronizeHeight uint64 `json:"synchronizeHeight"` // 同步高度
}

func (l FFFBlockChain) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l FFFBlockChain) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}
