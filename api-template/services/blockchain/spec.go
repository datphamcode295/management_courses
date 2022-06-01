package blockchain

type BlockChainService interface {
	SyncCurrentHeadBlock()
	CurrentHeadBlock() int64
}
