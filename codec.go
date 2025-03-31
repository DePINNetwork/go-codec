// Package go_codec provides codec functionality with dual import support
package go_codec

import (
	"github.com/DePINNetwork/go-codec/codec"
)

// Re-export key types and functions
type MsgpackHandle = codec.MsgpackHandle
type BasicHandle = codec.BasicHandle
type Encoder = codec.Encoder

// NewEncoder creates a new encoder
func NewEncoder(w interface{}, h interface{}) *Encoder {
	return codec.NewEncoder(w, h.(*MsgpackHandle))
}

// Handle is the default MsgpackHandle
var Handle = &codec.MsgpackHandle{}
