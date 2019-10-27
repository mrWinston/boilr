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
	log.Printf("args is: %+v\n", args)

	plate, err := config.LoadPlateFile(args.ConfigPath)
	log.Printf("%v\n", plate)

	valid := plate.ValidatePlateFile()
	if !valid {
		log.Fatalf("Platefile invalid")
	}

	if err != nil {
		log.Fatal(err)
	}

	templateFolder, err := filepath.Abs(args.ConfigPath)
	templateFolder = filepath.Dir(templateFolder)
	templateFolder = fmt.Sprintf("%s/%s", templateFolder, plate.Config.TemplateRoot)
	context := pongo2.Context{}
	context, err = plate.GetVarsFromUser()
	if err != nil {
		log.Fatal(err)
	}

	//if args.VarsFile != "" {
	//	varsFile, _ := filepath.Abs(args.VarsFile)
	//	context, err = config.GetVarsFromYaml(varsFile, vars)
	//	if err != nil {
	//		log.Fatalf("Error while reading vars from '%s' : %v\n", varsFile, err)
	//	}
	//} else {
	//	context = plate.GetVarsFromUser()
	//}

	err = templating.Render(context, templateFolder, args.OutputPath)
	if err != nil {
		log.Fatal(err)
	}

}
