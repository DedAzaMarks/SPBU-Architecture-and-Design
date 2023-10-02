package internal

import "testing"

func TestCat(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cat(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEcho(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Echo(tt.args.args...); got != tt.want {
				t.Errorf("Echo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPwd(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pwd()
		})
	}
}

func TestWc(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
		want3 string
	}{
		{
			args: args{
				filename: "commands.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := Wc(tt.args.filename)
			if got != tt.want {
				t.Errorf("Wc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Wc() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Wc() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("Wc() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
