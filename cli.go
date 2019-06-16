package main

import (
	"flag"
	"log"
)

type Config struct {
	ConfigPath string
	OutputPath string
	VarsFile   string
	Command    string
	Workdir    string
	Exec       bool
}

func GetArgsConfig() *Config {
	var configPath string
	var outputPath string
	var varsFile string
	var command string
	var workdir string
	var exec bool

	flag.StringVar(&configPath, "in", "./", "Path to .plate file")
	flag.StringVar(&outputPath, "out", "", "Output folder")
	flag.StringVar(&varsFile, "vars", "", "Prefilled Vars File")
	flag.StringVar(&command, "cmd", "", "Command to Execute after Rendering. Overwrites CMD from .plate file")
	flag.StringVar(&workdir, "wd", "./", "Workdir for cmd after templating (relative to <outputPath>). Overwrite WORKDIR from .plate file")
	flag.BoolVar(&exec, "exec", false, "Run Command after templating")
	flag.Parse()

	return &Config{
		ConfigPath: configPath,
		OutputPath: outputPath,
		VarsFile:   varsFile,
		Command:    command,
		Workdir:    workdir,
		Exec:       exec,
	}

}

func ValidateConfig(config *Config) {
	if config.OutputPath == "" {
		flag.Usage()
		log.Fatal("-out Must not be empty!")
	}

	if config.ConfigPath == "" {
		flag.Usage()
		log.Fatal("-in Must not be empty!")
	}
}
