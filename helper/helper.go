package helper

func Urlify(url []byte) string {
  return string(url[:len(string(url)) - 5])
}
