package helper

import (
  "testing"
)

func TestUrlify(t *testing.T) {
  expected := "https://github.com/garigari-kun/shutto"
  // git config: https://github.com/garigari-kun/shutto.git
  test_bytes := []byte{0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x69, 0x67, 0x61, 0x72, 0x69, 0x2d, 0x6b, 0x75, 0x6e, 0x2f, 0x73, 0x68, 0x75, 0x74, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74}
  urlified := Urlify(test_bytes)
  if urlified != expected {
    t.Fatal("Failed")
  }

  // ssh://git@github.com/garigari-kun/shutto.git
  test_bytes_ssh := []byte{0x73, 0x73, 0x68, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x69, 0x67, 0x61, 0x72, 0x69, 0x2d, 0x6b, 0x75, 0x6e, 0x2f, 0x73, 0x68, 0x75, 0x74, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74}
  urlified_ssh := Urlify(test_bytes_ssh)
  if urlified_ssh != expected {
    t.Fatal("Failed")
  }

  // git@github.com:garigari-kun/shutto.git
  test_bytes_git := []byte{0x67, 0x69, 0x74, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x3a, 0x67, 0x61, 0x72, 0x69, 0x67, 0x61, 0x72, 0x69, 0x2d, 0x6b, 0x75, 0x6e, 0x2f, 0x73, 0x68, 0x75, 0x74, 0x74, 0x6f, 0x2e, 0x67, 0x69, 0x74}
  urlified_git := Urlify(test_bytes_git)
  if urlified_git != expected {
    t.Fatal("Failed")
  }
}

func TestFormatBranchName(t *testing.T) {
  test_bytes := []byte{0x20, 0x74, 0x65, 0x73, 0x74, 0x20}
  formated := FormatBrachName(test_bytes)
  if formated != "test" {
    t.Fatal("Failed. Should be 'test'")
  }
}
