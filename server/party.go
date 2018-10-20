/*
This is the server side of party editor
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"path/filepath"
	"net/http"
	"os"
	gopath "path"
	pf"path/filepath"
	"strings"
	"log"
	"utils"
)

func init(){
  // init metadata store
	meatdataStore := MetadataStore.NewMetadataStore()

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

func main(){
	// start serving MeataDataStore

}
