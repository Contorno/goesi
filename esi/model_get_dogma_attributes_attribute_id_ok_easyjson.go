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

func easyjson6176f835DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetDogmaAttributesAttributeIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetDogmaAttributesAttributeIdOkList, 0, 0)
			} else {
				*out = GetDogmaAttributesAttributeIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetDogmaAttributesAttributeIdOk
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
func easyjson6176f835EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetDogmaAttributesAttributeIdOkList) {
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
func (v GetDogmaAttributesAttributeIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6176f835EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetDogmaAttributesAttributeIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6176f835EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetDogmaAttributesAttributeIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6176f835DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetDogmaAttributesAttributeIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6176f835DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson6176f835DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetDogmaAttributesAttributeIdOk) {
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
		case "default_value":
			out.DefaultValue = float32(in.Float32())
		case "description":
			out.Description = string(in.String())
		case "display_name":
			out.DisplayName = string(in.String())
		case "high_is_good":
			out.HighIsGood = bool(in.Bool())
		case "icon_id":
			out.IconId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "published":
			out.Published = bool(in.Bool())
		case "stackable":
			out.Stackable = bool(in.Bool())
		case "unit_id":
			out.UnitId = int32(in.Int32())
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
func easyjson6176f835EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetDogmaAttributesAttributeIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AttributeId != 0 {
		const prefix string = ",\"attribute_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AttributeId))
	}
	if in.DefaultValue != 0 {
		const prefix string = ",\"default_value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.DefaultValue))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.DisplayName != "" {
		const prefix string = ",\"display_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DisplayName))
	}
	if in.HighIsGood {
		const prefix string = ",\"high_is_good\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HighIsGood))
	}
	if in.IconId != 0 {
		const prefix string = ",\"icon_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.IconId))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Published {
		const prefix string = ",\"published\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Published))
	}
	if in.Stackable {
		const prefix string = ",\"stackable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Stackable))
	}
	if in.UnitId != 0 {
		const prefix string = ",\"unit_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.UnitId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetDogmaAttributesAttributeIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6176f835EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetDogmaAttributesAttributeIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6176f835EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetDogmaAttributesAttributeIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6176f835DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetDogmaAttributesAttributeIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6176f835DecodeGithubComContornoGoesiEsi1(l, v)
}
