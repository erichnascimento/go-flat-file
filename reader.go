package flat

import (
  "fmt"
  "io"
  "bytes"
  "bufio"
  "reflect"
  "strings"
)

// A ParseError is returned for parsing errors.
// The first line is one.
type ParseError struct {
  Line int
  Err error
}

func (e *ParseError) Error() string {
  return fmt.Sprintf("line %d: %s", e.Line, e.Err)
}

// A Reader reads records from a Flat-encoded file.
type Reader struct {
  // EOL is the end of line character.
  // It is set to new line char('\n') by NewReader
  EOL byte
  // Filler is the string used to fill fields in rows.
  // It is set to space by NewReader
  Filler string
  line int
  r *bufio.Reader
  field bytes.Buffer
  row string
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
  return &Reader{
    EOL: '\n',
    Filler: " ",
    r: bufio.NewReader(r),
  }
}

// error creates a new ParseError based on r.
func (r *Reader) error(err error) error {
  return &ParseError{
    Line: r.line,
    Err: err,
  }
}

// fillData fills the data in v
func (r *Reader) fillData(v interface{}) error {
  el := reflect.ValueOf(v).Elem()
  for _, field := range getFields(v, "flat") {
    if field.tag != true {
      continue
    }

    if field.finish >= len(r.row) {
      return r.error(fmt.Errorf(`The field "%s(%d:%d)" is out of range.`, field.name, field.start, field.finish))
    }

    data := strings.TrimRight(r.row[field.start-1:field.finish], r.Filler)
    el.FieldByName(field.name).SetString(data)
  }

  return nil
}

// Read reads one record from r. The record is sliced and
// v is filled.
func (r *Reader) Read(v interface{}) (err error) {
  if err := r.readRow(); err != nil {
    return err
  }
  return r.fillData(v)
}

// readRow read a row from r.
func (r *Reader) readRow() (err error) {
  r.line++

  r.row, err = r.r.ReadString(r.EOL)
  if err != nil && err != io.EOF {
    r.error(err)
  }

  return
}
