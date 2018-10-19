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
(BlockStore bs) storeBlock(hash string, block string){
  bs.BlockStore[hash] = block
  return proto_buf.Empty{}
}

/*
return a block by hash
*/
(BlockStore bs) getBlock(hash string) (proto_buf.Block){
  return Block{hash, bs.blockMap[hash]}
}

/*
check whether a block exits
*/
(BlockStore bs) hashBlock(hash string) (proto_buf.SimpleAnswer){
  if(_, ok := bs.blockMap[hash]; ok){
    return proto_buf.SimpleAnswer{true}
  }else{
    return proto_buf.SimpleAnswer{false}
  }
}
