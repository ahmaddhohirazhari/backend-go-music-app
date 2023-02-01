package server

import (
	"log"
	"net/http"
	"os"

	"backend/src/routers"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "server",
	Short: "start api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {

	

	if mainRoute, err := routers.New(); err == nil {
		var addrs string = ""

		if pr := os.Getenv("APP_PORT"); pr != "" {
			addrs = "127.0.0.1:" + pr
		}

		log.Println("App running on " + addrs)

		handler := cors.New(cors.Options{
			AllowedOrigins: []string{os.Getenv("ORIGIN_ALLOWED")},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"X-Requested-With"},
		}).Handler(mainRoute)

		if err := http.ListenAndServe(addrs, handler); err != nil {
			return err
		}


		return nil

	} else {
		return err
	}
}