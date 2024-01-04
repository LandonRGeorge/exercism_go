package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type caesar struct{}

type shift struct {
	distance int
}

type vigenere struct {
	key string
}
