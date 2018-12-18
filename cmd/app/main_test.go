package main
import (
    "os"
    "testing"
	"github.com/stretchr/testify/assert"
)


func TestMain(t *testing.T) {}

func TestNonExistentCredentialsFile(t *testing.T) {
	os.Args = []string{"","-c ./non-exist-crednetials-file"}
	assert.Panics(t, main, "main should Panic with a non-existent credentials file")
}

