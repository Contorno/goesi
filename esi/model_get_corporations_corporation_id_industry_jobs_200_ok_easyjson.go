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

func easyjsonB493deffDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdIndustryJobs200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdIndustryJobs200OkList, 0, 0)
			} else {
				*out = GetCorporationsCorporationIdIndustryJobs200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdIndustryJobs200Ok
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
func easyjsonB493deffEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdIndustryJobs200OkList) {
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
func (v GetCorporationsCorporationIdIndustryJobs200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB493deffEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdIndustryJobs200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB493deffEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdIndustryJobs200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB493deffDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdIndustryJobs200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB493deffDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonB493deffDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdIndustryJobs200Ok) {
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
		case "activity_id":
			out.ActivityId = int32(in.Int32())
		case "blueprint_id":
			out.BlueprintId = int64(in.Int64())
		case "blueprint_location_id":
			out.BlueprintLocationId = int64(in.Int64())
		case "blueprint_type_id":
			out.BlueprintTypeId = int32(in.Int32())
		case "completed_character_id":
			out.CompletedCharacterId = int32(in.Int32())
		case "completed_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CompletedDate).UnmarshalJSON(data))
			}
		case "cost":
			out.Cost = float64(in.Float64())
		case "duration":
			out.Duration = int32(in.Int32())
		case "end_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EndDate).UnmarshalJSON(data))
			}
		case "facility_id":
			out.FacilityId = int64(in.Int64())
		case "installer_id":
			out.InstallerId = int32(in.Int32())
		case "job_id":
			out.JobId = int32(in.Int32())
		case "licensed_runs":
			out.LicensedRuns = int32(in.Int32())
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "output_location_id":
			out.OutputLocationId = int64(in.Int64())
		case "pause_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.PauseDate).UnmarshalJSON(data))
			}
		case "probability":
			out.Probability = float32(in.Float32())
		case "product_type_id":
			out.ProductTypeId = int32(in.Int32())
		case "runs":
			out.Runs = int32(in.Int32())
		case "start_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartDate).UnmarshalJSON(data))
			}
		case "status":
			out.Status = string(in.String())
		case "successful_runs":
			out.SuccessfulRuns = int32(in.Int32())
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
func easyjsonB493deffEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdIndustryJobs200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ActivityId != 0 {
		const prefix string = ",\"activity_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.ActivityId))
	}
	if in.BlueprintId != 0 {
		const prefix string = ",\"blueprint_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.BlueprintId))
	}
	if in.BlueprintLocationId != 0 {
		const prefix string = ",\"blueprint_location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.BlueprintLocationId))
	}
	if in.BlueprintTypeId != 0 {
		const prefix string = ",\"blueprint_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.BlueprintTypeId))
	}
	if in.CompletedCharacterId != 0 {
		const prefix string = ",\"completed_character_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CompletedCharacterId))
	}
	if true {
		const prefix string = ",\"completed_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CompletedDate).MarshalJSON())
	}
	if in.Cost != 0 {
		const prefix string = ",\"cost\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Cost))
	}
	if in.Duration != 0 {
		const prefix string = ",\"duration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Duration))
	}
	if true {
		const prefix string = ",\"end_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.EndDate).MarshalJSON())
	}
	if in.FacilityId != 0 {
		const prefix string = ",\"facility_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.FacilityId))
	}
	if in.InstallerId != 0 {
		const prefix string = ",\"installer_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.InstallerId))
	}
	if in.JobId != 0 {
		const prefix string = ",\"job_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.JobId))
	}
	if in.LicensedRuns != 0 {
		const prefix string = ",\"licensed_runs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.LicensedRuns))
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
	if in.OutputLocationId != 0 {
		const prefix string = ",\"output_location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OutputLocationId))
	}
	if true {
		const prefix string = ",\"pause_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.PauseDate).MarshalJSON())
	}
	if in.Probability != 0 {
		const prefix string = ",\"probability\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Probability))
	}
	if in.ProductTypeId != 0 {
		const prefix string = ",\"product_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ProductTypeId))
	}
	if in.Runs != 0 {
		const prefix string = ",\"runs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Runs))
	}
	if true {
		const prefix string = ",\"start_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.StartDate).MarshalJSON())
	}
	if in.Status != "" {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	if in.SuccessfulRuns != 0 {
		const prefix string = ",\"successful_runs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SuccessfulRuns))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdIndustryJobs200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB493deffEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdIndustryJobs200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB493deffEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdIndustryJobs200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB493deffDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdIndustryJobs200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB493deffDecodeGithubComContornoGoesiEsi1(l, v)
}
