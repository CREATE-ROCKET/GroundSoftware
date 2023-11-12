package handler

import (
	"reflect"
	"testing"
)

func Test_byteArrayToString(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{bytes: []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}},
			want: "[72 101 108 108 111]",
		},
		{
			name: "Test case 2",
			args: args{bytes: []byte{0x47, 0x6F, 0x21}},
			want: "[71 111 33]",
		},
		{
			name: "Real Data",
			args: args{bytes: []byte{0x0f, 0x5a, 0xbb, 0x13, 0x71, 0xc2, 0xa0, 0x03, 0xb1, 0xc2, 0xa0, 0x03, 0xeb, 0x40, 0x47, 0x66, 0x0d, 0x0d, 0xf5, 0x7d, 0x7f, 0x3f, 0x73, 0x7f, 0xc1, 0x3c, 0x03, 0x7d, 0xca, 0x3c, 0x15, 0x94, 0x58, 0xbd, 0x46, 0x71, 0x0d, 0x0d, 0xf9, 0x7d, 0x7f, 0x3f, 0x9f, 0x70, 0xc1, 0x3c, 0x25, 0x79, 0xca, 0x3c, 0xdc, 0x93, 0x58, 0xbd, 0xce, 0x84, 0x0d, 0x0d, 0xf4, 0x7d, 0x7f, 0x3f, 0xd9, 0x66, 0xc1, 0x3c, 0x97, 0x96, 0xca, 0x3c, 0x65, 0x94, 0x58, 0xbd, 0x3f, 0x9c, 0x0d, 0x0d, 0x03, 0x7e, 0x7f, 0x3f, 0x17, 0x89, 0xc1, 0x3c, 0x6f, 0xe5, 0xca, 0x3c, 0x4f, 0x69, 0x58, 0xbd, 0xaf, 0xb3, 0x0d, 0x0d, 0xfa, 0x7d, 0x7f, 0x3f, 0x73, 0x85, 0xc1, 0x3c, 0x7c, 0xeb, 0xca, 0x3c, 0x26, 0x73, 0x58, 0xbd, 0x39, 0xc7, 0x0d, 0x0d, 0xf9, 0x7d, 0x7f, 0x3f, 0x48, 0x7f, 0xc1, 0x3c, 0xc8, 0xf6, 0xca, 0x3c, 0x47, 0x72, 0x58, 0xbd, 0xa6, 0xde, 0x0d, 0x0d, 0xf6, 0x7d, 0x7f, 0x3f, 0x26, 0x63, 0xc1, 0x3c, 0x5d, 0xec, 0xca, 0x3c, 0xcf, 0x7e, 0x58, 0xbd, 0x31, 0xf2, 0x0d, 0x0d, 0xfa, 0x7d, 0x7f, 0x3f, 0xa3, 0x69, 0xc1, 0x3c, 0xa6, 0xf6, 0xca, 0x3c, 0x4d, 0x76, 0x58, 0xbd, 0x3b, 0xfa, 0x0d, 0x0d, 0x9d, 0x6a, 0x40, 0xd4, 0xfe, 0x01, 0x9f, 0x19, 0x00}},
			want: "[15 90 187 19 113 194 160 3 177 194 160 3 235 64 71 102 13 13 245 125 127 63 115 127 193 60 3 125 202 60 21 148 88 189 70 113 13 13 249 125 127 63 159 112 193 60 37 121 202 60 220 147 88 189 206 132 13 13 244 125 127 63 217 102 193 60 151 150 202 60 101 148 88 189 63 156 13 13 3 126 127 63 23 137 193 60 111 229 202 60 79 105 88 189 175 179 13 13 250 125 127 63 115 133 193 60 124 235 202 60 38 115 88 189 57 199 13 13 249 125 127 63 72 127 193 60 200 246 202 60 71 114 88 189 166 222 13 13 246 125 127 63 38 99 193 60 93 236 202 60 207 126 88 189 49 242 13 13 250 125 127 63 163 105 193 60 166 246 202 60 77 118 88 189 59 250 13 13 157 106 64 212 254 1 159 25 0]",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteArrayToString(tt.args.bytes); got != tt.want {
				t.Errorf("byteArrayToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findStartIndex(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Pattern found at the beginning",
			args: args{input: []byte{0x0f, 0x5a, 0xbb, 0x13, 0x71, 0xc2, 0xa0, 0x03}},
			want: 0,
		},
		{
			name: "Pattern found in the middle",
			args: args{input: []byte{0x13, 0x71, 0xc2, 0xa0, 0x03, 0x0f, 0x5a, 0xbb}},
			want: 5,
		},
		{
			name: "Pattern not found",
			args: args{input: []byte{0x01, 0x02, 0x03, 0x04, 0x05}},
			want: -1,
		},
		{
			name: "Real Data",
			args: args{input: []byte{0x0f, 0x5a, 0xbb, 0x13, 0x71, 0xc2, 0xa0, 0x03, 0xb1, 0xc2, 0xa0, 0x03, 0xeb, 0x40, 0x47, 0x66, 0x0d, 0x0d, 0xf5, 0x7d, 0x7f, 0x3f, 0x73, 0x7f, 0xc1, 0x3c, 0x03, 0x7d, 0xca, 0x3c, 0x15, 0x94, 0x58, 0xbd, 0x46, 0x71, 0x0d, 0x0d, 0xf9, 0x7d, 0x7f, 0x3f, 0x9f, 0x70, 0xc1, 0x3c, 0x25, 0x79, 0xca, 0x3c, 0xdc, 0x93, 0x58, 0xbd, 0xce, 0x84, 0x0d, 0x0d, 0xf4, 0x7d, 0x7f, 0x3f, 0xd9, 0x66, 0xc1, 0x3c, 0x97, 0x96, 0xca, 0x3c, 0x65, 0x94, 0x58, 0xbd, 0x3f, 0x9c, 0x0d, 0x0d, 0x03, 0x7e, 0x7f, 0x3f, 0x17, 0x89, 0xc1, 0x3c, 0x6f, 0xe5, 0xca, 0x3c, 0x4f, 0x69, 0x58, 0xbd, 0xaf, 0xb3, 0x0d, 0x0d, 0xfa, 0x7d, 0x7f, 0x3f, 0x73, 0x85, 0xc1, 0x3c, 0x7c, 0xeb, 0xca, 0x3c, 0x26, 0x73, 0x58, 0xbd, 0x39, 0xc7, 0x0d, 0x0d, 0xf9, 0x7d, 0x7f, 0x3f, 0x48, 0x7f, 0xc1, 0x3c, 0xc8, 0xf6, 0xca, 0x3c, 0x47, 0x72, 0x58, 0xbd, 0xa6, 0xde, 0x0d, 0x0d, 0xf6, 0x7d, 0x7f, 0x3f, 0x26, 0x63, 0xc1, 0x3c, 0x5d, 0xec, 0xca, 0x3c, 0xcf, 0x7e, 0x58, 0xbd, 0x31, 0xf2, 0x0d, 0x0d, 0xfa, 0x7d, 0x7f, 0x3f, 0xa3, 0x69, 0xc1, 0x3c, 0xa6, 0xf6, 0xca, 0x3c, 0x4d, 0x76, 0x58, 0xbd, 0x3b, 0xfa, 0x0d, 0x0d, 0x9d, 0x6a, 0x40, 0xd4, 0xfe, 0x01, 0x9f, 0x19, 0x00}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStartIndex(tt.args.input); got != tt.want {
				t.Errorf("findStartIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findData(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Valid data",
			args: args{input: []byte{0x0f, 0x5a, 0xbb, 0x04, 0x01, 0x02, 0x03, 0x04, 0xff, 0xee}},
			want: []byte{0x04, 0x01, 0x02, 0x03, 0x04, 0xff, 0xee},
		},
		{
			name: "Empty data",
			args: args{input: []byte{0x0f, 0x5a, 0xbb, 0x00}},
			want: []byte{},
		},
		{
			name: "Invalid length",
			args: args{input: []byte{0x0f, 0x5a, 0xbb, 0x03, 0x01, 0x02, 0x03}},
			want: nil,
		},
		{
			name: "No start pattern",
			args: args{input: []byte{0x01, 0x02, 0x03, 0x04, 0xff, 0xee}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findData(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findData() = %v, want %v", got, tt.want)
			}
		})
	}
}
