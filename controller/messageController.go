package controller

import (
	"chatroom/model"
	"chatroom/util"
	"encoding/json"
	"net/http"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	if isInvalidSession(w, r) {
		return
	}
	var (
		message model.Message
		err error
	)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

   if err = message.CreateMessage(util.DB()); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, "message sent")
}
