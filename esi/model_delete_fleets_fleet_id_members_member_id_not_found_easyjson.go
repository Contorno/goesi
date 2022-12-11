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

func easyjsonCf3a58dcDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *DeleteFleetsFleetIdMembersMemberIdNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(DeleteFleetsFleetIdMembersMemberIdNotFoundList, 0, 4)
			} else {
				*out = DeleteFleetsFleetIdMembersMemberIdNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 DeleteFleetsFleetIdMembersMemberIdNotFound
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
func easyjsonCf3a58dcEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in DeleteFleetsFleetIdMembersMemberIdNotFoundList) {
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
func (v DeleteFleetsFleetIdMembersMemberIdNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3a58dcEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteFleetsFleetIdMembersMemberIdNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3a58dcEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteFleetsFleetIdMembersMemberIdNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3a58dcDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteFleetsFleetIdMembersMemberIdNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3a58dcDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonCf3a58dcDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *DeleteFleetsFleetIdMembersMemberIdNotFound) {
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
		case "error":
			out.Error_ = string(in.String())
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
func easyjsonCf3a58dcEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in DeleteFleetsFleetIdMembersMemberIdNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		const prefix string = ",\"error\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteFleetsFleetIdMembersMemberIdNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3a58dcEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteFleetsFleetIdMembersMemberIdNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3a58dcEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteFleetsFleetIdMembersMemberIdNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3a58dcDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteFleetsFleetIdMembersMemberIdNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3a58dcDecodeGithubComContornoGoesiEsi1(l, v)
}
