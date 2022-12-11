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

func easyjsonD15c35bDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdOpportunities200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdOpportunities200OkList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdOpportunities200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdOpportunities200Ok
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
func easyjsonD15c35bEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdOpportunities200OkList) {
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
func (v GetCharactersCharacterIdOpportunities200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD15c35bEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdOpportunities200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD15c35bEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdOpportunities200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD15c35bDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdOpportunities200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD15c35bDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonD15c35bDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdOpportunities200Ok) {
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
		case "completed_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CompletedAt).UnmarshalJSON(data))
			}
		case "task_id":
			out.TaskId = int32(in.Int32())
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
func easyjsonD15c35bEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdOpportunities200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"completed_at\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.CompletedAt).MarshalJSON())
	}
	if in.TaskId != 0 {
		const prefix string = ",\"task_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TaskId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdOpportunities200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD15c35bEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdOpportunities200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD15c35bEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdOpportunities200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD15c35bDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdOpportunities200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD15c35bDecodeGithubComContornoGoesiEsi1(l, v)
}
