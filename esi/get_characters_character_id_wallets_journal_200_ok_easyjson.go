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

func easyjson5867b91eDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletsJournal200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdWalletsJournal200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdWalletsJournal200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdWalletsJournal200Ok
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
func easyjson5867b91eEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdWalletsJournal200OkList) {
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
func (v GetCharactersCharacterIdWalletsJournal200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5867b91eEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdWalletsJournal200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5867b91eEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletsJournal200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5867b91eDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletsJournal200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5867b91eDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson5867b91eDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdWalletsJournal200Ok) {
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
		case "amount":
			out.Amount = float32(in.Float32())
		case "argument_name":
			out.ArgumentName = string(in.String())
		case "argument_value":
			out.ArgumentValue = int32(in.Int32())
		case "balance":
			out.Balance = float32(in.Float32())
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "first_party_id":
			out.FirstPartyId = int32(in.Int32())
		case "first_party_type":
			out.FirstPartyType = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "ref_id":
			out.RefId = int64(in.Int64())
		case "ref_type_id":
			out.RefTypeId = int32(in.Int32())
		case "second_party_id":
			out.SecondPartyId = int32(in.Int32())
		case "second_party_type":
			out.SecondPartyType = string(in.String())
		case "tax_amount":
			out.TaxAmount = float32(in.Float32())
		case "tax_reciever_id":
			out.TaxRecieverId = int32(in.Int32())
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
func easyjson5867b91eEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdWalletsJournal200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"amount\":")
		out.Float32(float32(in.Amount))
	}
	if in.ArgumentName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"argument_name\":")
		out.String(string(in.ArgumentName))
	}
	if in.ArgumentValue != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"argument_value\":")
		out.Int32(int32(in.ArgumentValue))
	}
	if in.Balance != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"balance\":")
		out.Float32(float32(in.Balance))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"date\":")
		out.Raw((in.Date).MarshalJSON())
	}
	if in.FirstPartyId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"first_party_id\":")
		out.Int32(int32(in.FirstPartyId))
	}
	if in.FirstPartyType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"first_party_type\":")
		out.String(string(in.FirstPartyType))
	}
	if in.Reason != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reason\":")
		out.String(string(in.Reason))
	}
	if in.RefId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ref_id\":")
		out.Int64(int64(in.RefId))
	}
	if in.RefTypeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ref_type_id\":")
		out.Int32(int32(in.RefTypeId))
	}
	if in.SecondPartyId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"second_party_id\":")
		out.Int32(int32(in.SecondPartyId))
	}
	if in.SecondPartyType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"second_party_type\":")
		out.String(string(in.SecondPartyType))
	}
	if in.TaxAmount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"tax_amount\":")
		out.Float32(float32(in.TaxAmount))
	}
	if in.TaxRecieverId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"tax_reciever_id\":")
		out.Int32(int32(in.TaxRecieverId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdWalletsJournal200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5867b91eEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdWalletsJournal200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5867b91eEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletsJournal200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5867b91eDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdWalletsJournal200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5867b91eDecodeGithubComAntihaxGoesiEsi1(l, v)
}
