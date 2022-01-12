// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package drop_origin

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropOriginal(t *testing.T) {
	e := new(dropOrigin)
	//Check whether override default config
	var input interface{}
	err := json.Unmarshal([]byte(`{
	 			"metrics_collected": {
	 				"cpu": {
	 					"drop_origin": true,
	 					"measurement": [
	 						"cpu_usage_guest"
	 					]
	 				}
	 			}}`), &input)
	assert.NoError(t, err)
	actualKey, actualVal := e.ApplyRule(input)
	expectedKey := "drop_origin"
	expectedVal := map[string][]string{
		"dropMetrics": {"cpu"},
	}

	assert.Equal(t, expectedKey, actualKey)
	assert.Equal(t, expectedVal, actualVal)

}
