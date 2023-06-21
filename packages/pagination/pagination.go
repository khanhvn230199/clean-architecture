package pagination

import (
	"errors"
	"fmt"
	arraycommom "github.com/example-golang-projects/clean-architecture/packages/common/cm_array"
	"reflect"
	"strings"
)

type Paging struct {
	Total  int      `json:"total,omitempty"`
	Limit  int      `json:"limit"`
	Page   int      `json:"page"`
	Offset int      `json:"offset"`
	Sorts  []string `json:"sorts"`
}

func (p *Paging) Validate(model interface{}) error {
	if model == nil {
		return nil
	}

	if len(p.Sorts) > 0 {
		r := reflect.ValueOf(model).Elem()
		typeOfT := r.Type()
		var normalizedColumns []string
		for i := 0; i < r.NumField(); i++ {
			fieldName := typeOfT.Field(i).Name
			normalizedColumns = append(normalizedColumns, strings.ToLower(fieldName))
		}
		// Example Sort: "Created_at desc"
		for _, sort := range p.Sorts {
			lowerSort := strings.ToLower(sort)             // created_at desc
			lowerSortStrs := strings.Split(lowerSort, " ") // ["created_at", "desc"]
			if len(lowerSortStrs) != 2 {
				return errors.New(fmt.Sprintf("Sort does not valid"))
			}
			sortField := lowerSortStrs[0]                             // "created_at"
			normalizedField := strings.ReplaceAll(sortField, "_", "") // createdat
			isContained := arraycommom.ListStringsContain(normalizedColumns, normalizedField)
			if !isContained {
				return errors.New(fmt.Sprintf("Sorted field %v does not exist in table", normalizedField))
			}
		}
	}
	return nil
}
