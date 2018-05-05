package mdopen

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/mdwhatcott/mdopen/internal/templates/github"
)

const md = `# test`

func TestOpenerOpen(t *testing.T) {
	f := strings.NewReader(md)
	opener := New(echoCMD())
	if err := opener.Open(f); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestOpener_prepareFile(t *testing.T) {
	opener := New()

	tests := []struct {
		name     string
		r        io.Reader
		wantErr  bool
		contains string
	}{
		{
			name:     "valid",
			r:        strings.NewReader(md),
			wantErr:  false,
			contains: "<h1>test</h1>",
		},
		{
			name:    "with error",
			r:       new(failingReader),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := new(bytes.Buffer)

			err := opener.prepareFile(w, tt.r)
			if tt.wantErr && err == nil {
				t.Errorf("expected error")
				return
			}

			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %s", err)
				return
			}

			if !strings.Contains(w.String(), tt.contains) {
				t.Errorf("body does not contains expected string: %s", tt.contains)
				if err := opener.Open(tt.r); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *Opener
	}{
		{
			name: "default",
			args: args{
				options: []Option{},
			},
			want: &Opener{
				cmdName: commandName(),
				layout:  template.Must(template.New("layout").Parse(github.Template)),
			},
		},
		{
			name: "github",
			args: args{
				options: []Option{
					GithubTemplate(),
				},
			},
			want: &Opener{
				cmdName: commandName(),
				layout:  template.Must(template.New("layout").Parse(github.Template)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

type failingReader struct{}

func (r *failingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("unexpected EOF")
}

func echoCMD() Option {
	return func(opener *Opener) {
		opener.cmdName = "echo"
	}
}
