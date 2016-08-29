package flat

import (
  "strings"
  "errors"
  "strconv"
)

var (
  ErrInvalidIndex = errors.New("invalid tag")
  InvalidIndex = -1
)

func ParseTag(tag string) (start, finish int, err error) {
  if idx := strings.Index(tag, ","); idx != -1 {
    if start, err = strconv.Atoi(tag[:idx]); err != nil {
      return InvalidIndex, InvalidIndex, err
    }

    if finish, err = strconv.Atoi(tag[idx+1:]); err != nil {
      return InvalidIndex, InvalidIndex, err
    }

    return
  }

  return -1, -1, ErrInvalidIndex
}
