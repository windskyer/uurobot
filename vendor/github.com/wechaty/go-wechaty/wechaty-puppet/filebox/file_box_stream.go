package filebox

import (
	"io"
)

type fileBoxStream struct {
	Reader io.Reader
}

func newFileBoxStream(reader io.Reader) *fileBoxStream {
	return &fileBoxStream{Reader: reader}
}

func (fb *fileBoxStream) toJSONMap() map[string]interface{} {
	return nil
}

func (fb *fileBoxStream) toBytes() ([]byte, error) {
	panic("im")
}

func (fb *fileBoxStream) toReader() (io.Reader, error) {
	return fb.Reader, nil
}
