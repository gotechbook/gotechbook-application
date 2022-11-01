package mongo

// TbFFFBlocksName Blocks 3f区块信息表
const TbFFFBlocksName = "tb_fff_blocks"

type TbFFFBlocks struct {
	Number           int64       `json:"number,omitempty" bson:"number"`
	Difficulty       string      `json:"difficulty,omitempty" bson:"difficulty"`
	ExtraData        string      `json:"extraData,omitempty" bson:"extraData"`
	GasLimit         int64       `json:"gasLimit,omitempty" bson:"gasLimit"`
	GasUsed          int64       `json:"gasUsed,omitempty" bson:"gasUsed"`
	Hash             string      `json:"hash,omitempty" bson:"hash"`
	LogsBloom        string      `json:"logsBloom,omitempty" bson:"logsBloom"`
	Miner            string      `json:"miner,omitempty" bson:"miner"`
	MixHash          string      `json:"mixHash,omitempty" bson:"mixHash"`
	Nonce            string      `json:"nonce,omitempty" bson:"nonce"`
	ParentHash       string      `json:"parentHash,omitempty" bson:"parentHash"`
	ReceiptsRoot     string      `json:"receiptsRoot,omitempty" bson:"receiptsRoot"`
	Sha3Uncles       string      `json:"sha3Uncles,omitempty" bson:"sha3Uncles"`
	Size             int64       `json:"size,omitempty" bson:"size"`
	StateRoot        string      `json:"stateRoot,omitempty" bson:"stateRoot"`
	Timestamp        int64       `json:"timestamp,omitempty" bson:"timestamp"`
	TotalDifficulty  int64       `json:"totalDifficulty,omitempty" bson:"totalDifficulty"`
	Transactions     []string    `json:"transactions,omitempty" bson:"transactions"`
	TransactionsRoot string      `json:"transactionsRoot,omitempty" bson:"transactionsRoot"`
	Uncles           interface{} `json:"uncles,omitempty" bson:"uncles"`
}
