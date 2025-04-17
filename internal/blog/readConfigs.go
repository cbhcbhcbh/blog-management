package blog

import (
	"blog/internal/blog/store"
	"blog/internal/pkg/log"
	"blog/pkg/db"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	recommendedHomeDir = ".blog"
	defaultConfigName  = "blog.yaml"
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		projectRoot := findProjectRoot()

		viper.AddConfigPath(filepath.Join(projectRoot, "configs"))
		viper.AddConfigPath(filepath.Join(projectRoot, recommendedHomeDir))
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}

	viper.AutomaticEnv()

	viper.SetEnvPrefix("BLOG")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}

	log.Debugw("Using config file", "file", viper.ConfigFileUsed())
}

func findProjectRoot() string {
	if root := os.Getenv("BLOG_ROOT"); root != "" {
		return root
	}

	dir, err := os.Getwd()
	if err != nil {
		return "/workspaces/miniblog"
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "/workspaces/miniblog"
}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

func initStore() error {
	dbOPtions := &db.PostgresOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}

	ins, err := db.NewPostgres(dbOPtions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)

	return nil
}
