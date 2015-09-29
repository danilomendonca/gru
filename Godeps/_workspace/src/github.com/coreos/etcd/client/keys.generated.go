// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package client

import (
	"errors"
	"fmt"
	codec1978 "github.com/elleFlorio/gru/Godeps/_workspace/src/github.com/ugorji/go/codec"
	"reflect"
	"runtime"
	"time"
)

const (
	codecSelferC_UTF81819         = 1
	codecSelferC_RAW1819          = 0
	codecSelverValueTypeArray1819 = 10
	codecSelverValueTypeMap1819   = 9
)

var (
	codecSelferBitsize1819                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1819 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1819 struct{}

func init() {
	if codec1978.GenVersion != 2 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			2, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 time.Time
		_ = v0
	}
}

func (x *Response) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yysep1 := !z.EncBinary()
		yy2arr1 := z.EncBasicHandle().StructToArray
		var yyfirst1 bool
		var yyq1 [3]bool
		_, _, _, _ = yysep1, yyfirst1, yyq1, yy2arr1
		const yyr1 bool = false
		if yyr1 || yy2arr1 {
			r.EncodeArrayStart(3)
		} else {
			var yynn1 int = 3
			for _, b := range yyq1 {
				if b {
					yynn1++
				}
			}
			r.EncodeMapStart(yynn1)
		}
		if yyr1 || yy2arr1 {
			r.EncodeString(codecSelferC_UTF81819, string(x.Action))
		} else {
			yyfirst1 = true
			r.EncodeString(codecSelferC_UTF81819, string("action"))
			if yysep1 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeString(codecSelferC_UTF81819, string(x.Action))
		}
		if yyr1 || yy2arr1 {
			if yysep1 {
				r.EncodeArrayEntrySeparator()
			}
			if x.Node == nil {
				r.EncodeNil()
			} else {
				x.Node.CodecEncodeSelf(e)
			}
		} else {
			if yyfirst1 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst1 = true
			}
			r.EncodeString(codecSelferC_UTF81819, string("node"))
			if yysep1 {
				r.EncodeMapKVSeparator()
			}
			if x.Node == nil {
				r.EncodeNil()
			} else {
				x.Node.CodecEncodeSelf(e)
			}
		}
		if yyr1 || yy2arr1 {
			if yysep1 {
				r.EncodeArrayEntrySeparator()
			}
			if x.PrevNode == nil {
				r.EncodeNil()
			} else {
				x.PrevNode.CodecEncodeSelf(e)
			}
		} else {
			if yyfirst1 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst1 = true
			}
			r.EncodeString(codecSelferC_UTF81819, string("prevNode"))
			if yysep1 {
				r.EncodeMapKVSeparator()
			}
			if x.PrevNode == nil {
				r.EncodeNil()
			} else {
				x.PrevNode.CodecEncodeSelf(e)
			}
		}
		if yysep1 {
			if yyr1 || yy2arr1 {
				r.EncodeArrayEnd()
			} else {
				r.EncodeMapEnd()
			}
		}
	}
}

func (x *Response) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if r.IsContainerType(codecSelverValueTypeMap1819) {
		yyl5 := r.ReadMapStart()
		if yyl5 == 0 {
			r.ReadMapEnd()
		} else {
			x.codecDecodeSelfFromMap(yyl5, d)
		}
	} else if r.IsContainerType(codecSelverValueTypeArray1819) {
		yyl5 := r.ReadArrayStart()
		if yyl5 == 0 {
			r.ReadArrayEnd()
		} else {
			x.codecDecodeSelfFromArray(yyl5, d)
		}
	} else {
		panic(codecSelferOnlyMapOrArrayEncodeToStructErr1819)
	}
}

func (x *Response) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys6Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys6Slc
	var yyhl6 bool = l >= 0
	for yyj6 := 0; ; yyj6++ {
		if yyhl6 {
			if yyj6 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
			if yyj6 > 0 {
				r.ReadMapEntrySeparator()
			}
		}
		yys6Slc = r.DecodeBytes(yys6Slc, true, true)
		yys6 := string(yys6Slc)
		if !yyhl6 {
			r.ReadMapKVSeparator()
		}
		switch yys6 {
		case "action":
			if r.TryDecodeAsNil() {
				x.Action = ""
			} else {
				x.Action = string(r.DecodeString())
			}
		case "node":
			if r.TryDecodeAsNil() {
				if x.Node != nil {
					x.Node = nil
				}
			} else {
				if x.Node == nil {
					x.Node = new(Node)
				}
				x.Node.CodecDecodeSelf(d)
			}
		case "prevNode":
			if r.TryDecodeAsNil() {
				if x.PrevNode != nil {
					x.PrevNode = nil
				}
			} else {
				if x.PrevNode == nil {
					x.PrevNode = new(Node)
				}
				x.PrevNode.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys6)
		} // end switch yys6
	} // end for yyj6
	if !yyhl6 {
		r.ReadMapEnd()
	}
}

func (x *Response) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj10 int
	var yyb10 bool
	var yyhl10 bool = l >= 0
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Action = ""
	} else {
		x.Action = string(r.DecodeString())
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		if x.Node != nil {
			x.Node = nil
		}
	} else {
		if x.Node == nil {
			x.Node = new(Node)
		}
		x.Node.CodecDecodeSelf(d)
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		if x.PrevNode != nil {
			x.PrevNode = nil
		}
	} else {
		if x.PrevNode == nil {
			x.PrevNode = new(Node)
		}
		x.PrevNode.CodecDecodeSelf(d)
	}
	for {
		yyj10++
		if yyhl10 {
			yyb10 = yyj10 > l
		} else {
			yyb10 = r.CheckBreak()
		}
		if yyb10 {
			break
		}
		if yyj10 > 1 {
			r.ReadArrayEntrySeparator()
		}
		z.DecStructFieldNotFound(yyj10-1, "")
	}
	r.ReadArrayEnd()
}

func (x *Node) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yysep14 := !z.EncBinary()
		yy2arr14 := z.EncBasicHandle().StructToArray
		var yyfirst14 bool
		var yyq14 [8]bool
		_, _, _, _ = yysep14, yyfirst14, yyq14, yy2arr14
		const yyr14 bool = false
		yyq14[1] = x.Dir != false
		yyq14[6] = x.Expiration != nil
		yyq14[7] = x.TTL != 0
		if yyr14 || yy2arr14 {
			r.EncodeArrayStart(8)
		} else {
			var yynn14 int = 5
			for _, b := range yyq14 {
				if b {
					yynn14++
				}
			}
			r.EncodeMapStart(yynn14)
		}
		if yyr14 || yy2arr14 {
			r.EncodeString(codecSelferC_UTF81819, string(x.Key))
		} else {
			yyfirst14 = true
			r.EncodeString(codecSelferC_UTF81819, string("key"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeString(codecSelferC_UTF81819, string(x.Key))
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq14[1] {
				r.EncodeBool(bool(x.Dir))
			} else {
				r.EncodeBool(false)
			}
		} else {
			if yyq14[1] {
				if yyfirst14 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst14 = true
				}
				r.EncodeString(codecSelferC_UTF81819, string("dir"))
				if yysep14 {
					r.EncodeMapKVSeparator()
				}
				r.EncodeBool(bool(x.Dir))
			}
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			r.EncodeString(codecSelferC_UTF81819, string(x.Value))
		} else {
			if yyfirst14 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst14 = true
			}
			r.EncodeString(codecSelferC_UTF81819, string("value"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeString(codecSelferC_UTF81819, string(x.Value))
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			if x.Nodes == nil {
				r.EncodeNil()
			} else {
				x.Nodes.CodecEncodeSelf(e)
			}
		} else {
			if yyfirst14 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst14 = true
			}
			r.EncodeString(codecSelferC_UTF81819, string("nodes"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			if x.Nodes == nil {
				r.EncodeNil()
			} else {
				x.Nodes.CodecEncodeSelf(e)
			}
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			r.EncodeUint(uint64(x.CreatedIndex))
		} else {
			if yyfirst14 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst14 = true
			}
			r.EncodeString(codecSelferC_UTF81819, string("createdIndex"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeUint(uint64(x.CreatedIndex))
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			r.EncodeUint(uint64(x.ModifiedIndex))
		} else {
			if yyfirst14 {
				r.EncodeMapEntrySeparator()
			} else {
				yyfirst14 = true
			}
			r.EncodeString(codecSelferC_UTF81819, string("modifiedIndex"))
			if yysep14 {
				r.EncodeMapKVSeparator()
			}
			r.EncodeUint(uint64(x.ModifiedIndex))
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq14[6] {
				if x.Expiration == nil {
					r.EncodeNil()
				} else {
					z.EncFallback(x.Expiration)
				}
			} else {
				r.EncodeNil()
			}
		} else {
			if yyq14[6] {
				if yyfirst14 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst14 = true
				}
				r.EncodeString(codecSelferC_UTF81819, string("expiration"))
				if yysep14 {
					r.EncodeMapKVSeparator()
				}
				if x.Expiration == nil {
					r.EncodeNil()
				} else {
					z.EncFallback(x.Expiration)
				}
			}
		}
		if yyr14 || yy2arr14 {
			if yysep14 {
				r.EncodeArrayEntrySeparator()
			}
			if yyq14[7] {
				r.EncodeInt(int64(x.TTL))
			} else {
				r.EncodeInt(0)
			}
		} else {
			if yyq14[7] {
				if yyfirst14 {
					r.EncodeMapEntrySeparator()
				} else {
					yyfirst14 = true
				}
				r.EncodeString(codecSelferC_UTF81819, string("ttl"))
				if yysep14 {
					r.EncodeMapKVSeparator()
				}
				r.EncodeInt(int64(x.TTL))
			}
		}
		if yysep14 {
			if yyr14 || yy2arr14 {
				r.EncodeArrayEnd()
			} else {
				r.EncodeMapEnd()
			}
		}
	}
}

func (x *Node) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if r.IsContainerType(codecSelverValueTypeMap1819) {
		yyl23 := r.ReadMapStart()
		if yyl23 == 0 {
			r.ReadMapEnd()
		} else {
			x.codecDecodeSelfFromMap(yyl23, d)
		}
	} else if r.IsContainerType(codecSelverValueTypeArray1819) {
		yyl23 := r.ReadArrayStart()
		if yyl23 == 0 {
			r.ReadArrayEnd()
		} else {
			x.codecDecodeSelfFromArray(yyl23, d)
		}
	} else {
		panic(codecSelferOnlyMapOrArrayEncodeToStructErr1819)
	}
}

func (x *Node) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys24Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys24Slc
	var yyhl24 bool = l >= 0
	for yyj24 := 0; ; yyj24++ {
		if yyhl24 {
			if yyj24 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
			if yyj24 > 0 {
				r.ReadMapEntrySeparator()
			}
		}
		yys24Slc = r.DecodeBytes(yys24Slc, true, true)
		yys24 := string(yys24Slc)
		if !yyhl24 {
			r.ReadMapKVSeparator()
		}
		switch yys24 {
		case "key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = string(r.DecodeString())
			}
		case "dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = bool(r.DecodeBool())
			}
		case "value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				x.Value = string(r.DecodeString())
			}
		case "nodes":
			if r.TryDecodeAsNil() {
				x.Nodes = nil
			} else {
				yyv28 := &x.Nodes
				yyv28.CodecDecodeSelf(d)
			}
		case "createdIndex":
			if r.TryDecodeAsNil() {
				x.CreatedIndex = 0
			} else {
				x.CreatedIndex = uint64(r.DecodeUint(64))
			}
		case "modifiedIndex":
			if r.TryDecodeAsNil() {
				x.ModifiedIndex = 0
			} else {
				x.ModifiedIndex = uint64(r.DecodeUint(64))
			}
		case "expiration":
			if r.TryDecodeAsNil() {
				if x.Expiration != nil {
					x.Expiration = nil
				}
			} else {
				if x.Expiration == nil {
					x.Expiration = new(time.Time)
				}
				z.DecFallback(x.Expiration, false)
			}
		case "ttl":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				x.TTL = int64(r.DecodeInt(64))
			}
		default:
			z.DecStructFieldNotFound(-1, yys24)
		} // end switch yys24
	} // end for yyj24
	if !yyhl24 {
		r.ReadMapEnd()
	}
}

func (x *Node) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj33 int
	var yyb33 bool
	var yyhl33 bool = l >= 0
	yyj33++
	if yyhl33 {
		yyb33 = yyj33 > l
	} else {
		yyb33 = r.CheckBreak()
	}
	if yyb33 {
		r.ReadArrayEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		x.Key = string(r.DecodeString())
	}
	yyj33++
	if yyhl33 {
		yyb33 = yyj33 > l
	} else {
		yyb33 = r.CheckBreak()
	}
	if yyb33 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		x.Dir = bool(r.DecodeBool())
	}
	yyj33++
	if yyhl33 {
		yyb33 = yyj33 > l
	} else {
		yyb33 = r.CheckBreak()
	}
	if yyb33 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		x.Value = string(r.DecodeString())
	}
	yyj33++
	if yyhl33 {
		yyb33 = yyj33 > l
	} else {
		yyb33 = r.CheckBreak()
	}
	if yyb33 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.Nodes = nil
	} else {
		yyv37 := &x.Nodes
		yyv37.CodecDecodeSelf(d)
	}
	yyj33++
	if yyhl33 {
		yyb33 = yyj33 > l
	} else {
		yyb33 = r.CheckBreak()
	}
	if yyb33 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.CreatedIndex = 0
	} else {
		x.CreatedIndex = uint64(r.DecodeUint(64))
	}
	yyj33++
	if yyhl33 {
		yyb33 = yyj33 > l
	} else {
		yyb33 = r.CheckBreak()
	}
	if yyb33 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.ModifiedIndex = 0
	} else {
		x.ModifiedIndex = uint64(r.DecodeUint(64))
	}
	yyj33++
	if yyhl33 {
		yyb33 = yyj33 > l
	} else {
		yyb33 = r.CheckBreak()
	}
	if yyb33 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		if x.Expiration != nil {
			x.Expiration = nil
		}
	} else {
		if x.Expiration == nil {
			x.Expiration = new(time.Time)
		}
		z.DecFallback(x.Expiration, false)
	}
	yyj33++
	if yyhl33 {
		yyb33 = yyj33 > l
	} else {
		yyb33 = r.CheckBreak()
	}
	if yyb33 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayEntrySeparator()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		x.TTL = int64(r.DecodeInt(64))
	}
	for {
		yyj33++
		if yyhl33 {
			yyb33 = yyj33 > l
		} else {
			yyb33 = r.CheckBreak()
		}
		if yyb33 {
			break
		}
		if yyj33 > 1 {
			r.ReadArrayEntrySeparator()
		}
		z.DecStructFieldNotFound(yyj33-1, "")
	}
	r.ReadArrayEnd()
}

func (x Nodes) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		h.encNodes((Nodes)(x), e)
	}
}

func (x *Nodes) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	h.decNodes((*Nodes)(x), d)
}

func (x codecSelfer1819) encNodes(v Nodes, e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	yys42 := !z.EncBinary()
	if yys42 {
		for yyi42, yyv42 := range v {
			if yyi42 > 0 {
				r.EncodeArrayEntrySeparator()
			}
			if yyv42 == nil {
				r.EncodeNil()
			} else {
				yyv42.CodecEncodeSelf(e)
			}
		}
		r.EncodeArrayEnd()
	} else {
		for _, yyv42 := range v {
			if yyv42 == nil {
				r.EncodeNil()
			} else {
				yyv42.CodecEncodeSelf(e)
			}
		}
	}
}

func (x codecSelfer1819) decNodes(v *Nodes, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv43 := *v
	yyh43, yyl43 := z.DecSliceHelperStart()

	var yyc43 bool
	_ = yyc43

	if yyv43 == nil {
		if yyl43 <= 0 {
			yyv43 = make(Nodes, 0)
		} else {
			yyv43 = make(Nodes, yyl43)
		}
		yyc43 = true
	}

	if yyl43 == 0 {
		if len(yyv43) != 0 {
			yyv43 = yyv43[:0]
			yyc43 = true
		}
	} else if yyl43 > 0 {

		yyn43 := yyl43
		if yyl43 > cap(yyv43) {
			yyv43 = make([]*Node, yyl43, yyl43)
			yyc43 = true

		} else if yyl43 != len(yyv43) {
			yyv43 = yyv43[:yyl43]
			yyc43 = true
		}
		yyj43 := 0
		for ; yyj43 < yyn43; yyj43++ {
			if r.TryDecodeAsNil() {
				if yyv43[yyj43] != nil {
					*yyv43[yyj43] = Node{}
				}
			} else {
				if yyv43[yyj43] == nil {
					yyv43[yyj43] = new(Node)
				}
				yyw44 := yyv43[yyj43]
				yyw44.CodecDecodeSelf(d)
			}

		}

	} else {
		for yyj43 := 0; !r.CheckBreak(); yyj43++ {
			if yyj43 >= len(yyv43) {
				yyv43 = append(yyv43, nil) // var yyz43 *Node
				yyc43 = true
			}
			if yyj43 > 0 {
				yyh43.Sep(yyj43)
			}

			if yyj43 < len(yyv43) {
				if r.TryDecodeAsNil() {
					if yyv43[yyj43] != nil {
						*yyv43[yyj43] = Node{}
					}
				} else {
					if yyv43[yyj43] == nil {
						yyv43[yyj43] = new(Node)
					}
					yyw45 := yyv43[yyj43]
					yyw45.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		yyh43.End()
	}
	if yyc43 {
		*v = yyv43
	}

}
