package main

// go:generate packr
// this application uses github.com/gobuffalo/packr
// to statically include frontend interface code in
// a binary distributable.

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	goji "goji.io"
	"goji.io/pat"
)

const ListenAddr = "0.0.0.0:8080"

func main() {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	mux := goji.NewMux()

	// frontendBox is used to statically embed the frontend
	frontendBox := packr.NewBox("index.html")

	mux.HandleFunc(pat.New("/ui"), func(res http.ResponseWriter, req *http.Request) {
		clientScriptFile, err := frontendBox.Open("")
		if err != nil {
			http.Error(res, err.Error(), 500)
			logger.Error().
				Err(err).
				Str("ep", req.URL.String()).
				Msg("failed to serve frontend")

			return
		}

		http.ServeContent(
			res,
			req,
			"index.html",
			time.Now(),
			clientScriptFile,
		)
	})

	logger.Info().Str("addr", ListenAddr).Msg("starting application server")
	logger.Info().Msg(fmt.Sprintf("go to http://%s/ui to use the web interface", ListenAddr))

	http.ListenAndServe(ListenAddr, mux)
}
