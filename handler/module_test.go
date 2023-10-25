package handler

import (
	"context"
	"testing"
)

func TestApp_ModuleSend(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test sending message to module",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				text: "Hello, module!",
			},
			want: "00 0e 11 01 00 00 00 00 ff ff ff ff 48 65 6c 6c 6f 2c 20 6d 6f 64 75 6c 65 21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				ctx: tt.fields.ctx,
			}
			if got := a.ModuleSend(tt.args.text); got != tt.want {
				t.Errorf("App.ModuleSend() = %v, want %v", got, tt.want)
			}
		})
	}
}
