package helper

import (
  "strings"
)

func Urlify(url []byte) string {
  url_string := strings.TrimSpace(string(url))
  return url_string[:len(url_string) - 4]
}

func FormatBrachName(branch_name []byte) string {
  return strings.TrimSpace(string(branch_name))
}
