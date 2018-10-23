package proto_buf

type FileInfo struct{
  Filename string     `json:"file_name"`
  Version int32       `json:"version"`
  Blocklist []string  `json:"block_list"`
}

(FileInfo f) addAllBlocklist (blocks []string){
  for block := blocks{
    append(f.Blocklist, block)
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
  Hash string   `json:"hash"`
  Data string   `json:"data"`
}

type WriteResult {
  Result Result           `json:"result"`
  Current_version int32   `json:"current_version"`
  Missing_blocks []string `json:"missing_blocks"`
}

type SimpleAnswer struct{
  Answer bool   `json:"answer"`
}
type NodeList struct{
  NodeList []int32    `json:"node_list"`
}

type Empty struct{}
