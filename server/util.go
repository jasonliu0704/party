package util

import "io/ioutil"
import "log"
import "crypto/sha256"

// set of configuration parameter
BLOCK_SIZE =

/*
support multi-threading
simulate the client process to upload file to the server
step:
1. divide file into blocks and get hash of each block
3. upload file to MetadataStore to get missing blocks
2. upload missing blocks to the blockStore
filePath is the filename
*/
uploadFile(filePath string){
  content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
  version := 1

  // divide into blocks
  file_blocks := make([]string, len(content)/BLOCK_SIZE + 1)
  byte_idx = 0
  block_idx = 0
  for (byte_idx < len(content)){
    if(byte_idx + BLOCK_SIZE >= len(content)){
      file_blocks[block_idx] = string(content[byte_idx:])
    }else{
      file_blocks[block_idx] = string(content[byte_idx: byte_idx+BLOCK_SIZE])
    }
    byte_idx += BLOCK_SIZE
    block_idx++
  }

  // hash blocks
  block_hashs := make([]string, len(file_blocks))
  h := sha256.New()

  for idx, block := file_blocks {
    block_hashs[idx] = h.Sum(block)
  }

  // get missing blocks from MetatdataStore
  // keeping modifying until version is up to date
  for true {
    file_obj := proto_buf.FileInfo{filePath, }
  }

}
