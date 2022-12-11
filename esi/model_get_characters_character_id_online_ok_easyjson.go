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

func easyjson93c87c1aDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdOnlineOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdOnlineOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdOnlineOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdOnlineOk
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
func easyjson93c87c1aEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdOnlineOkList) {
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
func (v GetCharactersCharacterIdOnlineOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson93c87c1aEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdOnlineOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson93c87c1aEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdOnlineOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson93c87c1aDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdOnlineOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson93c87c1aDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson93c87c1aDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdOnlineOk) {
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
		case "last_login":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastLogin).UnmarshalJSON(data))
			}
		case "last_logout":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastLogout).UnmarshalJSON(data))
			}
		case "logins":
			out.Logins = int32(in.Int32())
		case "online":
			out.Online = bool(in.Bool())
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
func easyjson93c87c1aEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdOnlineOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"last_login\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.LastLogin).MarshalJSON())
	}
	if true {
		const prefix string = ",\"last_logout\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastLogout).MarshalJSON())
	}
	if in.Logins != 0 {
		const prefix string = ",\"logins\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Logins))
	}
	if in.Online {
		const prefix string = ",\"online\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Online))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdOnlineOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson93c87c1aEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdOnlineOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson93c87c1aEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdOnlineOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson93c87c1aDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdOnlineOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson93c87c1aDecodeGithubComContornoGoesiEsi1(l, v)
}
