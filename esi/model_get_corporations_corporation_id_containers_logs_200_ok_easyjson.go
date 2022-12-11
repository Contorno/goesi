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

func easyjsonD2e11036DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdContainersLogs200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdContainersLogs200OkList, 0, 0)
			} else {
				*out = GetCorporationsCorporationIdContainersLogs200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdContainersLogs200Ok
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
func easyjsonD2e11036EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdContainersLogs200OkList) {
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
func (v GetCorporationsCorporationIdContainersLogs200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2e11036EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdContainersLogs200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2e11036EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdContainersLogs200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2e11036DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdContainersLogs200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2e11036DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonD2e11036DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdContainersLogs200Ok) {
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
		case "action":
			out.Action = string(in.String())
		case "character_id":
			out.CharacterId = int32(in.Int32())
		case "container_id":
			out.ContainerId = int64(in.Int64())
		case "container_type_id":
			out.ContainerTypeId = int32(in.Int32())
		case "location_flag":
			out.LocationFlag = string(in.String())
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "logged_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LoggedAt).UnmarshalJSON(data))
			}
		case "new_config_bitmask":
			out.NewConfigBitmask = int32(in.Int32())
		case "old_config_bitmask":
			out.OldConfigBitmask = int32(in.Int32())
		case "password_type":
			out.PasswordType = string(in.String())
		case "quantity":
			out.Quantity = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjsonD2e11036EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdContainersLogs200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Action != "" {
		const prefix string = ",\"action\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Action))
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
	if in.ContainerId != 0 {
		const prefix string = ",\"container_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ContainerId))
	}
	if in.ContainerTypeId != 0 {
		const prefix string = ",\"container_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ContainerTypeId))
	}
	if in.LocationFlag != "" {
		const prefix string = ",\"location_flag\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LocationFlag))
	}
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LocationId))
	}
	if true {
		const prefix string = ",\"logged_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LoggedAt).MarshalJSON())
	}
	if in.NewConfigBitmask != 0 {
		const prefix string = ",\"new_config_bitmask\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.NewConfigBitmask))
	}
	if in.OldConfigBitmask != 0 {
		const prefix string = ",\"old_config_bitmask\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.OldConfigBitmask))
	}
	if in.PasswordType != "" {
		const prefix string = ",\"password_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PasswordType))
	}
	if in.Quantity != 0 {
		const prefix string = ",\"quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Quantity))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdContainersLogs200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2e11036EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdContainersLogs200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2e11036EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdContainersLogs200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2e11036DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdContainersLogs200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2e11036DecodeGithubComContornoGoesiEsi1(l, v)
}
