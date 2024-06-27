package cache

import (
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestRedisCache_Set(t *testing.T) {
	type fields struct {
		client *redis.Client
		prefix string
	}
	type args struct {
		key    string
		value  string
		expire time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "TestRedisCache_Set",
			fields: fields{
				client: redis.NewClient(&redis.Options{
					Addr:     "192.168.1.197:6379",
					Password: "123456",
					DB:       0,
				}),
				prefix: "test",
			},
			args: args{
				key:    "test",
				value:  "test123",
				expire: time.Second * 60,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisCache{
				client: tt.fields.client,
				prefix: tt.fields.prefix,
			}
			if err := r.Set(tt.args.key, tt.args.value, tt.args.expire); err != nil {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisCache_Has(t *testing.T) {
	type fields struct {
		client *redis.Client
		prefix string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "TestRedisCache_Has",
			fields: fields{
				client: redis.NewClient(&redis.Options{
					Addr:     "192.168.1.197:6379",
					Password: "123456",
					DB:       0,
				}),
				prefix: "test",
			},
			args: args{
				key: "test",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisCache{
				client: tt.fields.client,
				prefix: tt.fields.prefix,
			}
			if got := r.Has(tt.args.key); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCache_Get(t *testing.T) {
	type fields struct {
		client *redis.Client
		prefix string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "TestRedisCache_Get",
			fields: fields{
				client: redis.NewClient(&redis.Options{
					Addr:     "192.168.1.197:6379",
					Password: "123456",
					DB:       0,
				}),
				prefix: "test",
			},
			args: args{
				key: "test",
			},
			want: "test123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisCache{
				client: tt.fields.client,
				prefix: tt.fields.prefix,
			}
			if got := r.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCache_Del(t *testing.T) {
	type fields struct {
		client *redis.Client
		prefix string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "TestRedisCache_Del",
			fields: fields{
				client: redis.NewClient(&redis.Options{
					Addr:     "192.168.1.197:6379",
					Password: "123456",
					DB:       0,
				}),
				prefix: "test",
			},
			args: args{
				key: "test",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisCache{
				client: tt.fields.client,
				prefix: tt.fields.prefix,
			}
			if err := r.Del(tt.args.key); err != nil {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
