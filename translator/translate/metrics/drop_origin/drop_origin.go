// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package drop_origin

import (
	"github.com/aws/amazon-cloudwatch-agent/translator"
	parent "github.com/aws/amazon-cloudwatch-agent/translator/translate/metrics"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/metrics/config"
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/metrics/metrics_collect"
)

type dropOrigin struct {
}

const SectionKey = "drop_origin"
const MetricName = "dropMetrics"

func (do *dropOrigin) ApplyRule(input interface{}) (returnKey string, returnVal interface{}) {
	im := input.(map[string]interface{})
	// result := map[string]string{}

	returnKey = ""
	returnVal = ""
	if _, ok := im[metrics_collect.SectionKey]; !ok {
		return
	} else {
		pluginMap := im[metrics_collect.SectionKey].(map[string]interface{})
		result := make(map[string][]string)

		for key, _ := range pluginMap {
			if _, isDropping := translator.DefaultCase(SectionKey, false, pluginMap[key]); !isDropping.(bool) {
				// Do nothing if not dropping origin
			} else {
				// put dropping metric name in the result
				returnKey = SectionKey
				result[MetricName] = append(result[MetricName], config.GetRealPluginName(key))
				returnVal = result
			}
		}
	}
	return
}

func init() {
	do := new(dropOrigin)
	parent.RegisterRule(SectionKey, do)
}
