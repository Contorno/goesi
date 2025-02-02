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

func easyjson30a3ef00DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdMailLabelsOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdMailLabelsOkList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdMailLabelsOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdMailLabelsOk
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
func easyjson30a3ef00EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdMailLabelsOkList) {
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
func (v GetCharactersCharacterIdMailLabelsOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30a3ef00EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMailLabelsOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30a3ef00EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLabelsOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30a3ef00DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLabelsOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30a3ef00DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson30a3ef00DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdMailLabelsOk) {
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
		case "labels":
			if in.IsNull() {
				in.Skip()
				out.Labels = nil
			} else {
				in.Delim('[')
				if out.Labels == nil {
					if !in.IsDelim(']') {
						out.Labels = make([]GetCharactersCharacterIdMailLabelsLabel, 0, 1)
					} else {
						out.Labels = []GetCharactersCharacterIdMailLabelsLabel{}
					}
				} else {
					out.Labels = (out.Labels)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdMailLabelsLabel
					(v4).UnmarshalEasyJSON(in)
					out.Labels = append(out.Labels, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "total_unread_count":
			out.TotalUnreadCount = int32(in.Int32())
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
func easyjson30a3ef00EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdMailLabelsOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Labels) != 0 {
		const prefix string = ",\"labels\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v5, v6 := range in.Labels {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.TotalUnreadCount != 0 {
		const prefix string = ",\"total_unread_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TotalUnreadCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdMailLabelsOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30a3ef00EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMailLabelsOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30a3ef00EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLabelsOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30a3ef00DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLabelsOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30a3ef00DecodeGithubComContornoGoesiEsi1(l, v)
}
