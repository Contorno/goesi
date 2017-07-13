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

func easyjsonD9f0de9eDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannelsAllowedList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdChatChannelsAllowedList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdChatChannelsAllowedList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdChatChannelsAllowed
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
func easyjsonD9f0de9eEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdChatChannelsAllowedList) {
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
func (v GetCharactersCharacterIdChatChannelsAllowedList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD9f0de9eEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdChatChannelsAllowedList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD9f0de9eEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannelsAllowedList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD9f0de9eDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannelsAllowedList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD9f0de9eDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonD9f0de9eDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdChatChannelsAllowed) {
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
		case "accessor_id":
			out.AccessorId = int32(in.Int32())
		case "accessor_type":
			out.AccessorType = string(in.String())
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
func easyjsonD9f0de9eEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdChatChannelsAllowed) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AccessorId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_id\":")
		out.Int32(int32(in.AccessorId))
	}
	if in.AccessorType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"accessor_type\":")
		out.String(string(in.AccessorType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdChatChannelsAllowed) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD9f0de9eEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdChatChannelsAllowed) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD9f0de9eEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannelsAllowed) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD9f0de9eDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdChatChannelsAllowed) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD9f0de9eDecodeGithubComAntihaxGoesiEsi1(l, v)
}
