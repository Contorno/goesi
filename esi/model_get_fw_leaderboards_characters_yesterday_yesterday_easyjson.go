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

func easyjson7e1bb9aeDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersYesterdayYesterdayList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsCharactersYesterdayYesterdayList, 0, 8)
			} else {
				*out = GetFwLeaderboardsCharactersYesterdayYesterdayList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsCharactersYesterdayYesterday
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
func easyjson7e1bb9aeEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsCharactersYesterdayYesterdayList) {
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
func (v GetFwLeaderboardsCharactersYesterdayYesterdayList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7e1bb9aeEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCharactersYesterdayYesterdayList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7e1bb9aeEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersYesterdayYesterdayList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7e1bb9aeDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersYesterdayYesterdayList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7e1bb9aeDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson7e1bb9aeDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsCharactersYesterdayYesterday) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
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
func easyjson7e1bb9aeEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsCharactersYesterdayYesterday) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Amount))
	}
	if in.CharacterId != 0 {
		const prefix string = ",\"character_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CharacterId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCharactersYesterdayYesterday) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7e1bb9aeEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCharactersYesterdayYesterday) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7e1bb9aeEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersYesterdayYesterday) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7e1bb9aeDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCharactersYesterdayYesterday) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7e1bb9aeDecodeGithubComContornoGoesiEsi1(l, v)
}
