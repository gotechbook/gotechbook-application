package fff

import (
	"context"
	"github.com/FinanceFutureFactory/3fcoin/core/common/hexutil"
	"github.com/FinanceFutureFactory/3fcoin/core/core/types"
	"github.com/gotechbook/gotechbook-application/config"
	"github.com/gotechbook/gotechbook-application/db/mongo"
	"github.com/gotechbook/gotechbook-application/db/redis"
	"github.com/gotechbook/gotechbook-application/db/redis/store"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

type Blocks struct {
	component.Base
	timer *timer.Timer
}

func NewBlocks() *Blocks {
	return &Blocks{}
}

func (b *Blocks) AfterInit() {
	ctx := context.Background()
	blockChainStore := store.NewBlockChainDataStore()
	blockChainStore.UseWithCtx(ctx)
	go b.insertStore(ctx, blockChainStore)
}

func (b *Blocks) insertStore(ctx context.Context, blockChainStore *store.FFFBlockChainDataStore) {
	for {
		data, _ := blockChainStore.Get()
		if data.CurrentHeight == 0 {
			leastData := &mongo.TbFFFBlocks{}
			mongo.FindOne(ctx, config.GOTECHBOOK_MONGO, mongo.DB_FFF, mongo.TbFFFBlocksName, bson.M{}, leastData, options.FindOne().SetSort(bson.M{"number": -1}))
			chainId, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.ChainID(ctx)
			if err != nil {
				logger.Log.Error("fff request ChainID err: ", err)
				return
			}
			blockNumber, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.BlockNumber(ctx)
			if err != nil {
				logger.Log.Error("fff request BlockNumber err: ", err)
				return
			}
			err = blockChainStore.Set(redis.FFFBlockChain{
				ChainId:           chainId.String(),
				Name:              config.GOTECHBOOK_FFF.Chain.Name,
				CurrentHeight:     blockNumber,
				SynchronizeHeight: leastData.Number,
			})
			if err != nil {
				logger.Log.Error("blockChainStore set err: ", err)
				return
			}
		}
		blockNumber, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.BlockNumber(ctx)
		if err != nil {
			logger.Log.Error("fff request BlockNumber err: ", err)
			return
		}
		if data.SynchronizeHeight < blockNumber {
			data.SynchronizeHeight = data.SynchronizeHeight + 1
			block, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.BlockByNumber(ctx, new(big.Int).SetUint64(data.SynchronizeHeight))
			if err != nil {
				logger.Log.Error("fff request BlockByNumber err: ", err)
				return
			}
			for _, tx := range block.Transactions() {
				chainID, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.NetworkID(ctx)
				if err != nil {
					logger.Log.Error("fff request NetworkID err: ", err)
				}
				msg, err := tx.AsMessage(types.NewEIP155Signer(chainID))
				if err != nil {
					logger.Log.Error("fff AsMessage err: ", err)
				}
				receipt, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.TransactionReceipt(ctx, tx.Hash())
				if err != nil {
					logger.Log.Error("fff TransactionReceipt err: ", err)
				}
				isExit := &mongo.TbFFFTransactions{}
				err = mongo.FindOne(ctx, config.GOTECHBOOK_MONGO, mongo.DB_FFF, mongo.TbFFFTransactionsName, bson.M{"transactionHash": tx.Hash().Hex()}, isExit)
				if err != nil {
					_, err = mongo.Insert(ctx, config.GOTECHBOOK_MONGO, mongo.DB_FFF, mongo.TbFFFTransactionsName, &mongo.TbFFFTransactions{
						TransactionHash: tx.Hash().Hex(),
						BlockHash:       block.Hash().Hex(),
						BlockNumber:     block.NumberU64(),
						From:            msg.From().Hex(),
						To:              tx.To().Hex(),
						Value:           tx.Value().String(),
						Gas:             tx.Gas(),
						GasPrice:        tx.GasPrice().Uint64(),
						Nonce:           tx.Nonce(),
						Data:            hexutil.Encode(tx.Data()),
						Size:            tx.Size().String(),
						Status:          receipt.Status,
						Logs:            receipt.Logs,
						Detail:          receipt,
					})
					if err != nil {
						logger.Log.Error("tb_fff_transactions insert err: ", err)
						return
					}
				}
			}
			isExit := &mongo.TbFFFBlocks{}
			err = mongo.FindOne(ctx, config.GOTECHBOOK_MONGO, mongo.DB_FFF, mongo.TbFFFBlocksName, bson.M{"number": data.SynchronizeHeight}, isExit)
			if err != nil {
				_, err = mongo.Insert(ctx, config.GOTECHBOOK_MONGO, mongo.DB_FFF, mongo.TbFFFBlocksName, &mongo.TbFFFBlocks{
					Number:           block.NumberU64(),
					Difficulty:       block.Difficulty().String(),
					ExtraData:        hexutil.Encode(block.Extra()),
					GasLimit:         block.GasLimit(),
					GasUsed:          block.GasUsed(),
					Hash:             block.Hash().Hex(),
					LogsBloom:        hexutil.Encode(block.Bloom().Bytes()),
					Miner:            block.Coinbase().Hex(),
					MixHash:          block.MixDigest().Hex(),
					Nonce:            block.Nonce(),
					ParentHash:       block.ParentHash().Hex(),
					ReceiptsRoot:     block.ReceiptHash().Hex(),
					Sha3Uncles:       block.UncleHash().Hex(),
					Size:             block.Size().String(),
					StateRoot:        block.Root().Hex(),
					Timestamp:        block.Time(),
					Transactions:     block.Transactions().Len(),
					TransactionsRoot: block.TxHash().Hex(),
					Uncles:           block.Uncles(),
				})
				if err != nil {
					logger.Log.Error("tb_fff_blocks insert err: ", err)
					return
				}
			}
			err = blockChainStore.Set(redis.FFFBlockChain{
				ChainId:           data.ChainId,
				Name:              config.GOTECHBOOK_FFF.Chain.Name,
				CurrentHeight:     blockNumber,
				SynchronizeHeight: data.SynchronizeHeight,
			})
		}
	}
}
