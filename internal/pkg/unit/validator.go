package unit

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Custom validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()

	// Register function to get tag name from json tags
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Validate validates a struct based on the validator tags and returns JSON error response
func Validate(obj interface{}) (bool, map[string]string) {
	err := validate.Struct(obj)
	if err == nil {
		return true, nil
	}

	errorMap := make(map[string]string)

	// Check for InvalidValidationError using errors.As
	var invalidErr *validator.InvalidValidationError
	if errors.As(err, &invalidErr) {
		errorMap["message"] = "Invalid validation error"
		return false, errorMap
	}

	// Check for ValidationErrors using errors.As
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, fieldErr := range validationErrs {
			// Use the field name from the JSON tag
			fieldName := fieldErr.Field()

			// Create custom error message based on the validation tag
			switch fieldErr.Tag() {
			case "required":
				errorMap[fieldName] = fieldName + " is required"
			case "email":
				errorMap[fieldName] = fieldName + " must be a valid email address"
			case "min":
				errorMap[fieldName] = fieldName + " must be at least " + fieldErr.Param() + " characters"
			case "max":
				errorMap[fieldName] = fieldName + " must be at most " + fieldErr.Param() + " characters"
			default:
				errorMap[fieldName] = "Validation failed on: " + fieldErr.Tag()
			}
		}
		return false, errorMap
	}

	// Handle other error types
	errorMap["message"] = "Unexpected error: " + err.Error()
	return false, errorMap
}

// ValidateJSON validates the JSON request body and responds with errors if validation fails
func ValidateJSON(c *gin.Context, obj interface{}) bool {
	// Bind JSON to struct
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return false
	}

	// Validate the struct
	valid, errors := Validate(obj)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Validation failed",
			"details": errors,
		})
		return false
	}

	return true
}
