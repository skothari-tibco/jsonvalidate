package jsonvalidate

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `
{
    "name": "jsonvalidate",
    "type": "flogo:activity",
    "ref": "github.com/skothari-tibco/jsonvalidate",
    "version": "0.0.1",
    "title": "JSON Validator",
    "description": "Simple JSON Validator Activity",
    "homepage": " ",
    "input":[
      {
        "name": "text",
        "type": "string",
        "value": ""
      },
      {
        "name": "path",
        "type":"string",
        "value":""
      }
    ],
    "output": [
      {
        "name": "isValid",
        "type": "bool"
      }
    ]
  }
  `

func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
