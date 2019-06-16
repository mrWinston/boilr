package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/mrWinston/boilr/config"
	"github.com/mrWinston/boilr/templating"
)

func main() {

	args := GetArgsConfig()

	ValidateConfig(args)

	if args.OutputPath == "" {
		log.Fatal("-out Must be set")
	}

	vars, err := config.LoadTemplateConfig(args.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	templateFolder, err := filepath.Abs(args.ConfigPath)
	templateFolder = filepath.Dir(templateFolder)
	templateFolder = fmt.Sprintf("%s/%s", templateFolder, vars["TEMPLATE_ROOT"])
	context := pongo2.Context{}
	if args.VarsFile != "" {
		varsFile, _ := filepath.Abs(args.VarsFile)
		context, err = config.GetVarsFromYaml(varsFile, vars)
		if err != nil {
			log.Fatalf("Error while reading vars from '%s' : %v\n", varsFile, err)
		}
	} else {
		context = config.QueryVarsFromUser(vars)
	}

	err = templating.Render(context, templateFolder, args.OutputPath)
	if err != nil {
		log.Fatal(err)
	}

}
