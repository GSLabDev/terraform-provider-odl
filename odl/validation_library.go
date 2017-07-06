package odl

import (
	"fmt"
	"log"
)

func validateOperation(v interface{}, k string) (warnings []string, errors []error) {
	operation := v.(string)
	log.Printf("[INFO] Validating operation %s", operation)
	if operation == "SET" || operation == "ADD" {
		return nil, nil
	}
	return returnError("Only ADD or SET options are allowed", fmt.Errorf("[ERROR] Invalid Operation"))
}

func returnError(message string, err error) (warnings []string, errors []error) {
	var errorVar []error
	var warningVar []string
	return append(warningVar, message), append(errorVar, err)
}
