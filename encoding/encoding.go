package encoding

type Codec interface {
    Marshal(v interface{}) ([]byte, error)
    Unmarshal(data []byte, v interface{}) error
    ContentType() string
}

var registry = make(map[string]Codec)

func Register(name string, c Codec) {
    registry[name] = c
}

func Get(name string) Codec {
    return registry[name]
}
