package jsonvalidate

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

/*
func TestWithoutPath(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("text", "{\"name\":\"Abc\"}")

	act.Eval(tc)

	result := tc.GetOutput("isValid")
	assert.Equal(t, result, true)

}
func TestWithoutPath2(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("text", "{\"name\":\"Abc\", \"age\":\"XYZ\"}")

	act.Eval(tc)

	result := tc.GetOutput("isValid")
	assert.Equal(t, result, false)

}
*/
/*
func TestWithNullValue(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("text", "")

	act.Eval(tc)

	result := tc.GetOutput("isValid")
	assert.Equal(t, result, false)

}
*/
func TestWithFilePath(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("text", "{\r\n  \"checked\": false,\r\n  \"dimenions\": {\r\n    \"width\": 5,\r\n    \"height\": 10\r\n  },\r\n  \"id\": 1,\r\n  \"name\": \"A green door\",\r\n  \"price\": 12.5,\r\n  \"tags\": [\r\n    \"home\",\r\n    \"green\"\r\n  ]\r\n}")
	tc.SetInput("path", "file:///Users/skothari-tibco/flogo/json_validator.json")

	act.Eval(tc)

	result := tc.GetOutput("isValid")
	assert.Equal(t, result, false)

}

/*
func TestWithHttpPath(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("text", "")
	tc.SetInput("path", "")

	act.Eval(tc)

	result := tc.GetOutput("isValid")
	assert.Equal(t, result, false)
}
*/
