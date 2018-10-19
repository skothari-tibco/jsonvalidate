package jsonvalidate

import (
	"bytes"
	"errors"
	"strings"

	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/xeipuuv/gojsonschema"
)

var activityLog = logger.GetLogger("activity-flogo-json-validate")

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

type JsonValidate struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &JsonValidate{metadata: metadata}
}

func (a *JsonValidate) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *JsonValidate) Eval(ctx activity.Context) (done bool, err error) {

	input := ctx.GetInput("text").(string)

	path := ctx.GetInput("path").(string)

	//If path is not defined directly try to Unmarshall
	if path != "" {
		if o, _ := isValid(input); o {
			documentLoader := gojsonschema.NewStringLoader(input)

			//Check is the path is a path to file/http endpoint. If not it's a string and check the Schema against it
			if isPath(path) {
				logger.Infof("Reference Loader")
				schemaLoder := gojsonschema.NewReferenceLoader(path)

				valid, err := check(schemaLoder, documentLoader)
				if valid {

					ctx.SetOutput("isValid", true)
					return true, nil
				}
				ctx.SetOutput("log", err)
				ctx.SetOutput("isValid", false)
				return true, nil
			}
			logger.Infof("String Loader")
			schemaLoder := gojsonschema.NewStringLoader(path)

			valid, err := check(schemaLoder, documentLoader)
			if valid {
				ctx.SetOutput("isValid", true)
				return true, nil
			}
			ctx.SetOutput("log", err)
			ctx.SetOutput("isValid", false)

			return true, nil
		}

	}

	logger.Debugf("string is", input)
	o, err := isValid(input)
	if o {

		ctx.SetOutput("isValid", true)
		return true, nil
	}
	ctx.SetOutput("log", err)
	ctx.SetOutput("isValid", false)

	return true, nil

}

func isValid(s string) (bool, error) {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(s), &js)
	if err == nil {
		return true, nil
	}

	return false, errors.New("Invalid JSON")

}

func isPath(s string) bool {

	return strings.Contains(s, "/")
}

func check(schemaLoader, documentLoader gojsonschema.JSONLoader) (bool, string) {
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		//logger.Error(err)
		return false, err.Error()
	}

	if result.Valid() {
		logger.Infof("The document is valid\n")
		return true, ""
	}
	//fmt.Println("The document is not valid. see errors :", result)
	//logger.Error("The document is not valid. see errors :\n")
	var b bytes.Buffer
	for _, desc := range result.Errors() {
		b.WriteString(desc.String())
		b.WriteString("\n")
	}

	return false, b.String()

}
