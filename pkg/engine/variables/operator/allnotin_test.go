package operator

import (
	"testing"

	"github.com/go-logr/logr"

	"github.com/kyverno/kyverno/pkg/engine/context"
)

func TestAllNotInHandler_Evaluate(t *testing.T) {
	type fields struct {
		ctx context.EvalInterface
		log logr.Logger
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "key is string and in value",
			args: args{
				key:   "kyverno",
				value: "kyverno",
			},
			want: false,
		},
		{
			name: "key is string and in value",
			args: args{
				key:   "kube-system",
				value: []interface{}{"default", "kube-*"},
			},
			want: false,
		},
		{
			name: "key is string and not in value",
			args: args{
				key:   "kyverno",
				value: "default",
			},
			want: true,
		},

		{
			name: "key is int and in value",
			args: args{
				key:   64,
				value: "64",
			},
			want: false,
		},
		{
			name: "key is int and not in value",
			args: args{
				key:   64,
				value: "default",
			},
			want: true,
		},
		{
			name: "key is array and in value",
			args: args{
				key:   []interface{}{"kube-system", "kube-public"},
				value: "kube-*",
			},
			want: false,
		},
		{
			name: "key is array and partially in value",
			args: args{
				key:   []interface{}{"kube-system", "default"},
				value: "kube-system",
			},
			want: false,
		},
		{
			name: "key is array and not in value",
			args: args{
				key:   []interface{}{"default", "kyverno"},
				value: "kube-*",
			},
			want: true,
		},
		{
			name: "key and value are array and not in value",
			args: args{
				key:   []interface{}{"default", "kyverno"},
				value: []interface{}{"kube-*", "kube-system"},
			},
			want: true,
		},
		{
			name: "key and value are array and partially in value",
			args: args{
				key:   []interface{}{"default", "kyverno"},
				value: []interface{}{"kube-*", "ky*"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allnin := AllNotInHandler{
				ctx: tt.fields.ctx,
				log: tt.fields.log,
			}
			if got := allnin.Evaluate(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
