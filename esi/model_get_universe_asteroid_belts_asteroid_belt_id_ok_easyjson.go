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

func easyjson1e086bfcDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetUniverseAsteroidBeltsAsteroidBeltIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseAsteroidBeltsAsteroidBeltIdOkList, 0, 1)
			} else {
				*out = GetUniverseAsteroidBeltsAsteroidBeltIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseAsteroidBeltsAsteroidBeltIdOk
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
func easyjson1e086bfcEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetUniverseAsteroidBeltsAsteroidBeltIdOkList) {
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
func (v GetUniverseAsteroidBeltsAsteroidBeltIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1e086bfcEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseAsteroidBeltsAsteroidBeltIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1e086bfcEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseAsteroidBeltsAsteroidBeltIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1e086bfcDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseAsteroidBeltsAsteroidBeltIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1e086bfcDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson1e086bfcDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetUniverseAsteroidBeltsAsteroidBeltIdOk) {
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
		case "name":
			out.Name = string(in.String())
		case "position":
			(out.Position).UnmarshalEasyJSON(in)
		case "system_id":
			out.SystemId = int32(in.Int32())
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
func easyjson1e086bfcEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetUniverseAsteroidBeltsAsteroidBeltIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if true {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Position).MarshalEasyJSON(out)
	}
	if in.SystemId != 0 {
		const prefix string = ",\"system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SystemId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseAsteroidBeltsAsteroidBeltIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1e086bfcEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseAsteroidBeltsAsteroidBeltIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1e086bfcEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseAsteroidBeltsAsteroidBeltIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1e086bfcDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseAsteroidBeltsAsteroidBeltIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1e086bfcDecodeGithubComContornoGoesiEsi1(l, v)
}
