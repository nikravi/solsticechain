package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type blockCount struct {
	Count int
}

func apiGetBlockCount(w http.ResponseWriter, r *http.Request, bc *Blockchain) {
	respondJSON(w, http.StatusOK, blockCount{bc.GetBestHeight()})
}

func apiGetBlock(w http.ResponseWriter, r *http.Request, bc *Blockchain) {
	id := mux.Vars(r)["id"]
	var block Block
	var errBlock error
	intID, err := strconv.Atoi(id)
	if err == nil {
		block, errBlock = bc.GetBlockByHeight(intID)

	} else {
		blockHash, errDecode := hex.DecodeString(id)
		if errDecode != nil {
			errBlock = errDecode
		} else {
			block, errBlock = bc.GetBlock(blockHash)
		}

	}

	if errBlock == nil {
		respondJSON(w, http.StatusOK, block)
	} else {
		respondError(w, http.StatusBadRequest, fmt.Sprintf("Block %s was not found", id))
	}
}
