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

func easyjson2d81da92DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetUniverseTypesTypeIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseTypesTypeIdOkList, 0, 0)
			} else {
				*out = GetUniverseTypesTypeIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseTypesTypeIdOk
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
func easyjson2d81da92EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetUniverseTypesTypeIdOkList) {
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
func (v GetUniverseTypesTypeIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2d81da92EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2d81da92EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2d81da92DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2d81da92DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson2d81da92DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetUniverseTypesTypeIdOk) {
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
		case "capacity":
			out.Capacity = float32(in.Float32())
		case "description":
			out.Description = string(in.String())
		case "dogma_attributes":
			if in.IsNull() {
				in.Skip()
				out.DogmaAttributes = nil
			} else {
				in.Delim('[')
				if out.DogmaAttributes == nil {
					if !in.IsDelim(']') {
						out.DogmaAttributes = make([]GetUniverseTypesTypeIdDogmaAttribute, 0, 8)
					} else {
						out.DogmaAttributes = []GetUniverseTypesTypeIdDogmaAttribute{}
					}
				} else {
					out.DogmaAttributes = (out.DogmaAttributes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetUniverseTypesTypeIdDogmaAttribute
					easyjson2d81da92DecodeGithubComContornoGoesiEsi2(in, &v4)
					out.DogmaAttributes = append(out.DogmaAttributes, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "dogma_effects":
			if in.IsNull() {
				in.Skip()
				out.DogmaEffects = nil
			} else {
				in.Delim('[')
				if out.DogmaEffects == nil {
					if !in.IsDelim(']') {
						out.DogmaEffects = make([]GetUniverseTypesTypeIdDogmaEffect, 0, 8)
					} else {
						out.DogmaEffects = []GetUniverseTypesTypeIdDogmaEffect{}
					}
				} else {
					out.DogmaEffects = (out.DogmaEffects)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetUniverseTypesTypeIdDogmaEffect
					easyjson2d81da92DecodeGithubComContornoGoesiEsi3(in, &v5)
					out.DogmaEffects = append(out.DogmaEffects, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "graphic_id":
			out.GraphicId = int32(in.Int32())
		case "group_id":
			out.GroupId = int32(in.Int32())
		case "icon_id":
			out.IconId = int32(in.Int32())
		case "market_group_id":
			out.MarketGroupId = int32(in.Int32())
		case "mass":
			out.Mass = float32(in.Float32())
		case "name":
			out.Name = string(in.String())
		case "packaged_volume":
			out.PackagedVolume = float32(in.Float32())
		case "portion_size":
			out.PortionSize = int32(in.Int32())
		case "published":
			out.Published = bool(in.Bool())
		case "radius":
			out.Radius = float32(in.Float32())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "volume":
			out.Volume = float32(in.Float32())
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
func easyjson2d81da92EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetUniverseTypesTypeIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Capacity != 0 {
		const prefix string = ",\"capacity\":"
		first = false
		out.RawString(prefix[1:])
		out.Float32(float32(in.Capacity))
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
	if len(in.DogmaAttributes) != 0 {
		const prefix string = ",\"dogma_attributes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.DogmaAttributes {
				if v6 > 0 {
					out.RawByte(',')
				}
				easyjson2d81da92EncodeGithubComContornoGoesiEsi2(out, v7)
			}
			out.RawByte(']')
		}
	}
	if len(in.DogmaEffects) != 0 {
		const prefix string = ",\"dogma_effects\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.DogmaEffects {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjson2d81da92EncodeGithubComContornoGoesiEsi3(out, v9)
			}
			out.RawByte(']')
		}
	}
	if in.GraphicId != 0 {
		const prefix string = ",\"graphic_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.GraphicId))
	}
	if in.GroupId != 0 {
		const prefix string = ",\"group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.GroupId))
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
	if in.MarketGroupId != 0 {
		const prefix string = ",\"market_group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MarketGroupId))
	}
	if in.Mass != 0 {
		const prefix string = ",\"mass\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Mass))
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
	if in.PackagedVolume != 0 {
		const prefix string = ",\"packaged_volume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.PackagedVolume))
	}
	if in.PortionSize != 0 {
		const prefix string = ",\"portion_size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PortionSize))
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
	if in.Radius != 0 {
		const prefix string = ",\"radius\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Radius))
	}
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	if in.Volume != 0 {
		const prefix string = ",\"volume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Volume))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseTypesTypeIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2d81da92EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2d81da92EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2d81da92DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2d81da92DecodeGithubComContornoGoesiEsi1(l, v)
}
func easyjson2d81da92DecodeGithubComContornoGoesiEsi3(in *jlexer.Lexer, out *GetUniverseTypesTypeIdDogmaEffect) {
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
func easyjson2d81da92EncodeGithubComContornoGoesiEsi3(out *jwriter.Writer, in GetUniverseTypesTypeIdDogmaEffect) {
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
func easyjson2d81da92DecodeGithubComContornoGoesiEsi2(in *jlexer.Lexer, out *GetUniverseTypesTypeIdDogmaAttribute) {
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
func easyjson2d81da92EncodeGithubComContornoGoesiEsi2(out *jwriter.Writer, in GetUniverseTypesTypeIdDogmaAttribute) {
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
