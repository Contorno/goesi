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

func easyjson2533c13aDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *PostFleetsFleetIdMembersUnprocessableEntityList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostFleetsFleetIdMembersUnprocessableEntityList, 0, 4)
			} else {
				*out = PostFleetsFleetIdMembersUnprocessableEntityList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostFleetsFleetIdMembersUnprocessableEntity
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
func easyjson2533c13aEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in PostFleetsFleetIdMembersUnprocessableEntityList) {
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
func (v PostFleetsFleetIdMembersUnprocessableEntityList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2533c13aEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFleetsFleetIdMembersUnprocessableEntityList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2533c13aEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFleetsFleetIdMembersUnprocessableEntityList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2533c13aDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFleetsFleetIdMembersUnprocessableEntityList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2533c13aDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson2533c13aDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *PostFleetsFleetIdMembersUnprocessableEntity) {
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
func easyjson2533c13aEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in PostFleetsFleetIdMembersUnprocessableEntity) {
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
func (v PostFleetsFleetIdMembersUnprocessableEntity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2533c13aEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFleetsFleetIdMembersUnprocessableEntity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2533c13aEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFleetsFleetIdMembersUnprocessableEntity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2533c13aDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFleetsFleetIdMembersUnprocessableEntity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2533c13aDecodeGithubComContornoGoesiEsi1(l, v)
}
