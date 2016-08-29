package flat_test

import (
  "fmt"
  "log"
  "strings"
  "io"
  "github.com/erichnascimento/go-flat-file"
)

type Person struct {
  Id string
  FirstName string `flat:"1,15"`
  LastName string `flat:"16,30"`
  Age string `flat:"31,33"`
  Birthdate string `flat:"34,43"`
  Salary string `flat:"44,55"`
  IsActive string `flat:"56,56"`
}

func (p Person) String() string {
  return fmt.Sprintf("%s %s, %s years old(%s). R$ %s (%s)",
    p.FirstName, p.LastName, p.Age, p.Birthdate, p.Salary, p.IsActive)
}

func ExampleReader() {
  in := `
Erich do       Nascimento     33 19-08-19831000000     s
Eduardo M. do  Nascimento     0  31-12-20150           s
`

  r := flat.NewReader(strings.NewReader(strings.TrimLeft(in, "\n")))

  for {
    person := &Person{}
    err := r.Read(person)
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }

    fmt.Printf("%s\n", person)
  }

  // Output:
  // Erich do Nascimento, 33 years old(19-08-1983). R$ 1000000 (s)
  // Eduardo M. do Nascimento, 0 years old(31-12-2015). R$ 0 (s)
}
