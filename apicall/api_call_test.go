package apicall

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestCaller_Get(t *testing.T) {
	type fields struct {
		Getter Getter
	}
	type args struct {
		apiCall *ApiCall
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				apiCall: &ApiCall{
					Url: "http://hello.com",
				},
			},
			want:    []byte("http://hello.com"),
		},
		{
			name: "error",
			fields: fields{
				Getter: &FakeGetter{
					Error: true,
				},
			},
			args: args{
				apiCall: &ApiCall{
					Url: "wrong.url",
				},
			},
			wantErr: &GetterErr{
				Err: "wrong.url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Caller{
				Getter: tt.fields.Getter,
			}
			got, err := c.Get(tt.args.apiCall)
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
		})
	}
}
