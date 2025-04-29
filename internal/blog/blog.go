package blog

import (
	"blog/internal/pkg/log"
	mw "blog/internal/pkg/middleware"
	"blog/internal/pkg/version/verflag"

	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func NewBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blog",
		Short: "blog subcommand",
		Long:  `This is a blog subcommand`,
		RunE: func(cmd *cobra.Command, args []string) error {
			verflag.PrintAndExitIfRequested()

			log.Init(logOptions())
			defer log.Sync()

			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the blog configuration file. Empty string for no configuration file.")

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run() error {
	if err := initStore(); err != nil {
		return err
	}

	gin.SetMode(viper.GetString("runmode"))
	router := gin.Default()

	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache(), mw.Cors(), mw.Secure(), mw.RequestID(), mw.Authn()}
	router.Use(mws...)

	if err := installRouters(router); err != nil {
		return err
	}

	httpsrv := startInsecureServer(router)

	httpssrv := startSecureServer(router)

	return gracefulShutdown(httpsrv, httpssrv)
}

func gracefulShutdown(srv *http.Server, ssrv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Infow("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("shutdown-timeout"))
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorw("Server forced to shutdown", "err", err)
		return err
	}

	if err := ssrv.Shutdown(ctx); err != nil {
		log.Errorw("Secure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("Server exited")
	return nil
}

func startInsecureServer(g *gin.Engine) *http.Server {
	httpsrv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}

	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	return httpsrv
}

func startSecureServer(g *gin.Engine) *http.Server {
	httpssrv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}

	log.Infow("Start to listening the incoming requests on https address", "addr", viper.GetString("tls.addr"))
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			if err := httpssrv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalw(err.Error())
			}
		}()
	}
	return httpssrv
}
