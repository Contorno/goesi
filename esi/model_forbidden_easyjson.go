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

func easyjsonAe9dec1fDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *ForbiddenList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ForbiddenList, 0, 2)
			} else {
				*out = ForbiddenList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Forbidden
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
func easyjsonAe9dec1fEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in ForbiddenList) {
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
func (v ForbiddenList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe9dec1fEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ForbiddenList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe9dec1fEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ForbiddenList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe9dec1fDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ForbiddenList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe9dec1fDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonAe9dec1fDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *Forbidden) {
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
		case "sso_status":
			out.SsoStatus = int32(in.Int32())
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
func easyjsonAe9dec1fEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in Forbidden) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		const prefix string = ",\"error\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Error_))
	}
	if in.SsoStatus != 0 {
		const prefix string = ",\"sso_status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SsoStatus))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Forbidden) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe9dec1fEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Forbidden) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe9dec1fEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Forbidden) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe9dec1fDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Forbidden) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe9dec1fDecodeGithubComContornoGoesiEsi1(l, v)
}
