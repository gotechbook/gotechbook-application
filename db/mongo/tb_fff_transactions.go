package mongo

import "github.com/FinanceFutureFactory/3fcoin/core/core/types"

// TbFFFTransactionsName Transactions 3fTransactions信息表
const TbFFFTransactionsName = "tb_fff_transactions"

type TbFFFTransactions struct {
	TransactionHash string         `json:"transactionHash" bson:"transactionHash"`
	BlockHash       string         `json:"blockHash"  bson:"blockHash"`
	BlockNumber     uint64         `json:"blockNumber" bson:"blockNumber"`
	From            string         `json:"from" bson:"from"`
	To              string         `json:"to" bson:"to"`
	Value           string         `json:"value" bson:"value"`
	Gas             uint64         `json:"gas" bson:"gas"`
	GasPrice        uint64         `json:"gasPrice" bson:"gasPrice"`
	Nonce           uint64         `json:"nonce" bson:"nonce"`
	Data            string         `json:"data" bson:"data"`
	Size            string         `json:"size,omitempty" bson:"size"`
	Status          uint64         `json:"status" bson:"status"`
	Logs            []*types.Log   `json:"logs" bson:"logs"`
	Detail          *types.Receipt `json:"detail" bson:"detail"`
}

//for _, tx := range block.Transactions() {
//	fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083c
//	fmt.Println(tx.Value().String())    // 10000000000000000
//	fmt.Println(tx.Gas())               // 105000
//	fmt.Println(tx.GasPrice().Uint64()) // 102000000000
//	fmt.Println(tx.Nonce())             // 110644
//	fmt.Println(tx.Data())              // []
//	fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
//	chainID, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.NetworkID(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err != nil {
//		fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
//	}
//}
