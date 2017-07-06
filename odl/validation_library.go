package odl

import (
	"fmt"
)

func validateOperation(v interface{}, k string) (warnings []string, errors []error) {
	operation := v.(string)
	if operation == "SET" || operation == "ADD" {
		return nil, nil
	}
	return []string{"Only ADD or SET options are allowed"}, []error{fmt.Errorf("[ERROR] Invalid Operation")}
}
