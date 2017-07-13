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

func easyjson3cad633cDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdSkillsSkillList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdSkillsSkillList, 0, 4)
			} else {
				*out = GetCharactersCharacterIdSkillsSkillList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdSkillsSkill
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
func easyjson3cad633cEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdSkillsSkillList) {
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
func (v GetCharactersCharacterIdSkillsSkillList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3cad633cEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSkillsSkillList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3cad633cEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillsSkillList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3cad633cDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillsSkillList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3cad633cDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson3cad633cDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdSkillsSkill) {
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
		case "current_skill_level":
			out.CurrentSkillLevel = int32(in.Int32())
		case "skill_id":
			out.SkillId = int32(in.Int32())
		case "skillpoints_in_skill":
			out.SkillpointsInSkill = int64(in.Int64())
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
func easyjson3cad633cEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdSkillsSkill) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CurrentSkillLevel != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"current_skill_level\":")
		out.Int32(int32(in.CurrentSkillLevel))
	}
	if in.SkillId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"skill_id\":")
		out.Int32(int32(in.SkillId))
	}
	if in.SkillpointsInSkill != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"skillpoints_in_skill\":")
		out.Int64(int64(in.SkillpointsInSkill))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdSkillsSkill) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3cad633cEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdSkillsSkill) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3cad633cEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillsSkill) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3cad633cDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdSkillsSkill) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3cad633cDecodeGithubComAntihaxGoesiEsi1(l, v)
}
