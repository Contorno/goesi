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

func easyjsonC59437f0DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseRegionsRegionIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseRegionsRegionIdOkList, 0, 1)
			} else {
				*out = GetUniverseRegionsRegionIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseRegionsRegionIdOk
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
func easyjsonC59437f0EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseRegionsRegionIdOkList) {
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
func (v GetUniverseRegionsRegionIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC59437f0EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseRegionsRegionIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC59437f0EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseRegionsRegionIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC59437f0DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseRegionsRegionIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC59437f0DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonC59437f0DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseRegionsRegionIdOk) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "constellations":
			if in.IsNull() {
				in.Skip()
				out.Constellations = nil
			} else {
				in.Delim('[')
				if out.Constellations == nil {
					if !in.IsDelim(']') {
						out.Constellations = make([]int32, 0, 16)
					} else {
						out.Constellations = []int32{}
					}
				} else {
					out.Constellations = (out.Constellations)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Constellations = append(out.Constellations, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "description":
			out.Description = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "region_id":
			out.RegionId = int32(in.Int32())
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
func easyjsonC59437f0EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseRegionsRegionIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Constellations) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"constellations\":")
		if in.Constellations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Constellations {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	if in.Description != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"description\":")
		out.String(string(in.Description))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.RegionId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"region_id\":")
		out.Int32(int32(in.RegionId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseRegionsRegionIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC59437f0EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseRegionsRegionIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC59437f0EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseRegionsRegionIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC59437f0DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseRegionsRegionIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC59437f0DecodeGithubComAntihaxGoesiEsi1(l, v)
}
