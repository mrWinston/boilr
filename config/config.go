// Package config provides measures of getting the configuration before running
// boilr. This includes parsing of the provided *.plate files as well as
// querying the User for configuration values
package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/go-yaml/yaml"
)

var reservedKeys = map[string]bool{
	"TEMPLATE_ROOT": true,
	"CMD":           true,
	"WORKDIR":       true,
}

// LoadTemplateConfig loads a .plate file for templating. It returns an error
// in case either the unmarshalling of the yml is unseccessful or the file can
// not be read.
// the Return value is intended to be plugged directly into QueryVarsFromUser
func LoadTemplateConfig(path string) (map[string]string, error) {
	m := make(map[string]string)
	fileContent, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileContent, &m)
	if err != nil {
		return nil, err
	}
	log.Println(m)
	return m, nil

}

// QueryVarsFromUser is used to query the configuration with a TUI interface
// from the user. It accepts a Configuration Map of the following format:
// map: string-> string
//
// Where the keys represent the name of the variable to query and the value
// represents the type of the variable to query. Supported types are "string"
// and "list.
//
// It converts the Input of the user to a pongo2.Context map for later use with
// templating.
func QueryVarsFromUser(config map[string]string) pongo2.Context {
	context := pongo2.Context{}
	reader := bufio.NewReader(os.Stdin)
	for key, value := range config {

		if reservedKeys[key] {
			continue
		}

		switch value {
		case "list":
			fmt.Printf("Provide Values for: %v \n Seperate by Linebreak, End with Empty Line\n", key)
			var list []string
			input, _ := reader.ReadString('\n')
			for input != "\n" {
				list = append(list, strings.Replace(input, "\n", "", 1))
				input, _ = reader.ReadString('\n')
			}
			context[key] = list
		case "string":
			fmt.Printf("Enter a Value for %v: ", key)
			var input string
			input, _ = reader.ReadString('\n')
			context[key] = strings.Replace(input, "\n", "", 1)
		default:
			log.Panicf("Invalid Config: %v is %v , supported is: list, string\n", key, value)
		}
	}
	return context
}

func GetVarsFromYaml(path string, config map[string]string) (pongo2.Context, error) {
	//	context := pongo2.Context{}
	m := make(map[string]interface{})

	fileContent, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileContent, &m)
	if err != nil {
		return nil, err
	}

	// check if all keys from vars file are present in config
	for k, _ := range m {
		if _, ok := config[k]; !ok {
			return nil, fmt.Errorf("Key '%v' from vars file not defined in Config", k)
		}
	}

	// check if all keys from config are present in vars file
	for k, _ := range config {
		if reservedKeys[k] { // we do not expect to find the reserved keys in the vars file
			continue
		}
		if _, ok := m[k]; !ok {
			return nil, fmt.Errorf("Required key '%v' not found in vars file", k)
		}
	}

	return m, nil

}
