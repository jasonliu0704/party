package proto_buf

type FileInfo struct{
  filename string
  version int32
  blocklist []string
}

(FileInfo f) addAllBlocklist (blocks []string){
  for block := blocks{
    append(f.blocklist, block)
  }
}

type Result int

const(
  OK Result = 0
  OLD_VERSION Result = 1
  MISSING_BLOCKS Result = 2
  NOT_LEADER Result = 3
)

type Block struct {
  hash string
  data string
}

(Block b) Hash(){
  return b.hash
}

(Block b) Data{
  return data
}

type WriteResult {
  result Result
  current_version int32
  missing_blocks []string
}

(WriteResult wr) Result(){
  return wr.result
}

(WriteResult wr) CurrentVersion(){
  return wr.current_version
}

(WriteResult wr) MissingBlocks(){
  return wr.missing_blocks
}

type SimpleAnswer struct{
  answer bool
}
type NodeList struct{
  nodeList []int32
}

type Empty struct{}
