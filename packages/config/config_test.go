package config

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	testingpkg "github.com/example-golang-projects/clean-architecture/packages/testing"

	"github.com/stretchr/testify/assert"
)

type TestType struct {
	Test1 Test1 `json:"test_1"`
}

type Test1 struct {
	Test11 []int  `json:"test_1_1"`
	Test12 Test12 `json:"test_1_2"`
}

type Test12 struct {
	Test121 string `json:"test_1_2_1"`
}

func Test_MustLoad(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	expectedResp := &TestType{
		Test1: Test1{
			Test11: []int{0, 1, 2, 3},
			Test12: Test12{
				Test121: "value_test_1_2_1",
			},
		},
	}
	testCases := []testingpkg.TestCase{
		{
			Name: "Failed case: Error when file path is empty",
			Ctx:  ctx,
			Req: []interface{}{
				FileType_JSON,
				"",
				&TestType{},
			},
			ExpectedResp: nil,
			ExpectedErr:  errors.New("error when config path is empty"),
			Setup: func(ctx context.Context) {
			},
		},
		{
			Name: "Failed case: Error when read config file",
			Ctx:  ctx,
			Req: []interface{}{
				FileType_JSON,
				"/testfile/config.test.json",
				&TestType{},
			},
			ExpectedResp: nil,
			ExpectedErr:  errors.New("error when read config file"),
			Setup: func(ctx context.Context) {
			},
		},
		{
			Name: "Failed case: Error when type of file is invalid",
			Ctx:  ctx,
			Req: []interface{}{
				FileType_NONE,
				fmt.Sprintf("./testfile/config.test.json"),
				&TestType{},
			},
			ExpectedResp: nil,
			ExpectedErr:  errors.New("error when file type is invalid"),
			Setup: func(ctx context.Context) {
			},
		},
		{
			Name: "Failed case: Error when decode bytes to struct",
			Ctx:  ctx,
			Req: []interface{}{
				FileType_JSON,
				fmt.Sprintf("./testfile/config.test_invalid.json"),
				&TestType{},
			},
			ExpectedResp: nil,
			ExpectedErr:  errors.New("error when decode file"),
			Setup: func(ctx context.Context) {
			},
		},
		{
			Name: "Happy case (fileType = FileType_JSON)",
			Ctx:  ctx,
			Req: []interface{}{
				FileType_JSON,
				fmt.Sprintf("./testfile/config.test.json"),
				&TestType{},
			},
			ExpectedResp: nil,
			Setup: func(ctx context.Context) {
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testCase.Setup(testCase.Ctx)

			fileType := testCase.Req.([]interface{})[0].(FileType)
			pathString := testCase.Req.([]interface{})[1].(string)
			resp := testCase.Req.([]interface{})[2].(*TestType)
			err := MustLoad(fileType, pathString, resp)
			if testCase.ExpectedErr != nil {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), testCase.ExpectedErr.Error())
			} else {
				assert.Equal(t, testCase.ExpectedErr, err)
				assert.Equal(t, expectedResp, resp)
			}
		})
	}
}
