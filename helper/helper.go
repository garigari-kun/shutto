package helper

import (
  "strings"
)

func Urlify(url []byte) string {
  url_string := strings.TrimSpace(string(url))

  if strings.Contains(url_string, "ssh://") {
    domain := url_string[10:len(url) - 4]
    return "https://" + domain
  } else if strings.Contains(url_string, "git@") {
    domain := url_string[4:]
    url_parts := strings.Split(domain, ":")
    trailing_path := url_parts[1]
    return "https://" + url_parts[0] + "/" + trailing_path[:len(trailing_path) - 4]
  }
  return url_string[:len(url_string) - 4]

}

func FormatBrachName(branch_name []byte) string {
  return strings.TrimSpace(string(branch_name))
}
