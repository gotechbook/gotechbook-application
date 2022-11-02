package mongo

// TbFFFTransactionsName Transactions 3fTransactions信息表
const TbFFFTransactionsName = "tb_fff_transactions"

type TbFFFTransactions struct {
	TransactionHash   string `json:"transactionHash" bson:"transactionHash"`
	BlockHash         string `json:"blockHash"  bson:"blockHash"`
	BlockNumber       uint64 `json:"blockNumber" bson:"blockNumber"`
	From              string `json:"from" bson:"from"`
	To                string `json:"to" bson:"to"`
	Value             string `json:"value" bson:"value"`
	Gas               uint64 `json:"gas" bson:"gas"`
	GasPrice          uint64 `json:"gasPrice" bson:"gasPrice"`
	Nonce             uint64 `json:"nonce" bson:"nonce"`
	Data              string `json:"data" bson:"data"`
	Size              string `json:"size,omitempty" bson:"size"`
	Status            uint64 `json:"status" bson:"status"`
	Type              uint8  `json:"type" bson:"type"`
	PostState         []byte `json:"postState" bson:"postState"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	GasUsed           uint64 `json:"gasUsed" bson:"gasUsed"`
	TransactionIndex  uint   `json:"transactionIndex" bson:"transactionIndex"`
	ContractAddress   string `json:"contractAddress" bson:"contractAddress"`
}
