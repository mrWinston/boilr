// Package config provides measures of getting the configuration before running
// boilr. This includes parsing of the provided *.plate files as well as
// querying the User for configuration values
package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flosch/pongo2"
	"github.com/go-yaml/yaml"
	"github.com/mrWinston/boilr/input"
)

var reservedKeys = map[string]bool{
	"TEMPLATE_ROOT": true,
	"CMD":           true,
	"WORKDIR":       true,
}

// CommandConfig stores the configuration for a run cmd
type CommandConfig struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Workdir string `yaml:"workdir"`
}

// Return the String representation
func (commandConfig *CommandConfig) String() string {
	return fmt.Sprintf("{Name: '%s', Command: '%s', Workdir: '%s'}", commandConfig.Name, commandConfig.Command, commandConfig.Workdir)
}

// Config Holds the configuration for rendering the template
type Config struct {
	TemplateRoot string `yaml:"template_root"`
}

// PlateFile Represents the *.plate files
type PlateFile struct {
	Vars     map[string]string `yaml:"vars"`
	Config   Config            `yaml:"config"`
	Commands []CommandConfig   `yaml:"commands"`
}

// ValidatePlateFile Validates a given PlateFile for correctnes
func (plateFile *PlateFile) ValidatePlateFile() bool {

	return true
}

// GetVarsFromUser Queries the user for input based on the vars defined in the platefile
func (plateFile *PlateFile) GetVarsFromUser() (pongo2.Context, error) {
	context := pongo2.Context{}
	for k, v := range plateFile.Vars {
		input, err := input.GetInputByName(v)
		if err != nil {
			return nil, err
		}

		context[k], err = input.Ask(v, os.Stdin, os.Stdout)
		if err != nil {
			return nil, err
		}
	}

	return context, nil
}

func (plateFile *PlateFile) GetVarsFromYaml(path string) (pongo2.Context, error) {

	m := make(map[string]interface{})

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileContent, &m)

	if err != nil {
		return nil, err
	}

	for key, _ := range m {
		if _, ok := plateFile.Vars[key]; !ok {
			return nil, fmt.Errorf("Key %s not defined in PlateFile", key)
		}
	}
	for key, _ := range plateFile.Vars {
		if _, ok := m[key]; !ok {
			return nil, fmt.Errorf("Key %s not defined in Vars File", key)
		}
	}

	fmt.Printf("%v", m)
	return m, err

}

// LoadPlateFile Loads a plate file from the given path. Returns an error when
// reading the file fails. Does not validate the plate file
func LoadPlateFile(path string) (*PlateFile, error) {
	var config PlateFile
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileContent, &config)

	fmt.Printf("%v", config)

	return &config, err
}
