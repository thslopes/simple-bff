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
		apiCall     ApiCall
		queryString map[string]string
		headers     map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				apiCall: ApiCall{
					Url: "http://hello.com",
				},
			},
			want: "http://hello.com",
		},
		{
			name: "with path params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				apiCall: ApiCall{
					Url: "http://hello.com",
					PathParams: []Param{
						{
							Name:  "name1",
							Value: "value",
							Type:  "querystring",
						},
					},
				},
				queryString: map[string]string{
					"value": "pathValue",
				},
			},
			want: "http://hello.com/name1pathValue",
		},
		{
			name: "with path headers",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				apiCall: ApiCall{
					Url: "http://hello.com",
					Headers: []Param{
						{
							Name:  "name1",
							Value: "value",
							Type:  "querystring",
						},
					},
				},
				queryString: map[string]string{
					"value": "headerValue",
				},
			},
			want: "http://hello.comHname1headerValue",
		},
		{
			name: "with constant qs params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				apiCall: ApiCall{
					Url: "http://hello.com",
					QueryParams: []Param{
						{
							Name:  "name",
							Value: "value",
							Type:  "constant",
						},
					},
				},
			},
			want: "http://hello.comnamevalue",
		},
		{
			name: "with qs to qs params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				apiCall: ApiCall{
					Url: "http://hello.com",
					QueryParams: []Param{
						{
							Name:  "name",
							Value: "qsKey",
							Type:  "querystring",
						},
					},
				},
				queryString: map[string]string{
					"qsKey": "qsValue",
				},
			},
			want: "http://hello.comnameqsValue",
		},
		{
			name: "with header to qs params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				apiCall: ApiCall{
					Url: "http://hello.com",
					QueryParams: []Param{
						{
							Name:  "name",
							Value: "headerKey",
							Type:  "header",
						},
					},
				},
				headers: map[string]string{
					"headerKey": "headerValue",
				},
			},
			want: "http://hello.comnameheaderValue",
		},
		{
			name: "with qs and constant qs params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				apiCall: ApiCall{
					Url: "http://hello.com",
					QueryParams: []Param{
						{
							Name:  "name",
							Value: "value",
							Type:  "constant",
						},
						{
							Name:  "name2",
							Value: "qsKey",
							Type:  "querystring",
						},
					},
				},
				queryString: map[string]string{
					"qsKey": "qsValue",
				},
			},
			want: "http://hello.comnamevaluename2qsValue",
		},
		{
			name: "error",
			fields: fields{
				Getter: &FakeGetter{
					Error: true,
				},
			},
			args: args{
				apiCall: ApiCall{
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
			got, err := c.Get(tt.args.apiCall, tt.args.queryString, tt.args.headers)
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, string(got)); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
		})
	}
}
