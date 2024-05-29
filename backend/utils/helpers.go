package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/pjebs/jsonerror"
)

func GetLogger() *log.Logger {
	l := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	return l
}

func runCommand(command string, args ...string) ([]byte, error) {

	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}

	return output, nil
}

func ReadRequestBody(c *gin.Context, targetInterface interface{}) (interface{}, error) {
	var bd []byte

	// Read the request body
	bd, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, errors.New("failed to read request body")
	}
	defer c.Request.Body.Close()

	// Parse JSON data into SnappicFile struct
	err = json.Unmarshal(bd, &targetInterface)
	if err != nil {
		log.Println((err))
		return nil, errors.New("failed to parse JSON data")
	}

	return targetInterface, nil

}

// Generate secure token
func GenerateSecureToken(length int) (string, error) {
	// Create a byte slice to hold the random data
	token := make([]byte, length)

	// Read random bytes into the slice
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to a hexadecimal string
	return hex.EncodeToString(token), nil
}

// Build errors
type MultipleErrors []map[string]string

type JsonErrors struct {
	message string
	errors  []jsonerror.JE
}

type JsonErrorsString struct {
	Message string         `json:"message,omitempty"`
	Errors  MultipleErrors `json:"errors"`
}

func (e *JsonErrors) Add(err jsonerror.JE) {
	e.errors = append(e.errors, err)
}

func (e *JsonErrors) Message(message string) {
	e.message = message
}

func (e *JsonErrors) ToString() JsonErrorsString {
	var result MultipleErrors
	var ret JsonErrorsString

	for _, err := range e.errors {
		result = append(result, err.Render())
	}
	// Set errors
	ret.Errors = result
	ret.Message = e.message

	return ret
}
