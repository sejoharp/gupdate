package sub

import (
	"testing"
)

func TestConsolidate(t *testing.T) {
	type args struct {
		input []Result
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no input",
			args: args{
				input: []Result{},
			},
			want: "Nothing happened!",
		},
		{
			name: "pulling only",
			args: args{
				input: []Result{Pulled{Name: "n1"}, Pulled{Name: "n2"}},
			},
			want: "\nPulled 2 repositories.",
		},
		{
			name: "cloning only",
			args: args{
				input: []Result{Cloned{Name: "n1", Message: "path"}, Cloned{Name: "n2", Message: "path"}},
			},
			want: "\nCloned 2 new repositories:\n\tn1:\tpath\n\tn2:\tpath\n",
		},
		{
			name: "local archived only",
			args: args{
				input: []Result{LocalArchived{Name: "n1", Message: "path"}, LocalArchived{Name: "n2", Message: "path"}},
			},
			want: "\nLocal copies of archived repositories:\n\tn1:\tpath\n\tn2:\tpath\n",
		},
		{
			name: "errors only",
			args: args{
				input: []Result{Error{Name: "n1", Message: "something happened"}, Error{Name: "n2", Message: "something else happened"}},
			},
			want: "\nErrors happened in these repositories:\n\tn1:\tsomething happened\n\tn2:\tsomething else happened\n",
		},
		{
			name: "mixture",
			args: args{
				input: []Result{Pulled{Name: "n1"}, Pulled{Name: "n2"}, Cloned{Name: "n3", Message: "path"}, LocalArchived{Name: "n4", Message: "path"}, Error{Name: "n1", Message: "error message"}},
			},
			want: `
Cloned 1 new repositories:
	n3:	path
Local copies of archived repositories:
	n4:	path
Errors happened in these repositories:
	n1:	error message
Pulled 2 repositories.`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Consolidate(tt.args.input); got != tt.want {
				t.Errorf("Consolidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
