package blog

import (
	"blog/internal/pkg/log"
	"blog/internal/pkg/version/verflag"
	"errors"
	"fmt"
	"net/http"

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
	gin.SetMode(viper.GetString("runmode"))

	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Page Not Found",
			"path":    c.Request.URL.Path,
		})
	})

	if err := router.Run(viper.GetString("addr")); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalw("Failed to start server", "error", err)
		return err
	}

	return nil
}
