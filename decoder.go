package amf

type Decoder interface {
	Decode() (interface{}, error)
}
