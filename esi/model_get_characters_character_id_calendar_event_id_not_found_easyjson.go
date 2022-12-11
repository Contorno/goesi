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

func easyjson3d914852DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdCalendarEventIdNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdCalendarEventIdNotFoundList, 0, 4)
			} else {
				*out = GetCharactersCharacterIdCalendarEventIdNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdCalendarEventIdNotFound
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
func easyjson3d914852EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdCalendarEventIdNotFoundList) {
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
func (v GetCharactersCharacterIdCalendarEventIdNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3d914852EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdCalendarEventIdNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3d914852EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdCalendarEventIdNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3d914852DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdCalendarEventIdNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3d914852DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson3d914852DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdCalendarEventIdNotFound) {
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
func easyjson3d914852EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdCalendarEventIdNotFound) {
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
func (v GetCharactersCharacterIdCalendarEventIdNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3d914852EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdCalendarEventIdNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3d914852EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdCalendarEventIdNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3d914852DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdCalendarEventIdNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3d914852DecodeGithubComContornoGoesiEsi1(l, v)
}
