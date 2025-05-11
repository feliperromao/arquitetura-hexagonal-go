package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jsonError(t *testing.T) {
	var jsonData = "any error message"
	result := jsonError(jsonData)
	require.Equal(t, []byte(`{"message":"any error message"}`), result)
}