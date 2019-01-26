package config

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// func init() {
// 	viper.SetConfigName(".perf_config")
// 	viper.AddConfigPath("$HOME/.perf_collector")
// 	viper.AddConfigPath(".")
// 	err := viper.ReadInConfig() // Find and read the config file
// 	if err != nil {             // Handle errors reading the config file
// 		panic(fmt.Errorf("Fatal error config file: %s ", err))
// 	}
// 	fmt.Println("VIPER INIT", viper.GetBool("collectors.metrics.cpu"))
// }

func LoadConfig(cmd *cobra.Command) error {
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// from a config file
		viper.SetConfigName(".perf_config")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/.perf_collector")
	}
	// from the environment
	viper.SetEnvPrefix("PERF")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// NOTE: this will require that you have config file somewhere in the paths specified. It can be reading from JSON, TOML, YAML, HCL, and Java properties files.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil

}
