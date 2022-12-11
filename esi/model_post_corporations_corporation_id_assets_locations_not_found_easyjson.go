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

func easyjsonC2effe49DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *PostCorporationsCorporationIdAssetsLocationsNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCorporationsCorporationIdAssetsLocationsNotFoundList, 0, 4)
			} else {
				*out = PostCorporationsCorporationIdAssetsLocationsNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCorporationsCorporationIdAssetsLocationsNotFound
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
func easyjsonC2effe49EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in PostCorporationsCorporationIdAssetsLocationsNotFoundList) {
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
func (v PostCorporationsCorporationIdAssetsLocationsNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC2effe49EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCorporationsCorporationIdAssetsLocationsNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC2effe49EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCorporationsCorporationIdAssetsLocationsNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC2effe49DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCorporationsCorporationIdAssetsLocationsNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC2effe49DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonC2effe49DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *PostCorporationsCorporationIdAssetsLocationsNotFound) {
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
func easyjsonC2effe49EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in PostCorporationsCorporationIdAssetsLocationsNotFound) {
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
func (v PostCorporationsCorporationIdAssetsLocationsNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC2effe49EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCorporationsCorporationIdAssetsLocationsNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC2effe49EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCorporationsCorporationIdAssetsLocationsNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC2effe49DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCorporationsCorporationIdAssetsLocationsNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC2effe49DecodeGithubComContornoGoesiEsi1(l, v)
}
