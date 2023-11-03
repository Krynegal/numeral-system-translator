package validators

import (
	"fmt"
	"github.com/Krynegal/numeral-system-translator.git/internal/models"
)

type errFieldIsUndefined struct {
	fieldName string
}

func (e errFieldIsUndefined) Error() string {
	return fmt.Sprintf("field %s is undefined", e.fieldName)
}

func CheckRequest(request *models.Request) error {
	if request.Number == nil {
		return errFieldIsUndefined{fieldName: "number"}
	}

	if request.Base == nil {
		return errFieldIsUndefined{fieldName: "base"}
	}
	if request.ToBase == nil {
		return errFieldIsUndefined{fieldName: "to_base"}
	}

	return nil
}

func CheckBase(base int) error {
	if base < 2 || base > 16 {
		return fmt.Errorf("base number cannot be grater than 16 or less than 2")
	}

	return nil
}
