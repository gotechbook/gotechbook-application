package mongo

// TbFFFTransactionsLogsName Transactions 3fTransactionsLogs信息表
const TbFFFTransactionsLogsName = "tb_fff_transactions_logs"

type TbFFFTransactionsLogs struct {
	BlockNumber uint64   `json:"blockNumber" bson:"blockNumber"`
	BlockHash   string   `json:"blockHash" bson:"blockHash"`
	TxHash      string   `json:"transactionHash" bson:"transactionHash"`
	Address     string   `json:"address" bson:"address"`
	Topics      []string `json:"topics" bson:"topics"`
	Data        []byte   `json:"data"  bson:"data"`
	TxIndex     uint     `json:"transactionIndex" bson:"transactionIndex"`
	Index       uint     `json:"logIndex" bson:"logIndex"`
	Removed     bool     `json:"removed" bson:"removed"`
}
