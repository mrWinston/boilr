package templating

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flosch/pongo2"
)

type renderJob struct {
	inFile  string
	outFile string
	context pongo2.Context
}

func splitJob(job renderJob) ([]renderJob, error) {
	matched, wholeMatch, createdVarName, arrayVarName := extractTemplate(job.outFile)
	if !matched {
		return []renderJob{job}, nil
	}

	if _, ok := job.context[arrayVarName]; !ok {
		return nil, fmt.Errorf("Can not find key %v for job %v", arrayVarName, job)
	}

	values := job.context[arrayVarName].([]string)

	var jobs []renderJob
	for _, value := range values {
		newContext := copyContext(job.context)
		newContext[createdVarName] = value

		jobs = append(jobs, renderJob{
			inFile:  job.inFile,
			outFile: strings.Replace(job.outFile, wholeMatch, fmt.Sprintf("%v", value), 1),
			context: newContext,
		})
	}
	return jobs, nil

}

func renderPathTemplates(jobs []renderJob) ([]renderJob, error) {
	var outjobs []renderJob
	for _, job := range jobs {
		tpl, err := pongo2.FromString(job.outFile)
		if err != nil {
			log.Printf("Got an error during Path Rendering: %v", err)
			return nil, err
		}
		renderedPath, err := tpl.Execute(job.context)
		if err != nil {
			log.Printf("Got an error during Path Rendering: %v", err)
			return nil, err
		}

		job.outFile = renderedPath
		outjobs = append(outjobs, job)
	}

	return outjobs, nil
}

func extractTemplate(input string) (bool, string, string, string) {
	re := regexp.MustCompile(forRegex)
	matches := re.FindStringSubmatch(input)
	if matches == nil || len(matches) < 3 {
		return false, "", "", ""
	}

	return true, matches[0], matches[1], matches[2]

}

func copyContext(context pongo2.Context) pongo2.Context {
	newContext := pongo2.Context{}
	for key, value := range context {
		newContext[key] = value
	}
	return newContext
}

func postProcessJobs(jobs []renderJob) ([]renderJob, []error) {
	var errors []error
	for !isProcessingDone(jobs) {
		var tmpArray []renderJob
		for _, job := range jobs {
			splittedJobs, err := splitJob(job)
			if err != nil {
				errors = append(errors, err)
			}
			tmpArray = append(tmpArray, splittedJobs...)
		}
		jobs = tmpArray
	}
	if len(errors) > 0 {
		return nil, errors
	}
	return jobs, nil
}

func renderJobs(jobs []renderJob) error {
	var templates []*pongo2.Template
	for _, job := range jobs {
		template, err := pongo2.FromFile(job.inFile)
		if err != nil {
			return err
		}
		templates = append(templates, template)
	}

	log.Printf("Made Templates: %v", len(templates))

	for i, job := range jobs {
		// create folder for file
		outFolder := filepath.Dir(job.outFile)
		if err := os.MkdirAll(outFolder, os.ModePerm); err != nil {
			log.Printf("Got an Erro while creating Dir: %v : %v", outFolder, err)
		}
		if strings.HasSuffix(job.inFile, ".j2") {

			file, err := os.Create(job.outFile[:len(job.outFile)-3])
			if err != nil {
				return err
			}
			defer closeAndLogError(file)

			if err := templates[i].ExecuteWriter(job.context, file); err != nil {
				fmt.Printf("Error Writing Template %v: %v", jobs[i].outFile, err)
			}
		} else {
			_, err := copyFile(job.inFile, job.outFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isProcessingDone(jobs []renderJob) bool {
	re := regexp.MustCompile(forRegex)

	for _, job := range jobs {
		if re.MatchString(job.outFile) {
			return false
		}
	}
	return true

}

type closable interface {
	Close() error
}

func closeAndLogError(c closable) {
	if err := c.Close(); err != nil {
		fmt.Println(err)
	}
}
