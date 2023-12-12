package utils

import (
	"io/ioutil"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
)

const defaultMetaPath = "/etc/mse/pod/labels"

// ExtractInstanceMeta extract the pod meta info.
func ExtractInstanceMeta() map[string]string {
	data, err := ioutil.ReadFile(defaultMetaPath)
	if err != nil {
		klog.Warnf("load metadata failed: %v", err)
	}
	tags := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		tags[parts[0]] = strings.ReplaceAll(parts[1], "\"", "")
	}
	return tags
}
