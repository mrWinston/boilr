package templating

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/flosch/pongo2"
)

const forRegex = `__for_(?P<var>\w+)_in_(?P<vars>\w+)__`
const substRegex = `__(?P<var>\w+)__`

// Render is responsible for actually creating the templated files and folder
// from the templateRoot and rendering the to outRoot. You will need to make
// sure beforehand, that the context provided contains all variables that are
// used in the templates.
func Render(context pongo2.Context, templateRoot string, outRoot string) error {
	absTempPath, err := convertToAbsPath(templateRoot, nil)
	absOutPath, err := convertToAbsPath(outRoot, err)
	err = errorForFolderExist(absTempPath, err)
	err = createFolderIfNotExist(absOutPath, err)

	if err != nil {
		return err
	}

	jobs, err := preprocessTemplateRoot(absTempPath, absOutPath, context)
	if err != nil {
		log.Printf("Error during Preprocessing")
		return err
	}
	jobs, postProcessErrs := postProcessJobs(jobs)
	jobs, err = renderPathTemplates(jobs)
	if err != nil {
		return err
	}
	if len(postProcessErrs) != 0 {
		var errorStringBuilder strings.Builder
		errorStringBuilder.WriteString(fmt.Sprintf("Got %d Errors while Processing Templates: \n", len(postProcessErrs)))
		for _, curErr := range postProcessErrs {
			errorStringBuilder.WriteString(fmt.Sprintf("\t %v \n", curErr))
		}
	}

	err = renderJobs(jobs)

	if err != nil {
		log.Printf("Error during Rendering")
	}
	return err
}

func createFolderIfNotExist(path string, err error) error {
	if err != nil {
		return err
	}
	return os.MkdirAll(path, os.ModePerm)
}
func errorForFolderExist(path string, err error) error {
	if err != nil {
		return err
	}
	_, err = os.Stat(path)
	return err
}
func convertToAbsPath(path string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return filepath.Abs(path)
}

func preprocessTemplateRoot(templateRoot string, outRoot string, defaultContext pongo2.Context) ([]renderJob, error) {
	var jobs []renderJob

	err := filepath.Walk(templateRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error Accessing Template Path: %v -- \t %v", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		outFile, _ := templatePathToNewPath(templateRoot, outRoot, path)
		jobs = append(jobs, renderJob{
			inFile:  path,
			outFile: outFile,
			context: defaultContext,
		})
		return nil
	})
	return jobs, err
}

func templatePathToNewPath(templateRoot string, outRoot string, path string) (string, error) {
	if !strings.HasPrefix(path, templateRoot) {
		return "", fmt.Errorf("Path not in Template Root: %s", templateRoot)
	}
	//absReturnPath, err := filepath.Abs(strings.Replace(path, templateRoot, outRoot, 1))

	return filepath.Abs(strings.Replace(path, templateRoot, outRoot, 1))

}

func convertToStringArray(in interface{}) ([]string, error) {
	var outArray []string
	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice:
		outArray = in.([]string)
	default:
		return nil, fmt.Errorf("Type is Incorrect: got %v, Want %v", reflect.TypeOf(in), reflect.Slice)
	}
	return outArray, nil

}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer closeAndLogError(source)

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer closeAndLogError(destination)
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
