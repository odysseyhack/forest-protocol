package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/unchainio/pkg/xlogger"
	"github.com/unchainio/pkg/iferr"
	"github.com/BurntSushi/toml"
	"github.com/go-chi/chi"
	"github.com/goware/cors"
	"gitlab.com/henkvanramshorst/forest/recommender-wrapper/pkg/domain"
	"fmt"
	"os/exec"
	"github.com/go-chi/render"
	"encoding/json"
)

func main() {
	logger, err := xlogger.New(nil)
	iferr.Exit(err)

	logger.Printf("Start Recommendation Wrapper")

	cfgFile, err := ioutil.ReadFile("./config/dev/config.toml")
	iferr.Exit(err)

	cfg := new(domain.Config)
	err = toml.Unmarshal(cfgFile, cfg)
	iferr.Exit(err)

	logger.Printf("config loaded: %+v", cfg)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})


	r := chi.NewRouter()

	r.Use(c.Handler)

	// Health Check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Recommender system"))
	})

	r.Post("/recommend", func(w http.ResponseWriter, r *http.Request) {
		var d domain.RecommendationRequest
		render.DecodeJSON(r.Body, d)

		cmd := exec.Command(fmt.Sprintf("python ./script/factorization_machine.py %b %b %b %b %b", d.Irr, d.Biodiversity, d.EnvProtection, d.JobOpportunities, d.FoodSecurity))

		out, err := cmd.Output()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("could not fetch output of script"))
			return
		}

		var response domain.RecommendationResponse

		err = json.Unmarshal(out, &response)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("could not parse output"))
		}

		// response parses so we can return out
		w.Write(out)
	})

	log.Printf("listening on: %s", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, r))

}