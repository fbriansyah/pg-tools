package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpModule "github.com/fbriansyah/pg-tools/port/http"
	controller_auth "github.com/fbriansyah/pg-tools/port/http/controller/auth"
	controller_dashboard "github.com/fbriansyah/pg-tools/port/http/controller/dashboard"
	controller_qris "github.com/fbriansyah/pg-tools/port/http/controller/qris"
	controller_setting "github.com/fbriansyah/pg-tools/port/http/controller/setting"
	controller_user "github.com/fbriansyah/pg-tools/port/http/controller/user"
	controller_virtualAccount "github.com/fbriansyah/pg-tools/port/http/controller/virtualAccount"
)

//go:embed public
var FS embed.FS

func main() {
	port := "3033"
	router := httpModule.NewRoute(&FS, &httpModule.RouterModule{
		DashboardCtrl: controller_dashboard.New(),
		AuthCtrl:      controller_auth.New(),
		VACtrl:        controller_virtualAccount.New(),
		QrisCtrl:      controller_qris.New(),
		UserCtrl:      controller_user.New(),
		SettingCtrl:   controller_setting.New(),
	})

	// The HTTP Server
	server := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", port), Handler: router}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
