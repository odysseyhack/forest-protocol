package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/unchainio/interfaces/logger"
	"github.com/unchainio/pkg/iferr"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/backend-api"
)

const (
	campaignsPath     = "/campaigns"
	conversationsPath = "/conversations"
	dacsPath          = "/dacs"
	donationsPath     = "/donations"
	eventsPath        = "/events"
	milestonesPath    = "/milestones"
	pledgeAdminsPath  = "/pledgeAdmins"
	usersPath         = "/users"
	whitelistPath     = "/whitelist"
)

type BlockchainHandler struct {
	logger logger.Logger
	cfg    *backend_api.Config
}

func NewBlockchainHandler(logger logger.Logger, cfg *backend_api.Config) *BlockchainHandler {
	return &BlockchainHandler{logger: logger, cfg: cfg}
}

func (h *BlockchainHandler) Campaigns(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(campaignsPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) Conversations(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(conversationsPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) DACs(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(dacsPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) Donations(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(donationsPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) Events(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(eventsPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) Milestones(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(milestonesPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) PledgeAdmins(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(pledgeAdminsPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) Users(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(usersPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) Whitelist(w http.ResponseWriter, r *http.Request) {
	body, err := h.callFeatherAPI(whitelistPath)
	iferr.Respond(w, err)

	w.Write(body)
}

func (h *BlockchainHandler) callFeatherAPI(path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", h.cfg.FeatherBaseURL, path))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	h.logger.Printf(string(body))

	return body, nil
}
