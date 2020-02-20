package twitter

import (
	"../http_util"
	"../http_util/oauth1"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const envFile = "config/twitter.secret.json"


var credentials = oauth1.Credentials{}

func init(){
	if err := loadEnvFile(envFile); err != nil{
		panic(err)
	}
}

func loadEnvFile(envFile string) error {
	bytes, err := ioutil.ReadFile(envFile)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &credentials)
	return err
}

func Test(){

}

func Tweet(text string) {
	fmt.Println("Tweet")
	req := oauth1.Request{
		Request: http_util.Request{
			Method: http_util.POST,
			Url:    "https://api.twitter.com/1.1/statuses/update.json",
			Params: http_util.Params{
				"status": text,
			},
			Headers: nil,
		},
	}
	req.SetCredentials(credentials)
	fmt.Println("credentials: ", credentials)

	result := req.Do()
	fmt.Println(result)
}
