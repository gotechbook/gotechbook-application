package redis

import (
	"encoding/json"
)

const TbBlockChainKey = "tb_block_chain"

type BlockChainData struct {
	Data []BlockChain `json:"data"`
}

type BlockChain struct {
	Id                string `json:"id"`
	ChainId           string `json:"chainId"`           // 链ID
	Name              string `json:"name"`              // 链名称
	CurrentHeight     int64  `json:"currentHeight"`     // 当前高度
	SynchronizeHeight int64  `json:"synchronizeHeight"` // 同步高度
}

func (l BlockChainData) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l BlockChainData) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}

func (l BlockChain) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l BlockChain) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}
