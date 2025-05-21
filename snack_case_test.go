package str

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToSnackCase(t *testing.T) {
	tests := []struct {
		str   string
		want  string
		upper bool
	}{
		{str: "SomeStr", want: "SOME_STR", upper: true},
		{str: "", want: "", upper: true},
		{str: "F", want: "F", upper: true},
		{str: "foo", want: "FOO", upper: true},
		{str: "FooBar", want: "FOO_BAR", upper: true},
		{str: "fooBarBaz", want: "FOO_BAR_BAZ", upper: true},
		{str: "fooBar_baz", want: "FOO_BAR_BAZ", upper: true},
		{str: " foo_bar\n", want: "FOO_BAR", upper: true},
		{str: " foo-bar\t", want: "FOO_BAR", upper: true},
		{str: " foo bar\r", want: "FOO_BAR", upper: true},
		{str: "HTTP_status_code", want: "HTTP_STATUS_CODE", upper: true},
		{str: "foo_bar_v2_baz", want: "FOO_BAR_V2_BAZ", upper: true},
		{str: "fooBarV2Baz", want: "FOO_BAR_V2_BAZ", upper: true},

		{str: "some_str", want: "some_str", upper: false},
		{str: "", want: "", upper: false},
		{str: "F", want: "f", upper: false},
		{str: "foo", want: "foo", upper: false},
		{str: "FooBar", want: "foo_bar", upper: false},
		{str: "fooBarBaz", want: "foo_bar_baz", upper: false},
		{str: "fooBar_baz", want: "foo_bar_baz", upper: false},
		{str: " foo_bar\n", want: "foo_bar", upper: false},
		{str: " foo-bar\t", want: "foo_bar", upper: false},
		{str: " foo bar\r", want: "foo_bar", upper: false},
		{str: "HTTP_status_code", want: "http_status_code", upper: false},
		{str: "foo_bar_v2_baz", want: "foo_bar_v2_baz", upper: false},
		{str: "fooBarV2Baz", want: "foo_bar_v2_baz", upper: false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			got := ToSnackCase(tt.str, tt.upper)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ToSnackCase() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		str   string
		want  string
		upper bool
	}{
		{str: "SomeStr", want: "SOME-STR", upper: true},
		{str: "", want: "", upper: true},
		{str: "F", want: "F", upper: true},
		{str: "foo", want: "FOO", upper: true},
		{str: "FooBar", want: "FOO-BAR", upper: true},
		{str: "fooBarBaz", want: "FOO-BAR-BAZ", upper: true},
		{str: "fooBar-baz", want: "FOO-BAR-BAZ", upper: true},
		{str: " foo-bar\n", want: "FOO-BAR", upper: true},
		{str: " foo-bar\t", want: "FOO-BAR", upper: true},
		{str: " foo bar\r", want: "FOO-BAR", upper: true},
		{str: "HTTP-status-code", want: "HTTP-STATUS-CODE", upper: true},
		{str: "foo-bar-v2-baz", want: "FOO-BAR-V2-BAZ", upper: true},
		{str: "fooBarV2Baz", want: "FOO-BAR-V2-BAZ", upper: true},

		{str: "some-str", want: "some-str", upper: false},
		{str: "", want: "", upper: false},
		{str: "F", want: "f", upper: false},
		{str: "foo", want: "foo", upper: false},
		{str: "FooBar", want: "foo-bar", upper: false},
		{str: "fooBarBaz", want: "foo-bar-baz", upper: false},
		{str: "fooBar-baz", want: "foo-bar-baz", upper: false},
		{str: " foo-bar\n", want: "foo-bar", upper: false},
		{str: " foo-bar\t", want: "foo-bar", upper: false},
		{str: " foo bar\r", want: "foo-bar", upper: false},
		{str: "HTTP-status-code", want: "http-status-code", upper: false},
		{str: "foo-bar-v2-baz", want: "foo-bar-v2-baz", upper: false},
		{str: "fooBarV2Baz", want: "foo-bar-v2-baz", upper: false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			got := ToKebabCase(tt.str, tt.upper)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ToKebabCase() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
