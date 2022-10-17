package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetStock(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	stockRequest := vars["stock"]
	stockCode:=strings.TrimLeft(stockRequest,"/stock=")

	resp, err := http.Get("https://stooq.com/q/l/?s="+stockCode+"&f=sd2t2ohlcv&h&e=csv")
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusOK, "Invalid request payload")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		respondWithError(w, http.StatusOK, "Stock Bot: error"+err.Error())
	}
	stockArray:=strings.Split(string(body),",")

	if len (stockArray) > 7 {
		if stockArray[7] != "N/D"{
			volumeArray:=strings.Split(stockArray[7],"Volume")
			quoteArray := strings.Split(volumeArray[1], "\n")

			if strings.Contains(string(body), "N/D"){
				respondWithJSON(w, http.StatusOK, "Stock Bot:"+ quoteArray[1] +" quote not found")
			}else{
				respondWithJSON(w, http.StatusOK,"Stock Bot:"+ quoteArray[1] +" quote is $"+ stockArray[11]+ " per share")
			}
		}
	}else{
		respondWithJSON(w, http.StatusOK, "Stock Bot: Timeout error. Please try again.")
	}
}
