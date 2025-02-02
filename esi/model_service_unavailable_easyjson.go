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

func easyjson5322fa8aDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *ServiceUnavailableList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ServiceUnavailableList, 0, 4)
			} else {
				*out = ServiceUnavailableList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 ServiceUnavailable
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
func easyjson5322fa8aEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in ServiceUnavailableList) {
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
func (v ServiceUnavailableList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5322fa8aEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceUnavailableList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5322fa8aEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceUnavailableList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5322fa8aDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceUnavailableList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5322fa8aDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson5322fa8aDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *ServiceUnavailable) {
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
func easyjson5322fa8aEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in ServiceUnavailable) {
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
func (v ServiceUnavailable) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5322fa8aEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceUnavailable) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5322fa8aEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceUnavailable) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5322fa8aDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceUnavailable) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5322fa8aDecodeGithubComContornoGoesiEsi1(l, v)
}
