# ToYAML

## Usage
``` go
type bar struct {
  HogeFuga int
}
type example struct {
  Foo string
  Bar bar
}

func main() {
  b, _ := toml.Marshal(&example{Foo: "foo", Bar: bar{HogeFuga: 100}})
  yamlText, _ := ToYAML(b)

  fmt.Println(yamlText)

  // Bar:
  //   HogeFuga: 100
  // Foo: foo
}

```
