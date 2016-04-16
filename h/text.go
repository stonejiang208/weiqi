package h

import (
	"fmt"
	"net/http"
)

//Text
func text(w http.ResponseWriter, msg string, code int) {
	http.Error(w, msg, code)
}


//FormatStatusLine to a string
func formatStatusLine(code int) string {
	return fmt.Sprintf("%d %s", code, http.StatusText(code))
}

func textStatus(w http.ResponseWriter, msg string, code int) {
	text(w, fmt.Sprintf("%s\n%s", formatStatusLine(code), msg), code)
}