package string

import "testing"

func Test(t *testing.T) {
    var tests = []struct {
        s, want string
    } {
        {"Backward", "drawkcaB"},
        {"Hello, 中国", "国中 ,olleH"},
        {"", ""},
    }

    for _, c := range tests {
        got := Reverse(c.s);
        if got != c.want {
            t.Error("Reverse(%q) == %q, want %q", c.s, got, c.want)
        }
    }
}
