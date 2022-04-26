package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"yamlier/cmd/internal/util"
	"yamlier/config"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get yaml's value",
	Long: `example:
yamlier get [file_path] [key]
How to access hierarchical value: key.innerKey.innerinnerKey
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < config.GetMinRequiredArgumentsLength {
			log.Fatalf("yamlier get [file_path] [key]")
		}

		filename, _ := filepath.Abs(args[0])
		findingKey := args[1]

		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		yamlMap := make(map[interface{}]interface{})
		err = yaml.Unmarshal(yamlFile, &yamlMap)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		innerMap, matchedKey := util.FindValue(findingKey, yamlMap)
		yamlOutput, err := yaml.Marshal(innerMap[matchedKey])
		if err != nil {
			log.Fatalf("error: %v", err)
		} else if innerMap[matchedKey] == nil {
			log.Fatalf("error: There's no value of [%s]", findingKey)
		}
		fmt.Println(string(yamlOutput))

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
