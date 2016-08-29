package flat_test

import (
  "testing"
  "github.com/erichnascimento/go-flat-file"
)

func TestInvalidTagParsing(t *testing.T)  {
  start, finish, err := flat.ParseTag(``)
  if start != flat.InvalidIndex {
    t.Fatalf("start = %d, want %d", start, flat.InvalidIndex)
  }

  if finish != flat.InvalidIndex {
    t.Fatalf("finish = %d, want %d", finish, flat.InvalidIndex)
  }

  if err != flat.ErrInvalidIndex {
    t.Fatalf("err = %s, want %s", err, flat.ErrInvalidIndex)
  }
}



func TestTagParsing(t *testing.T)  {
  start, finish, err := flat.ParseTag(`2,20`)
  if err != nil {
    t.Fatalf("err = %s, want nil", err)
  }
  if start != 2 {
    t.Fatalf("start = %d, want 2", start)
  }

  if finish != 20 {
    t.Fatalf("finish = %d, want 20", finish)
  }
}
