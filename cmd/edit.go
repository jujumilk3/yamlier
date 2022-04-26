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

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit yaml's value",
	Long: `example:
yamlier edit [file_path] [key] [value] | [output_filepath]

If [output_filepath] doesn't set, This creates 'yamlier-outout.yaml'
How to access hierarchical value: key.innerKey.innerinnerKey
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < config.EditMinRequiredArgumentsLength {
			log.Fatalf("yamlier edit [file_path] [key] [value] | [options]")
		}

		filename, _ := filepath.Abs(args[0])
		findingKey := args[1]
		inputValue := args[2]
		outputFilename := "./yamlier-output.yaml"
		if len(args) > config.EditMinRequiredArgumentsLength {
			outputFilename = args[3]
		}

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
		if innerMap[matchedKey] == nil {
			log.Fatalf("error: There's no value of [%s]", findingKey)
		}
		innerMap[matchedKey] = inputValue

		yamlOutput, err := yaml.Marshal(yamlMap)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = ioutil.WriteFile(outputFilename, yamlOutput, 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Println("done!")

	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
