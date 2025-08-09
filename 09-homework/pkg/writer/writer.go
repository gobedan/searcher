package writer

import "io"

func WriteStrings(w io.Writer, args ...any) {
	for _, arg := range args {
		if str, ok := arg.(string); ok {
			w.Write([]byte(str))
		}
	}
}
