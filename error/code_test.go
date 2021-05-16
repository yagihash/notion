package error

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCode_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tt := []struct {
		name    string
		args    args
		want    Code
		wantErr bool
	}{
		{
			name: `"invalid_json"`,
			args: args{
				data: []byte(`"invalid_json"`),
			},
			want:    CodeInvalidJson,
			wantErr: false,
		},
		{
			name: `"undefined_code"`,
			args: args{
				data: []byte(`"undefined_code"`),
			},
			want:    Code(""),
			wantErr: true,
		},
		{
			name: `unsupported type bool`,
			args: args{
				data: []byte("true"),
			},
			want:    Code(""),
			wantErr: true,
		},
		{
			name: `unsupported type int`,
			args: args{
				data: []byte("100"),
			},
			want:    Code(""),
			wantErr: true,
		},
	}

	for _, c := range tt {
		t.Run(c.name, func(t *testing.T) {
			code := Code("")
			err := code.UnmarshalJSON(c.args.data)

			if err != nil && !c.wantErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && c.wantErr {
				t.Error("error expected")
			} else {
				if diff := cmp.Diff(code, c.want); diff != "" {
					t.Errorf("unexpected return value\n%v", diff)
				}
			}
		})
	}
}
