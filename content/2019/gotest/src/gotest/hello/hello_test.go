package hello

import (
	"fmt"
	"reflect"
	"testing"
	"testing/quick"
)

// START1 OMIT
func TestHello(t *testing.T) {
	s := Hello("Gopher")
	want := "Hello, Gopher!"
	if s != want {
		t.Fatalf(`Hello("Gopher") = %s"; want "%s"`, s, want)
	}
}

// END1 OMIT

// START2 OMIT
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Gopher")
	}
}

// END2 OMIT

// START3 OMIT

func ExampleHello() {
	fmt.Println(Hello("Gopher"))

	// Output: Hello, Gopher!
}

// END3 OMIT

// START4 OMIT
func TestHelloShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped")
	}

	// 後続の時間のかかるテスト
}

// END4 OMIT

// START5 OMIT
func TestHelloSubTest(t *testing.T) {
	// 前処理etc

	t.Run("Hoge", func(t *testing.T) {
		s := Hello("Hoge")
		want := "Hello, Hoge!"
		if s != want {
			t.Fatalf(`Hello("Hoge") = %s"; want "%s"`, s, want)
		}
	})
	t.Run("Moke", func(t *testing.T) {
		// さらにネストできる
		t.Run("Fuga", func(t *testing.T) {
			s := Hello("Fuga")
			want := "Hello, Fuga!"
			if s != want {
				t.Fatalf(`Hello("Fuga") = %s"; want "%s"`, s, want)
			}
		})
	})

	// 後処理etc
}

// END5 OMIT

// START6 OMIT
func TestHelloTableDriven(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"Hoge", "Hello, Hoge!"},
		{"Moke", "Hello, Moke!"},
		{"Fuga", "Hello, Fuga!"},
		{"Bosukete", "Hello, Bosukete!"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			s := Hello(tt.in)
			if s != tt.out {
				t.Fatalf(`Hello("Fuga") = %s"; want "%s"`, s, tt.out)
			}
		})
	}
}

// END6 OMIT

// START7 OMIT
func assertEquals(t *testing.T, actual, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(actual, want) {
		// t.Helper()を呼ばないと↓のファイル名、行番号が出力されてしまう
		t.Errorf("not equals; actual:%v, want:%v", actual, want)
	}
}

func TestHelloHelper(t *testing.T) {
	assertEquals(t, Hello("Hoge"), "Hello, Hoge!")
	assertEquals(t, Hello("Moke"), "Hello, Moke!")
	assertEquals(t, Hello("Fuga"), "Hello, Fuga!")
}

// END7 OMIT

// START8 OMIT
func TestHelloQuick(t *testing.T) {
	oldHello := func(name string) string {
		return "Hello, " + name + "!"
	}

	f := func(s string) bool {
		return Hello(s) == oldHello(s)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// END8 OMIT
