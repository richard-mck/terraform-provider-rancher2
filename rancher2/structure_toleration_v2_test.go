package rancher2

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

var (
	testTolerationsV2Conf      []corev1.Toleration
	testTolerationsV2Interface []interface{}
)

func init() {
	seconds := int64(10)
	testTolerationsV2Conf = []corev1.Toleration{
		{
			Key:               "key",
			Value:             "value",
			Effect:            corev1.TaintEffectNoSchedule,
			Operator:          corev1.TolerationOpEqual,
			TolerationSeconds: &seconds,
		},
	}
	testTolerationsV2Interface = []interface{}{
		map[string]interface{}{
			"key":      "key",
			"value":    "value",
			"effect":   "NoSchedule",
			"operator": "Equal",
			"seconds":  10,
		},
	}
}

func TestFlattenTolerationsV2(t *testing.T) {

	cases := []struct {
		Input          []corev1.Toleration
		ExpectedOutput []interface{}
	}{
		{
			testTolerationsV2Conf,
			testTolerationsV2Interface,
		},
	}

	for _, tc := range cases {
		output := flattenTolerationsV2(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandTolerationsV2(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput []corev1.Toleration
	}{
		{
			testTolerationsV2Interface,
			testTolerationsV2Conf,
		},
	}

	for _, tc := range cases {
		output := expandTolerationsV2(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
