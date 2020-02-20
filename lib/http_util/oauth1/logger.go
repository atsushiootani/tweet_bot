package oauth1

import (
	"log"
	"os"
)

var logger *log.Logger

func init(){
	file, err := os.OpenFile("log/development.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0777)
	if err != nil{
		panic(err)
	}

	logger = log.New(file, "[oauth1]", log.Ldate | log.Ltime | log.LUTC | log.Llongfile)
}
