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

func easyjson682940b2DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsVictoryPointsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsVictoryPointsList, 0, 1)
			} else {
				*out = GetFwLeaderboardsVictoryPointsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsVictoryPoints
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
func easyjson682940b2EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsVictoryPointsList) {
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
func (v GetFwLeaderboardsVictoryPointsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson682940b2EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsVictoryPointsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson682940b2EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsVictoryPointsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson682940b2DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsVictoryPointsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson682940b2DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson682940b2DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsVictoryPoints) {
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
		case "yesterday":
			if in.IsNull() {
				in.Skip()
				out.Yesterday = nil
			} else {
				in.Delim('[')
				if out.Yesterday == nil {
					if !in.IsDelim(']') {
						out.Yesterday = make([]GetFwLeaderboardsYesterday1, 0, 8)
					} else {
						out.Yesterday = []GetFwLeaderboardsYesterday1{}
					}
				} else {
					out.Yesterday = (out.Yesterday)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetFwLeaderboardsYesterday1
					easyjson682940b2DecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.Yesterday = append(out.Yesterday, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "last_week":
			if in.IsNull() {
				in.Skip()
				out.LastWeek = nil
			} else {
				in.Delim('[')
				if out.LastWeek == nil {
					if !in.IsDelim(']') {
						out.LastWeek = make([]GetFwLeaderboardsLastWeek1, 0, 8)
					} else {
						out.LastWeek = []GetFwLeaderboardsLastWeek1{}
					}
				} else {
					out.LastWeek = (out.LastWeek)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetFwLeaderboardsLastWeek1
					(v5).UnmarshalEasyJSON(in)
					out.LastWeek = append(out.LastWeek, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "active_total":
			if in.IsNull() {
				in.Skip()
				out.ActiveTotal = nil
			} else {
				in.Delim('[')
				if out.ActiveTotal == nil {
					if !in.IsDelim(']') {
						out.ActiveTotal = make([]GetFwLeaderboardsActiveTotal1, 0, 8)
					} else {
						out.ActiveTotal = []GetFwLeaderboardsActiveTotal1{}
					}
				} else {
					out.ActiveTotal = (out.ActiveTotal)[:0]
				}
				for !in.IsDelim(']') {
					var v6 GetFwLeaderboardsActiveTotal1
					easyjson682940b2DecodeGithubComAntihaxGoesiEsi3(in, &v6)
					out.ActiveTotal = append(out.ActiveTotal, v6)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson682940b2EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsVictoryPoints) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Yesterday) != 0 {
		const prefix string = ",\"yesterday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.Yesterday {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjson682940b2EncodeGithubComAntihaxGoesiEsi2(out, v8)
			}
			out.RawByte(']')
		}
	}
	if len(in.LastWeek) != 0 {
		const prefix string = ",\"last_week\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.LastWeek {
				if v9 > 0 {
					out.RawByte(',')
				}
				(v10).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.ActiveTotal) != 0 {
		const prefix string = ",\"active_total\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.ActiveTotal {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjson682940b2EncodeGithubComAntihaxGoesiEsi3(out, v12)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsVictoryPoints) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson682940b2EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsVictoryPoints) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson682940b2EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsVictoryPoints) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson682940b2DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsVictoryPoints) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson682940b2DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson682940b2DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetFwLeaderboardsActiveTotal1) {
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
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "amount":
			out.Amount = int32(in.Int32())
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
func easyjson682940b2EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetFwLeaderboardsActiveTotal1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Amount))
	}
	out.RawByte('}')
}
func easyjson682940b2DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetFwLeaderboardsYesterday1) {
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
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "amount":
			out.Amount = int32(in.Int32())
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
func easyjson682940b2EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetFwLeaderboardsYesterday1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Amount))
	}
	out.RawByte('}')
}
