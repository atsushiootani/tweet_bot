package simple

import (
	"log"
	"os"
)

var logger *log.Logger

func init(){
	logger = log.New(os.Stdout, "[simple]", log.Ldate | log.Ltime | log.LUTC | log.Llongfile)
}
