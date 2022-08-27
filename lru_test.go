package lru

import (
	"reflect"
	"sync"
	"testing"
)

type t1 struct {
	a string
	b int
	c float64
	d rune
}

type t2 struct {
	e rune
	f float64
	g int
	h string
}

var cacheForTesting = testPreparation()

func testPreparation() Cache {
	cacheV := New(100)

	for i := 150; i < 300; i++ {
		cacheV.Store(i, i)
	}

	cacheV.Store("123", "123")
	cacheV.Store(123, 123)
	cacheV.Store(1.23, 1.23)
	cacheV.Store('r', 'r')
	cacheV.Store(
		t1{
			a: "123",
			b: 123,
			c: 1.23,
			d: 'e',
		},
		t2{
			e: 'e',
			f: 1.23,
			g: 123,
			h: "123",
		},
	)

	return cacheV
}

func TestNew(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want Cache
	}{
		{
			name: "Size < 100",
			args: args{
				size: 10,
			},
			want: &cache{
				data:     make(map[string]dataCache, 0),
				latch:    sync.Mutex{},
				capacity: 100,
				border:   98,
				cleaning: false,
			},
		},
		{
			name: "Size > 100",
			args: args{
				size: 10000,
			},
			want: &cache{
				data:     make(map[string]dataCache, 0),
				latch:    sync.Mutex{},
				capacity: 10000,
				border:   9800,
				cleaning: false,
			},
		},
		{
			name: "Not normal size",
			args: args{
				size: 236,
			},
			want: &cache{
				data:     make(map[string]dataCache, 0),
				latch:    sync.Mutex{},
				capacity: 236,
				border:   196,
				cleaning: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cache_Store(t *testing.T) {
	type args struct {
		key   any
		value any
	}

	tests := []struct {
		name string
		c    Cache
		args args
		want bool
	}{
		{
			name: "Save #1 (string)",
			c:    New(100),
			args: args{
				key:   "2 + 2",
				value: "4",
			},
			want: true,
		},
		{
			name: "Save #2 (int)",
			c:    New(100),
			args: args{
				key:   6,
				value: 8,
			},
			want: true,
		},
		{
			name: "Save #3 (float)",
			c:    New(100),
			args: args{
				key:   6.78,
				value: 1.48,
			},
			want: true,
		},
		{
			name: "Save #4 (rune)",
			c:    New(100),
			args: args{
				key:   'a',
				value: 'b',
			},
			want: true,
		},
		{
			name: "Save #5 (struct)",
			c:    New(100),
			args: args{
				key: t1{
					a: "123",
					b: 123,
					c: 1.23,
					d: 'e',
				},
				value: t2{
					e: 'e',
					f: 1.23,
					g: 123,
					h: "123",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Store(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("cache.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cache_Load(t *testing.T) {
	type args struct {
		key any
	}
	tests := []struct {
		name    string
		c       Cache
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Load #1 (string)",
			c:    cacheForTesting,
			args: args{
				key: "123",
			},
			want:    "123",
			wantErr: false,
		},
		{
			name: "Load #2 (int)",
			c:    cacheForTesting,
			args: args{
				key: 123,
			},
			want:    123,
			wantErr: false,
		},
		{
			name: "Load #3 (float)",
			c:    cacheForTesting,
			args: args{
				key: 1.23,
			},
			want:    1.23,
			wantErr: false,
		},
		{
			name: "Load #4 (rune)",
			c:    cacheForTesting,
			args: args{
				key: 'r',
			},
			want:    'r',
			wantErr: false,
		},
		{
			name: "Load #5 (struct)",
			c:    cacheForTesting,
			args: args{
				key: t1{
					a: "123",
					b: 123,
					c: 1.23,
					d: 'e',
				},
			},
			want: t2{
				e: 'e',
				f: 1.23,
				g: 123,
				h: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Load(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("cache.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cache.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
