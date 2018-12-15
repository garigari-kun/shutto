package helper

import (
  "strings"
)

func Urlify(url []byte) string {
  return string(url[:len(string(url)) - 5])
}

func FormatBrachName(branch_name []byte) string {
  return strings.TrimSpace(string(branch_name))
}
