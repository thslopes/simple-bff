package apicall

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCaller_Get(t *testing.T) {
	Resources = map[string]string{
		"hello": "http://hello.com",
		"wrongurl": "wrong.url",
	}
	Queries = map[string]Query{
		"success": {
			Resource: "hello",
		},
		"withPathParams": {
			Resource: "hello",
			PathParams: []Param{
				{
					Name:  "name1",
					Value: "value",
					Type:  "querystring",
				},
			},
		},
		"withHeaders": {
			Resource: "hello",
			Headers: []Param{
				{
					Name:  "name1",
					Value: "value",
					Type:  "querystring",
				},
			},
		},
		"withConstantQs": {
			Resource: "hello",
			QueryParams: []Param{
				{
					Name:  "name",
					Value: "value",
					Type:  "constant",
				},
			},
		},
		"withQsToQs": {
			Resource: "hello",
			QueryParams: []Param{
				{
					Name:  "name",
					Value: "qsKey",
					Type:  "querystring",
				},
			},
		},
		"withHeaderToQs": {
			Resource: "hello",
			QueryParams: []Param{
				{
					Name:  "name",
					Value: "headerKey",
					Type:  "header",
				},
			},
		},
		"withQsAndConstantQs": {
			Resource: "hello",
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
		"error": {
			Resource: "wrongurl",
		},
	}
	type fields struct {
		Getter *FakeGetter
	}
	type args struct {
		query       string
		queryString map[string]string
		headers     map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FakeGetter
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				query: "success",
			},
			want: &FakeGetter{
				Url:        "http://hello.com",
				Qs:         map[string]string{},
				PathParams: map[string]string{},
				Headers:    map[string]string{},
			},
		},
		{
			name: "with path params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				query: "withPathParams",
				queryString: map[string]string{
					"value": "pathValue",
				},
			},
			want: &FakeGetter{
				Url:        "http://hello.com",
				Qs:         map[string]string{},
				PathParams: map[string]string{"name1": "pathValue"},
				Headers:    map[string]string{},
			},
		},
		{
			name: "with path headers",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				query: "withHeaders",
				queryString: map[string]string{
					"value": "headerValue",
				},
			},
			want: &FakeGetter{
				Url:        "http://hello.com",
				Qs:         map[string]string{},
				PathParams: map[string]string{},
				Headers:    map[string]string{"name1": "headerValue"},
			},
		},
		{
			name: "with constant qs params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				query: "withConstantQs",
			},
			want: &FakeGetter{
				Url:        "http://hello.com",
				Qs:         map[string]string{"name": "value"},
				PathParams: map[string]string{},
				Headers:    map[string]string{},
			},
		},
		{
			name: "with qs to qs params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				query: "withQsToQs",
				queryString: map[string]string{
					"qsKey": "qsValue",
				},
			},
			want: &FakeGetter{
				Url:        "http://hello.com",
				Qs:         map[string]string{"name": "qsValue"},
				PathParams: map[string]string{},
				Headers:    map[string]string{},
			},
		},
		{
			name: "with header to qs params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				query: "withHeaderToQs",
				headers: map[string]string{
					"headerKey": "headerValue",
				},
			},
			want: &FakeGetter{
				Url:        "http://hello.com",
				Qs:         map[string]string{"name": "headerValue"},
				PathParams: map[string]string{},
				Headers:    map[string]string{},
			},
		},
		{
			name: "with qs and constant qs params",
			fields: fields{
				Getter: &FakeGetter{},
			},
			args: args{
				query: "withQsAndConstantQs",
				queryString: map[string]string{
					"qsKey": "qsValue",
				},
			},
			want: &FakeGetter{
				Url:        "http://hello.com",
				Qs:         map[string]string{"name": "value", "name2": "qsValue"},
				PathParams: map[string]string{},
				Headers:    map[string]string{},
			},
		},
		{
			name: "error",
			fields: fields{
				Getter: &FakeGetter{
					Error: true,
				},
			},
			args: args{
				query: "error",
			},
			wantErr: &GetterErr{
				Err: "wrong.url",
			},
			want: &FakeGetter{
				Error: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Caller{
				Getter: tt.fields.Getter,
			}
			_, err := c.Get(tt.args.query, tt.args.queryString, tt.args.headers)
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, tt.fields.Getter); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
		})
	}
}
