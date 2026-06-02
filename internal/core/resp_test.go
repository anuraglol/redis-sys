package core

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    interface{}
		wantErr bool
	}{
		{
			name:    "Simple String",
			input:   []byte("+OK\r\n"),
			want:    "OK",
			wantErr: false,
		},
		{
			name:    "Error",
			input:   []byte("-ERR unknown command\r\n"),
			want:    "ERR unknown command",
			wantErr: false,
		},
		{
			name:    "Integer Positive",
			input:   []byte(":1000\r\n"),
			want:    int64(1000),
			wantErr: false,
		},
		{
			name:    "Integer Zero",
			input:   []byte(":0\r\n"),
			want:    int64(0),
			wantErr: false,
		},
		{
			name:    "Bulk String",
			input:   []byte("$5\r\nhello\r\n"),
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "Empty Bulk String",
			input:   []byte("$0\r\n\r\n"),
			want:    "",
			wantErr: false,
		},
		{
			name:    "Empty Array",
			input:   []byte("*0\r\n"),
			want:    []interface{}{},
			wantErr: false,
		},
		{
			name:    "Array of Bulk Strings",
			input:   []byte("*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"),
			want:    []interface{}{"foo", "bar"},
			wantErr: false,
		},
		{
			name:    "Mixed Array",
			input:   []byte("*3\r\n:1\r\n:2\r\n+OK\r\n"),
			want:    []interface{}{int64(1), int64(2), "OK"},
			wantErr: false,
		},
		{
			name:    "Nested Array",
			input:   []byte("*2\r\n*2\r\n:1\r\n:2\r\n*1\r\n+PING\r\n"),
			want:    []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{"PING"}},
			wantErr: false,
		},
		{
			name:    "Empty Input Error",
			input:   []byte{},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
