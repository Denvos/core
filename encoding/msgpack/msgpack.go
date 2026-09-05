package msgpack
import "github.com/Denvos/core/encoding"

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
    return nil, nil
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
    return nil
}

func (c Codec) ContentType() string {
    return "application/msgpack"
}

func init() {
    encoding.Register("msgpack", Codec{})
}
