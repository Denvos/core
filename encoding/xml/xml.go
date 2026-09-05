package xml
import "github.com/Denvos/core/encoding"

import (
    "encoding/xml"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
    return xml.Marshal(v)
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
    return xml.Unmarshal(data, v)
}

func (c Codec) ContentType() string {
    return "application/xml"
}

func init() {
    encoding.Register("xml", Codec{})
}
