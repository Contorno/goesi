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

func easyjson460d9473DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetUniverseTypesTypeIdDogmaAttributeList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseTypesTypeIdDogmaAttributeList, 0, 8)
			} else {
				*out = GetUniverseTypesTypeIdDogmaAttributeList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseTypesTypeIdDogmaAttribute
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
func easyjson460d9473EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetUniverseTypesTypeIdDogmaAttributeList) {
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
func (v GetUniverseTypesTypeIdDogmaAttributeList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson460d9473EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaAttributeList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson460d9473EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaAttributeList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson460d9473DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaAttributeList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson460d9473DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson460d9473DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetUniverseTypesTypeIdDogmaAttribute) {
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
		case "attribute_id":
			out.AttributeId = int32(in.Int32())
		case "value":
			out.Value = float32(in.Float32())
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
func easyjson460d9473EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetUniverseTypesTypeIdDogmaAttribute) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AttributeId != 0 {
		const prefix string = ",\"attribute_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AttributeId))
	}
	if in.Value != 0 {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaAttribute) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson460d9473EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaAttribute) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson460d9473EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaAttribute) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson460d9473DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaAttribute) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson460d9473DecodeGithubComContornoGoesiEsi1(l, v)
}
