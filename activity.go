package jsonvalidate

import (
	"fmt"
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
		if isValid(input) {
			documentLoader := gojsonschema.NewStringLoader(input)

			//Check is the path is a path to file/http endpoint. If not it's a string and check the Schema against it
			if isPath(path) {
				logger.Infof("Reference Loader")
				schemaLoder := gojsonschema.NewReferenceLoader(path)

				valid, _ := check(schemaLoder, documentLoader)
				if valid {
					ctx.SetOutput("isValid", true)
					return true, nil
				}
				ctx.SetOutput("isValid", false)
				return true, nil
			}
			logger.Infof("String Loader")
			schemaLoder := gojsonschema.NewStringLoader(path)

			valid, _ := check(schemaLoder, documentLoader)
			if valid {
				ctx.SetOutput("isValid", true)
				return true, nil
			}
			ctx.SetOutput("isValid", false)
			return true, nil
		}

	}

	logger.Debugf("string is", input)
	if isValid(input) {
		ctx.SetOutput("isValid", true)
		return true, nil
	}
	ctx.SetOutput("isValid", false)

	return true, nil

}

func isValid(s string) bool {
	var js map[string]interface{}

	return json.Unmarshal([]byte(s), &js) == nil
}

func isPath(s string) bool {

	return strings.Contains(s, "/")
}

func check(schemaLoader, documentLoader gojsonschema.JSONLoader) (bool, error) {
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	fmt.Println(result, err)
	if err != nil {
		//logger.Error(err)
		return false, nil
	}

	if result.Valid() {
		logger.Infof("The document is valid\n")
		return false, nil
	}
	fmt.Println("The document is not valid. see errors :", result)
	//logger.Error("The document is not valid. see errors :\n")
	return false, nil

}
