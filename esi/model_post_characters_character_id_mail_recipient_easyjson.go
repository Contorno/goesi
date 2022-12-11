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

func easyjsonFd1c367DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *PostCharactersCharacterIdMailRecipientList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCharactersCharacterIdMailRecipientList, 0, 2)
			} else {
				*out = PostCharactersCharacterIdMailRecipientList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCharactersCharacterIdMailRecipient
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
func easyjsonFd1c367EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in PostCharactersCharacterIdMailRecipientList) {
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
func (v PostCharactersCharacterIdMailRecipientList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFd1c367EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdMailRecipientList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFd1c367EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdMailRecipientList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFd1c367DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdMailRecipientList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFd1c367DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonFd1c367DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *PostCharactersCharacterIdMailRecipient) {
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
		case "recipient_id":
			out.RecipientId = int32(in.Int32())
		case "recipient_type":
			out.RecipientType = string(in.String())
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
func easyjsonFd1c367EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in PostCharactersCharacterIdMailRecipient) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RecipientId != 0 {
		const prefix string = ",\"recipient_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.RecipientId))
	}
	if in.RecipientType != "" {
		const prefix string = ",\"recipient_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCharactersCharacterIdMailRecipient) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFd1c367EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdMailRecipient) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFd1c367EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdMailRecipient) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFd1c367DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdMailRecipient) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFd1c367DecodeGithubComContornoGoesiEsi1(l, v)
}
