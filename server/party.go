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
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/Jeffail/leaps/lib/acl"
	"github.com/Jeffail/leaps/lib/api"
	apiio "github.com/Jeffail/leaps/lib/api/io"
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
  // scan all local files and prepares them

}
