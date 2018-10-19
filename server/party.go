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
	"os/exec"
	"os/signal"
	gopath "path"
	pf"path/filepath"
	"strings"
	"syscall"
	"time"
	"log"

	"github.com/Jeffail/leaps/lib/acl"
	"github.com/Jeffail/leaps/lib/api"
	"github.com/Jeffail/leaps/lib/audit"
	"github.com/Jeffail/leaps/lib/curator"
	"github.com/Jeffail/leaps/lib/store"
	"github.com/Jeffail/leaps/lib/util"
	"github.com/Jeffail/leaps/lib/util/service/log"
	"github.com/Jeffail/leaps/lib/util/service/metrics"
	"github.com/gorilla/websocket"
)

func init(){
  // init metadata store
	meatdataStore := MetadataStore.NewMetadataStore()
  // scan all local files and prepares them
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
			//
			return nil
		})
	}



}
