package utils

import "errors"

// Checks the types of the variables
func CheckVariableType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

// Makes sure only allowed values are being passed
func CheckVariableValues(v interface{}, canTake []string) error {
	for _, values := range canTake {
		if values == v {
			return nil
		}
	}
	return errors.New("can't take the passed value")
}

// Checks the type of the variable and the value passed
func CheckTypeAndValues(v interface{}, variableType string, canTake []string) error {
	if CheckVariableType(v) != variableType {
		return errors.New("variable type must be a " + variableType)
	}
	if err := CheckVariableValues(v, canTake); err != nil {
		return err
	}
	return nil
}
