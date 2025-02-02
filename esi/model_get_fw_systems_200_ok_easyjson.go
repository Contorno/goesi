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

func easyjsonB042087bDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetFwSystems200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwSystems200OkList, 0, 1)
			} else {
				*out = GetFwSystems200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwSystems200Ok
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
func easyjsonB042087bEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetFwSystems200OkList) {
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
func (v GetFwSystems200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB042087bEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwSystems200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB042087bEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwSystems200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB042087bDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwSystems200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB042087bDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonB042087bDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetFwSystems200Ok) {
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
		case "contested":
			out.Contested = string(in.String())
		case "occupier_faction_id":
			out.OccupierFactionId = int32(in.Int32())
		case "owner_faction_id":
			out.OwnerFactionId = int32(in.Int32())
		case "solar_system_id":
			out.SolarSystemId = int32(in.Int32())
		case "victory_points":
			out.VictoryPoints = int32(in.Int32())
		case "victory_points_threshold":
			out.VictoryPointsThreshold = int32(in.Int32())
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
func easyjsonB042087bEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetFwSystems200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Contested != "" {
		const prefix string = ",\"contested\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Contested))
	}
	if in.OccupierFactionId != 0 {
		const prefix string = ",\"occupier_faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.OccupierFactionId))
	}
	if in.OwnerFactionId != 0 {
		const prefix string = ",\"owner_faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.OwnerFactionId))
	}
	if in.SolarSystemId != 0 {
		const prefix string = ",\"solar_system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SolarSystemId))
	}
	if in.VictoryPoints != 0 {
		const prefix string = ",\"victory_points\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VictoryPoints))
	}
	if in.VictoryPointsThreshold != 0 {
		const prefix string = ",\"victory_points_threshold\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VictoryPointsThreshold))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwSystems200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB042087bEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwSystems200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB042087bEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwSystems200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB042087bDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwSystems200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB042087bDecodeGithubComContornoGoesiEsi1(l, v)
}
