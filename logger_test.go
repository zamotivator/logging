package logging

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	http_request := GetLogger("http.request")
	assert.Equal(t, http_request.path, ParseName("http.request"))
	db := GetLogger("db")
	assert.Equal(t, db.path, ParseName("db"))
}
