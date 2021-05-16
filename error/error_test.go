package error

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNotionError_Error(t *testing.T) {
	type fields struct {
		Status  int
		Code    Code
		Message string
		Object  string
	}
	tt := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "JointAaaBbb",
			fields: fields{
				Status:  400,
				Code:    "Aaa",
				Message: "Bbb",
				Object:  "error",
			},
			want: "Aaa: Bbb",
		},
	}
	for _, tt := range tt {
		t.Run(tt.name, func(t *testing.T) {
			ne := &NotionError{
				Status:  tt.fields.Status,
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Object:  tt.fields.Object,
			}

			got := ne.Error()

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("unexpected return value\n%v", diff)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		response string
	}
	tt := []struct {
		name    string
		args    args
		want    *NotionError
		wantErr bool
	}{
		{
			name: "validation error",
			args: args{
				response: `
{
    "code": "validation_error",
    "message": "The provided database ID is not a valid Notion UUID: aaa.",
    "object": "error",
    "status": 400
}
`,
			},
			want: &NotionError{
				Status:  400,
				Code:    CodeValidationError,
				Message: "The provided database ID is not a valid Notion UUID: aaa.",
				Object:  "error",
			},
			wantErr: false,
		},
		{
			name: "undefined code",
			args: args{
				response: `
{
    "code": "undefined code",
    "message": "undefined error message",
    "object": "error",
    "status": 400
}
`,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tt {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse([]byte(tt.args.response))

			if err != nil && !tt.wantErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tt.wantErr {
				t.Error("error expected")
			} else if !tt.wantErr { // first return value should not be used when the error is returned
				if diff := cmp.Diff(got, tt.want); diff != "" {
					t.Errorf("unexpected return value\n%v", diff)
				}
			}
		})
	}
}
