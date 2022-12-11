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

func easyjson87fa4a08DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *PutFleetsFleetIdSquadsSquadIdNamingList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PutFleetsFleetIdSquadsSquadIdNamingList, 0, 4)
			} else {
				*out = PutFleetsFleetIdSquadsSquadIdNamingList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PutFleetsFleetIdSquadsSquadIdNaming
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
func easyjson87fa4a08EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in PutFleetsFleetIdSquadsSquadIdNamingList) {
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
func (v PutFleetsFleetIdSquadsSquadIdNamingList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87fa4a08EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PutFleetsFleetIdSquadsSquadIdNamingList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87fa4a08EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PutFleetsFleetIdSquadsSquadIdNamingList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87fa4a08DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PutFleetsFleetIdSquadsSquadIdNamingList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87fa4a08DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson87fa4a08DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *PutFleetsFleetIdSquadsSquadIdNaming) {
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
		case "name":
			out.Name = string(in.String())
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
func easyjson87fa4a08EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in PutFleetsFleetIdSquadsSquadIdNaming) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PutFleetsFleetIdSquadsSquadIdNaming) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87fa4a08EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PutFleetsFleetIdSquadsSquadIdNaming) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87fa4a08EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PutFleetsFleetIdSquadsSquadIdNaming) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87fa4a08DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PutFleetsFleetIdSquadsSquadIdNaming) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87fa4a08DecodeGithubComContornoGoesiEsi1(l, v)
}
