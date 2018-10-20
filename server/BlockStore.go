package BlockStore

import "proto_buf"
type BlockStore struct {
  blockMap map[string][]string
}

NewBlockStore(){
  // setup config
  return BlockStore{}
}

/*
store a block
*/
(BlockStore bs) StoreBlock(block proto_buf.Block){
  bs.BlockStore[block.Hash()] = block
  return proto_buf.Empty{}
}

/*
return a block by hash
*/
(BlockStore bs) GetBlock(hash string) (proto_buf.Block){
  return Block{hash, bs.blockMap[hash]}
}

/*
check whether a block exits
*/
(BlockStore bs) HasBlock(hash string) (proto_buf.SimpleAnswer){
  if(_, ok := bs.blockMap[hash]; ok){
    return proto_buf.SimpleAnswer{true}
  }else{
    return proto_buf.SimpleAnswer{false}
  }
}
