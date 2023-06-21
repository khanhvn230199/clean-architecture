package testingpkg

import (
	"context"
	"fmt"
)

type TestCase struct {
	Name         string
	Ctx          context.Context
	Req          interface{}
	ExpectedResp interface{}
	ExpectedErr  error
	Setup        func(ctx context.Context)
}

var (
	ErrDefault = fmt.Errorf("something wrongs")
)
