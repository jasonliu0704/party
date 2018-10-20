package utils

import "io/ioutil"
import "log"
import "crypto/sha256"
import "MetaDataStore"
import "BlockStore"
import "proto_buf"

// set of configuration parameter
BLOCK_SIZE =

/*
keep upload the file to MetatdataStore until version is up to date
*/
tryUpload(ms MetaDataStore.MetadataStore, filePath string, update_version *int32, block_hashs)
(proto_buf.WriteResult){
  // keeping modifying until version is up to date
  for true {
    file_obj := proto_buf.FileInfo{filePath, *update_version, block_hashs}

    // try modify file
    wr := ms.modifyFile(file_obj)

    log.Println("uploadFile: get version from MetaDataStore is " + wr.result)
    if(wr.result != proto_buf.OLD_VERSION)break;

    *(update_version)+=1
  }
  return wr
}

/*
support multi-threading
simulate the client process to upload file to the server
step:
1. divide file into blocks and get hash of each block
3. upload file to MetadataStore to get missing blocks
2. upload missing blocks to the blockStore
filePath is the filename
*/
UploadFile(ms MetaDataStore.MetadataStore, bs filePath string)(bool){
  content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
  version := 1

  // divide into blocks
  var localBlockMap map[string]string
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

  update_version := version
  // get missing blocks from MetatdataStore
  wr := tryUpload(ms, filePath, update_version, block_hashs)

  log.Printf("Upload missing blocks: %v\n", wr.missing_blocks)
  // upload missing blocks to blockstore
  if(wr.result == proto_buf.MISSING_BLOCKS){
    for ms := wr.missing_blocks {
      log.Println("miss block: " + ms)
      bs.StoreBlock(proto_buf.Block{ms, localBlockMap[ms]})
    }

    // upload to MeataDataStore after blocks are uploaded
    wr = tryUpload(ms, filePath, update_version, block_hashs)
  }

  // succeed
  if(wr.Result() == proto_buf.OK){
    log.Println("file uploaded")
    return true
  }

  return false
}
