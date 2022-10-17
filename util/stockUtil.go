package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetStock(stockCode string) string{
	resp, err := http.Get("https://stooq.com/q/l/?s="+stockCode+"&f=sd2t2ohlcv&h&e=csv")
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	timeMessage:=time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		return timeMessage+" Stock Bot: error"+err.Error()
	}
	stockArray:=strings.Split(string(body),",")
	volumeArray:=strings.Split(stockArray[7],"Volume")
	quoteArray := strings.Split(volumeArray[1], "\n")

	if strings.Contains(string(body), "N/D"){
		return timeMessage+" Stock Bot:"+ quoteArray[1] +" quote not found"
	}else{
		return timeMessage+" Stock Bot:"+ quoteArray[1] +" quote is $"+ stockArray[11]+ " per share"
	}

}
