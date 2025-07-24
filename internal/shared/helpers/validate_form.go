// internal/shared/helpers/validate_form.go
package helpers

import "fmt"

func ValidateFormFields(filename, outputFilename, width string) error {
	if outputFilename == "" {
		return fmt.Errorf("field output name not filled")
	}
	if filename == "" {
		return fmt.Errorf("field file not selected")
	}
	if width == "" {
		return fmt.Errorf("field width not filled")
	}
	for _, ch := range width {
		if ch < '0' || ch > '9' {
			return fmt.Errorf("width must contain only digits")
		}
	}
	return nil
}
