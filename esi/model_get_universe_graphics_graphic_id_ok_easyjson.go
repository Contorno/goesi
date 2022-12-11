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

func easyjsonF7263cfaDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetUniverseGraphicsGraphicIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseGraphicsGraphicIdOkList, 0, 0)
			} else {
				*out = GetUniverseGraphicsGraphicIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseGraphicsGraphicIdOk
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
func easyjsonF7263cfaEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetUniverseGraphicsGraphicIdOkList) {
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
func (v GetUniverseGraphicsGraphicIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF7263cfaEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseGraphicsGraphicIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF7263cfaEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseGraphicsGraphicIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF7263cfaDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseGraphicsGraphicIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF7263cfaDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonF7263cfaDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetUniverseGraphicsGraphicIdOk) {
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
		case "collision_file":
			out.CollisionFile = string(in.String())
		case "graphic_file":
			out.GraphicFile = string(in.String())
		case "graphic_id":
			out.GraphicId = int32(in.Int32())
		case "icon_folder":
			out.IconFolder = string(in.String())
		case "sof_dna":
			out.SofDna = string(in.String())
		case "sof_fation_name":
			out.SofFationName = string(in.String())
		case "sof_hull_name":
			out.SofHullName = string(in.String())
		case "sof_race_name":
			out.SofRaceName = string(in.String())
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
func easyjsonF7263cfaEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetUniverseGraphicsGraphicIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CollisionFile != "" {
		const prefix string = ",\"collision_file\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.CollisionFile))
	}
	if in.GraphicFile != "" {
		const prefix string = ",\"graphic_file\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GraphicFile))
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
	if in.IconFolder != "" {
		const prefix string = ",\"icon_folder\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IconFolder))
	}
	if in.SofDna != "" {
		const prefix string = ",\"sof_dna\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SofDna))
	}
	if in.SofFationName != "" {
		const prefix string = ",\"sof_fation_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SofFationName))
	}
	if in.SofHullName != "" {
		const prefix string = ",\"sof_hull_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SofHullName))
	}
	if in.SofRaceName != "" {
		const prefix string = ",\"sof_race_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SofRaceName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseGraphicsGraphicIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF7263cfaEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseGraphicsGraphicIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF7263cfaEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseGraphicsGraphicIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF7263cfaDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseGraphicsGraphicIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF7263cfaDecodeGithubComContornoGoesiEsi1(l, v)
}
