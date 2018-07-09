package model

import (
	"testing"

	"github.com/alustau/issuer/database"
	"github.com/alustau/issuer/request"
	"github.com/stretchr/testify/assert"
)

func TestNewIssue(t *testing.T) {
	connection, _ := database.GetConnection("localhost", "3306", "root", "alustau321", "issuer_test")
	defer connection.Close()

	request := &request.Issue{
		Action: "reopened",
		Issue: map[string]interface{}{
			"number": 22,
			"labels": map[string]interface{}{
				"foo": "bar",
			},
		},
	}

	issue, _ := NewIssue(connection, request)

	// assert equality
	assert.Equal(t, request.Issue["number"], issue.Number)
}
