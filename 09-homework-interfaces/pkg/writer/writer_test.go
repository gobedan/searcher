package writer

import (
	"bytes"
	"testing"
)

func TestWriteStrings(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			"Test #1",
			args{[]any{123, "321"}},
			"321",
		},
		{
			"Test #2",
			args{[]any{"123", "321"}},
			"123321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			WriteStrings(w, tt.args.args...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("WriteStrings() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
