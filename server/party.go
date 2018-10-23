/*
This is the server side of party editor
*/
package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	pf"path/filepath"
	"strings"
	"log"
	"utils"
	"github.com/julienschmidt/httprouter"
	"BLockStore"
	"MeataDataStore"
	"net/http"
	"golang.org/x/net/websocket"
)

var meatdataStore MetadataStore.MetadataStore
var blockStore BlockStore.BlockStore

func init(){
  // init metadata and block store
	meatdataStore = MetadataStore.NewMetadataStore()
	blockStore = BlockStore.NewBlockStore()

  // scan all local files and stores them in MetatdataStore
	if pwd, err := os.Getwd(); err != nil{
		log.Fatal(err)
	}else{
		pf.Walk(pwd, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			fmt.Printf("visited file or dir: %q\n", path)
			// put the file into our metatdata store
			utils.UploadFile(ms, path)
			return nil
		})
	}

	// TODO: setup come configuration
}

// setup apis

// blockstore apis
// check if block exist
// read block_hash from header
// send SimpleAnswer to respond
func hasBlock(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	 answer := blockStore.HasBlock(r.Header.Get("block_hash"))
	 answerBytes, err := json.Marshal(answer)
	 if err != nil {
		 log.Printf("hasBlock: marshall error %#v\n", answerBytes)
		 w.WriteHeader(http.StatusInternalServerError)
		 return
	 }
	 w.WriteHeader(http.StatusOK)
	 w.Write(answerBytes)
}

// store a block to blockStore
// get a Block object from request body
// send Empty response
func storeBlock(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	// unmarshall Block object
	body, err := ioutil.ReadAll(req.Body)
  if err != nil {
      panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
  }
  log.Println(string(body))
  block := new(proto_buf.Block)
  err = json.Unmarshal(body, block)
  if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
			return
  }

	// actually store the block
	EmptyBytes, err := json.Marshal(blockStore.StoreBlock(*block))
	w.WriteHeader(http.StatusOK)
	w.Write(EmptyBytes)
}

// Get a block from blockStore
// get the block hash from reqeust header block_hash
// respond with proto_buf.Block object
func getBlock(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	block_hash := r.Header.Get("block_hash")
	BlockBytes, err := json.Marshal(blockStore.GetBlock(block_hash))
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
			return
  }
	w.WriteHeader(http.StatusOK)
	w.Write(BlockBytes)
}

// MetaDataStore apis

// readFile read file
// request has file_name as header
// reponse with proto_buf.FileInfo
func readFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	file_name := r.Header.Get("file_name")
	FileInfoBytes, err := json.Marshal(meatdataStore.readFile(file_name))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(FileInfoBytes)
}

// modifyFile modify file
func modifyFile()


func main(){
	// start serving MeataDataStore
	// only the ModifyFile endpoint requires web socket connection
	router := httprouter.New()
  router.GET("/hasBlock", hasBlock)
  router.POST("/storeBlock", storeBlock)
	router.Get("/getBlock", getBlock)
	router.Get("/readFile", readFile)

	// modifyFile use websocket
	http.Handle("/ws/modifyFile", websocket.Handler(modifyFile))

  log.Fatal(http.ListenAndServe(":8080", router))
}
