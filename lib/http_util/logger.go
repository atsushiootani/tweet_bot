package http_util

import (
	"log"
	"os"
)

var logger *log.Logger

func init(){
	logger = log.New(os.Stdout, "[http_util]", log.Ldate | log.Ltime | log.LUTC | log.Llongfile)
}
