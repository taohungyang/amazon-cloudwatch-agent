// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package cloudwatch

type DropOriginConfig struct {
	MetricNames []string `toml:"dropMetrics"`
}

func GetDroppingOriginMetrics(dropOriginConfig DropOriginConfig) map[string]struct{} {
	result := make(map[string]struct{})

	for _, metricName := range dropOriginConfig.MetricNames {
		result[metricName] = struct{}{}
	}
	return result
}