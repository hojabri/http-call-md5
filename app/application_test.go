package app

import (
	"encoding/hex"
	"testing"
)

func TestMD5(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				inputText: "test1",
			},
			want: "5a105e8b9d40e1329780d62ea2265d8a",
		},
		{
			name: "test2 (blank)",
			args: args{
				inputText: "",
			},
			want: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name: "test3(adjust string)",
			args: args{
				inputText: "Adjust is the industry leader in mobile measurement and fraud prevention. By making marketing simpler, smarter and more secure, we empower data-driven marketers to succeed.",
			},
			want: "e53646d67f608e0bac5c47bd2c47dc6c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got :=hex.EncodeToString(MD5(tt.args.inputText)); got!= tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_checkUrlFormat(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1 - google.com",
			args:    args{
				address: "google.com",
			},
			wantErr: false,
		},
		{
			name:    "test2 - http://google.com",
			args:    args{
				address: "http://google.com",
			},
			wantErr: false,
		},
		{
			name:    "test3 - http://google.com%",
			args:    args{
				address: "http://google.com%",
			},
			wantErr: true,
		},
		{
			name:    "test3 - (blank)",
			args:    args{
				address: "",
			},
			wantErr: true,
		},
		{
			name:    "test3 - 123456",
			args:    args{
				address: "123456",
			},
			wantErr: false,
		},
		{
			name:    "test3 - %&$",
			args:    args{
				address: "%&$",
			},
			wantErr: true,
		},
		{
			name:    "test3 - (    )",
			args:    args{
				address: "    ",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkUrlFormat(&tt.args.address); (err != nil) != tt.wantErr {
				t.Errorf("checkUrlFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpCall(t *testing.T) {



	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			//	Can't use every url for the test cases
			//	because in most cases their responses will be different
			//	with each call

			name:    "test - http://example.com",
			args:    args{
				url: "http://example.com/",
			},
			want:    "http://example.com/ 84238dfc8092e5d9c0dac8ef93371a07",
			wantErr: false,
		},
		//TODO: finding some websites with static and fix response
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := httpCall(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("httpCall() got = %v, want %v", got, tt.want)
			}
		})
	}
}