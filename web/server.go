package web

import (
	"app/config"
	"app/web/middlerware"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
)

func RunServer(wg *sync.WaitGroup) {
	mux := http.NewServeMux()
	manager := middlerware.NewManager()

	InitRoutes(mux, manager)
	wg.Add(1)

	go func() {
		defer wg.Done()

		conf := config.GetConfig()
		addr := fmt.Sprintf(":%d", conf.HttpPort)

		slog.Info(fmt.Sprintf("Listening at %s", addr))
		if err := http.ListenAndServe(addr, mux); err != nil {
			slog.Error(err.Error())
		}
	}()
}
