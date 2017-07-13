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

func easyjsonF7d26a12DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseGraphicsGraphicIdNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseGraphicsGraphicIdNotFoundList, 0, 4)
			} else {
				*out = GetUniverseGraphicsGraphicIdNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseGraphicsGraphicIdNotFound
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
func easyjsonF7d26a12EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseGraphicsGraphicIdNotFoundList) {
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
func (v GetUniverseGraphicsGraphicIdNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF7d26a12EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseGraphicsGraphicIdNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF7d26a12EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseGraphicsGraphicIdNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF7d26a12DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseGraphicsGraphicIdNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF7d26a12DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonF7d26a12DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseGraphicsGraphicIdNotFound) {
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
		key := in.UnsafeString()
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
func easyjsonF7d26a12EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseGraphicsGraphicIdNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"error\":")
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseGraphicsGraphicIdNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF7d26a12EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseGraphicsGraphicIdNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF7d26a12EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseGraphicsGraphicIdNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF7d26a12DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseGraphicsGraphicIdNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF7d26a12DecodeGithubComAntihaxGoesiEsi1(l, v)
}
