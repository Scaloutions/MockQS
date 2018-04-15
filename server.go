package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func usage() {
	fmt.Println("usage: example -logtostderr=true -stderrthreshold=[INFO|WARN|FATAL|ERROR] -log_dir=[string]\n")
	flag.PrintDefaults()
}

func echoString(c *gin.Context) {
	c.String(http.StatusOK, "Quote server is UP!")
}

type Quote struct {
	Price float64
	Stock string
	// UserId    string
	Timestamp int64
	CryptoKey string
}

type Request struct {
	Stock string
}

func getParams(c *gin.Context) Request {
	request := Request{}
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		glog.Error("Error processing request: %s", err)
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		glog.Error("Error parsing JSON: %s", err)
	}

	return request
}

func handleQuoteReq(c *gin.Context) {
	req := getParams(c)
	quote := getQuote(req.Stock)
	//randomly wait

	c.BindJSON(&quote)
	c.IndentedJSON(http.StatusOK, quote)
}

func main() {
	router := gin.Default()

	//glog initialization flags
	flag.Usage = usage
	flag.Parse()

	api := router.Group("/api")
	{
		api.GET("/test", echoString)
		// api.GET("/get_quote", getQuoteReq)
		api.POST("/quote", echoString)
	}

	log.Fatal(router.Run(":4444"))

}

func getQuote(stock string) Quote {
	r := rand.New(rand.NewSource(getCurrentTs()))
	price := r.Float64()
	price = float64(int(price*100)) / 100
	return Quote{
		Price: price,
		Stock: stock,
		// UserId:    userid,
		Timestamp: getCurrentTs(),
		CryptoKey: "PXdxruf7H5p9Br19Si5hq",
	}

}

func getCurrentTs() int64 {
	return time.Now().UnixNano() / 1000000
}
