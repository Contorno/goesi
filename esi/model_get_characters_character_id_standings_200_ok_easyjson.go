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

func easyjsonCd2065b3DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdStandings200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdStandings200OkList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdStandings200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdStandings200Ok
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
func easyjsonCd2065b3EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdStandings200OkList) {
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
func (v GetCharactersCharacterIdStandings200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCd2065b3EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStandings200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCd2065b3EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStandings200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCd2065b3DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStandings200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCd2065b3DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonCd2065b3DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdStandings200Ok) {
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
		case "from_id":
			out.FromId = int32(in.Int32())
		case "from_type":
			out.FromType = string(in.String())
		case "standing":
			out.Standing = float32(in.Float32())
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
func easyjsonCd2065b3EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdStandings200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FromId != 0 {
		const prefix string = ",\"from_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.FromId))
	}
	if in.FromType != "" {
		const prefix string = ",\"from_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FromType))
	}
	if in.Standing != 0 {
		const prefix string = ",\"standing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Standing))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdStandings200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCd2065b3EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStandings200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCd2065b3EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStandings200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCd2065b3DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStandings200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCd2065b3DecodeGithubComContornoGoesiEsi1(l, v)
}
