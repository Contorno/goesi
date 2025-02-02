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

func easyjsonFa740090DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersKillsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsCharactersKillsList, 0, 0)
			} else {
				*out = GetFwLeaderboardsCharactersKillsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsCharactersKills
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
func easyjsonFa740090EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsCharactersKillsList) {
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
func (v GetFwLeaderboardsCharactersKillsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa740090EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCharactersKillsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa740090EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersKillsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa740090DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersKillsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa740090DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonFa740090DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersKills) {
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
		case "active_total":
			if in.IsNull() {
				in.Skip()
				out.ActiveTotal = nil
			} else {
				in.Delim('[')
				if out.ActiveTotal == nil {
					if !in.IsDelim(']') {
						out.ActiveTotal = make([]GetFwLeaderboardsCharactersActiveTotalActiveTotal, 0, 8)
					} else {
						out.ActiveTotal = []GetFwLeaderboardsCharactersActiveTotalActiveTotal{}
					}
				} else {
					out.ActiveTotal = (out.ActiveTotal)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetFwLeaderboardsCharactersActiveTotalActiveTotal
					(v4).UnmarshalEasyJSON(in)
					out.ActiveTotal = append(out.ActiveTotal, v4)
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
						out.LastWeek = make([]GetFwLeaderboardsCharactersLastWeekLastWeek, 0, 8)
					} else {
						out.LastWeek = []GetFwLeaderboardsCharactersLastWeekLastWeek{}
					}
				} else {
					out.LastWeek = (out.LastWeek)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetFwLeaderboardsCharactersLastWeekLastWeek
					(v5).UnmarshalEasyJSON(in)
					out.LastWeek = append(out.LastWeek, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "yesterday":
			if in.IsNull() {
				in.Skip()
				out.Yesterday = nil
			} else {
				in.Delim('[')
				if out.Yesterday == nil {
					if !in.IsDelim(']') {
						out.Yesterday = make([]GetFwLeaderboardsCharactersYesterdayYesterday, 0, 8)
					} else {
						out.Yesterday = []GetFwLeaderboardsCharactersYesterdayYesterday{}
					}
				} else {
					out.Yesterday = (out.Yesterday)[:0]
				}
				for !in.IsDelim(']') {
					var v6 GetFwLeaderboardsCharactersYesterdayYesterday
					(v6).UnmarshalEasyJSON(in)
					out.Yesterday = append(out.Yesterday, v6)
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
func easyjsonFa740090EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsCharactersKills) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.ActiveTotal) != 0 {
		const prefix string = ",\"active_total\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v7, v8 := range in.ActiveTotal {
				if v7 > 0 {
					out.RawByte(',')
				}
				(v8).MarshalEasyJSON(out)
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
			for v11, v12 := range in.Yesterday {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCharactersKills) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa740090EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCharactersKills) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa740090EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersKills) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa740090DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersKills) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa740090DecodeGithubComContornoGoesiEsi1(l, v)
}
