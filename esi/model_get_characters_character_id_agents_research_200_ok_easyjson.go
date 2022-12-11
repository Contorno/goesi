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

func easyjsonB249e244DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdAgentsResearch200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdAgentsResearch200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdAgentsResearch200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdAgentsResearch200Ok
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
func easyjsonB249e244EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdAgentsResearch200OkList) {
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
func (v GetCharactersCharacterIdAgentsResearch200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB249e244EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdAgentsResearch200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB249e244EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdAgentsResearch200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB249e244DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdAgentsResearch200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB249e244DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonB249e244DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdAgentsResearch200Ok) {
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
		case "agent_id":
			out.AgentId = int32(in.Int32())
		case "points_per_day":
			out.PointsPerDay = float32(in.Float32())
		case "remainder_points":
			out.RemainderPoints = float32(in.Float32())
		case "skill_type_id":
			out.SkillTypeId = int32(in.Int32())
		case "started_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartedAt).UnmarshalJSON(data))
			}
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
func easyjsonB249e244EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdAgentsResearch200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AgentId != 0 {
		const prefix string = ",\"agent_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AgentId))
	}
	if in.PointsPerDay != 0 {
		const prefix string = ",\"points_per_day\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.PointsPerDay))
	}
	if in.RemainderPoints != 0 {
		const prefix string = ",\"remainder_points\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.RemainderPoints))
	}
	if in.SkillTypeId != 0 {
		const prefix string = ",\"skill_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SkillTypeId))
	}
	if true {
		const prefix string = ",\"started_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.StartedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdAgentsResearch200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB249e244EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdAgentsResearch200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB249e244EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdAgentsResearch200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB249e244DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdAgentsResearch200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB249e244DecodeGithubComContornoGoesiEsi1(l, v)
}
