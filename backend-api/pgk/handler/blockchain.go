package handler

import (
	"github.com/unchainio/interfaces/logger"
	"github.com/unchainio/pkg/iferr"
	"io/ioutil"
	"net/http"
)

type BlockchainHandler struct {
	logger logger.Logger
}

func NewBlockchainHandler (logger logger.Logger) *BlockchainHandler {
	return &BlockchainHandler{logger: logger}
}

func (h *BlockchainHandler) Campains(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:3030/campaigns/")
	if iferr.Respond(w, err) {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	h.logger.Printf(string(body))

	w.Write(body)
}

