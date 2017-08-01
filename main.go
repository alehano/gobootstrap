package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/alehano/gobootstrap/sys/cmd"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/alehano/gobootstrap/config"
	"github.com/spf13/cobra"
	"github.com/alehano/gobootstrap/sys/urls"
	_ "github.com/alehano/gobootstrap/models"
	// Add all views to enable them
	_ "github.com/alehano/gobootstrap/views/common"
	_ "github.com/alehano/gobootstrap/views/admin"
	_ "github.com/alehano/gobootstrap/views/home"
)

func main() {
	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "run_server",
		Short: "Start Application Web Server",
		Long: fmt.Sprintf("Start Application Web Server. You have to set config file "+
			"either by set %q enviroment variable with full path or place %q into a project working directory.",
			config.ENVName, config.Filename),
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Server running on :%d\n", config.Get().Port)
			runServer()
		},
	})

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func runServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	urls.AddAll(r)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Port), r)
}
