package greeting

import (
  "testing"
  "regexp"
)

func TestHelloGary(t *testing.T) {
  name := "gary"
  want := regexp.MustCompile(`\b` + name + `\b`)
  msg, err := Hello(name)
  if(err != nil || !want.MatchString(msg)) {
    t.Errorf(`Hello("gary") = %q, %v, want match for %#q, nil`, msg, err, want)
  }
}
