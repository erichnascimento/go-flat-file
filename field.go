package flat

import (
  "reflect"
)

type field struct {
  name string
  tag bool
  typ reflect.Type
  start int
  finish int
}

func getFields(v interface{}, tag string) (fields []*field) {
  t := reflect.TypeOf(v)
  if t.Kind() != reflect.Ptr {
    return
  }

  pt := t.Elem()
  fields = make([]*field, pt.NumField())
  for i := 0; i < pt.NumField(); i++ {
    sf := pt.Field(i)
    f := new(field)
    f.name = sf.Name
    f.typ = sf.Type

    start, finish, err := ParseTag(sf.Tag.Get(tag))
    f.start, f.finish, f.tag = start, finish, err == nil

    fields[i] = f
  }

  return
}
