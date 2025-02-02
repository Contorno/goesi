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

func easyjson6cb8c3e4DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetUniverseTypesTypeIdDogmaEffectList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseTypesTypeIdDogmaEffectList, 0, 8)
			} else {
				*out = GetUniverseTypesTypeIdDogmaEffectList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseTypesTypeIdDogmaEffect
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
func easyjson6cb8c3e4EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetUniverseTypesTypeIdDogmaEffectList) {
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
func (v GetUniverseTypesTypeIdDogmaEffectList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6cb8c3e4EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaEffectList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6cb8c3e4EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaEffectList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6cb8c3e4DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaEffectList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6cb8c3e4DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson6cb8c3e4DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetUniverseTypesTypeIdDogmaEffect) {
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
		case "effect_id":
			out.EffectId = int32(in.Int32())
		case "is_default":
			out.IsDefault = bool(in.Bool())
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
func easyjson6cb8c3e4EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetUniverseTypesTypeIdDogmaEffect) {
	out.RawByte('{')
	first := true
	_ = first
	if in.EffectId != 0 {
		const prefix string = ",\"effect_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.EffectId))
	}
	if in.IsDefault {
		const prefix string = ",\"is_default\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsDefault))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaEffect) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6cb8c3e4EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaEffect) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6cb8c3e4EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaEffect) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6cb8c3e4DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaEffect) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6cb8c3e4DecodeGithubComContornoGoesiEsi1(l, v)
}
