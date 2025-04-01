package codec

import (

	"testing"
)

func TestJsonMapKeysAsStringCanonical(t *testing.T) {
	handle := new(JsonHandle)
	handle.MapKeyAsString = true
	handle.Canonical = true

	var b []byte
	enc := NewEncoderBytes(&b, handle)

	mapType := map[interface{}]interface{}{
		2:  "interface{} key",
		11: false,
	}
	enc.MustEncode(mapType)

	// We expect "11" to be encoded before "2" because its string representation
	// is lexicographically less than "2".
	expected := `{"11":false,"2":"interface{} key"}`

	if string(b) != expected {
		t.Errorf("unexpected encoding: `%s`", b)
	}
}
