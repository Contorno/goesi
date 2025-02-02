// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson1a6ccbc6DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetContractsPublicBidsContractId200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetContractsPublicBidsContractId200OkList, 0, 2)
			} else {
				*out = GetContractsPublicBidsContractId200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetContractsPublicBidsContractId200Ok
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1a6ccbc6EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetContractsPublicBidsContractId200OkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetContractsPublicBidsContractId200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1a6ccbc6EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetContractsPublicBidsContractId200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1a6ccbc6EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetContractsPublicBidsContractId200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1a6ccbc6DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetContractsPublicBidsContractId200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1a6ccbc6DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson1a6ccbc6DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetContractsPublicBidsContractId200Ok) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "amount":
			out.Amount = float32(in.Float32())
		case "bid_id":
			out.BidId = int32(in.Int32())
		case "date_bid":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateBid).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1a6ccbc6EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetContractsPublicBidsContractId200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		first = false
		out.RawString(prefix[1:])
		out.Float32(float32(in.Amount))
	}
	if in.BidId != 0 {
		const prefix string = ",\"bid_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.BidId))
	}
	if true {
		const prefix string = ",\"date_bid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.DateBid).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetContractsPublicBidsContractId200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1a6ccbc6EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetContractsPublicBidsContractId200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1a6ccbc6EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetContractsPublicBidsContractId200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1a6ccbc6DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetContractsPublicBidsContractId200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1a6ccbc6DecodeGithubComContornoGoesiEsi1(l, v)
}
