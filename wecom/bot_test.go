package wecom

import (
	"context"
	"testing"
)

func TestSendMarkdown(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
		msg *RobotMsgMarkdown
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "hello world",
			args: args{
				ctx: context.TODO(),
				// https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=
				key: "1ad0dd28-2e9c-4dce-b609-0e27db4729db",
				msg: &RobotMsgMarkdown{
					Content: "> Hello World",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMarkdown(tt.args.ctx, tt.args.key, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendMarkdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendMsg(t *testing.T) {
	type args struct {
		ctx      context.Context
		question string
		reply    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "hello world",
			args: args{
				ctx:      context.TODO(),
				question: "问题",
				reply:    "reply",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMsg(tt.args.ctx, tt.args.question, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("SendMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
