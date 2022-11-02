package mongo

// TbFFFBlocksName Blocks 3f区块信息表
const TbFFFBlocksName = "tb_fff_blocks"

type TbFFFBlocks struct {
	Number           uint64      `json:"number,omitempty" bson:"number"`
	Difficulty       string      `json:"difficulty,omitempty" bson:"difficulty"`
	ExtraData        string      `json:"extraData,omitempty" bson:"extraData"`
	GasLimit         uint64      `json:"gasLimit,omitempty" bson:"gasLimit"`
	GasUsed          uint64      `json:"gasUsed,omitempty" bson:"gasUsed"`
	Hash             string      `json:"hash,omitempty" bson:"hash"`
	LogsBloom        string      `json:"logsBloom,omitempty" bson:"logsBloom"`
	Miner            string      `json:"miner,omitempty" bson:"miner"`
	MixHash          string      `json:"mixHash,omitempty" bson:"mixHash"`
	Nonce            uint64      `json:"nonce,omitempty" bson:"nonce"`
	ParentHash       string      `json:"parentHash,omitempty" bson:"parentHash"`
	ReceiptsRoot     string      `json:"receiptsRoot,omitempty" bson:"receiptsRoot"`
	Sha3Uncles       string      `json:"sha3Uncles,omitempty" bson:"sha3Uncles"`
	Size             string      `json:"size,omitempty" bson:"size"`
	StateRoot        string      `json:"stateRoot,omitempty" bson:"stateRoot"`
	Timestamp        uint64      `json:"timestamp,omitempty" bson:"timestamp"`
	Transactions     int         `json:"transactions,omitempty" bson:"transactions"`
	TransactionsRoot string      `json:"transactionsRoot,omitempty" bson:"transactionsRoot"`
	Uncles           interface{} `json:"uncles,omitempty" bson:"uncles"`
}
