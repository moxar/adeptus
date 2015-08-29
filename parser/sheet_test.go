package parser

import (
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_ParseSheet(t *testing.T) {
	cases := []struct {
		in    io.Reader
		out   Sheet
		err   bool
		panic bool
	}{
		{
			in:    nil,
			out:   Sheet{},
			err:   false,
			panic: true,
		},
		{
			in: strings.NewReader(`#  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum rhoncus porta tellus, eget placerat libero dictum ac. Sed massa ex, vehicula eget egestas quis, rutrum non quam. Quisque blandit lacus ac erat posuere pellentesque. Integer vitae eleifend nisi. Duis hendrerit facilisis blandit. Ut sed semper eros. Vestibulum laoreet consequat leo id interdum. Ut mauris justo, tincidunt ac tortor eu, auctor auctor eros. 
//  Duis id tincidunt lacus. Etiam a tincidunt urna. In et sapien vitae eros tempus cursus vitae dignissim felis. Etiam eget scelerisque nibh, eget maximus quam. Sed vitae elit id velit feugiat elementum. Praesent viverra tincidunt mi, nec posuere diam sodales ut. Fusce lectus nisi, venenatis sit amet accumsan placerat, luctus nec nulla. Donec eget metus sed ante ornare iaculis. 
`),
			out:   Sheet{},
			err:   false,
			panic: true,
		},
		{
			in: strings.NewReader(`Babar: Celeste

2015/08/10 Babar love celeste
	* Peanuts [200]
`),
			out:   Sheet{},
			err:   true,
			panic: false,
		},
		{
			in: strings.NewReader(`Name: Celeste

Babar love celeste
Babar also love peanuts
`),
			out:   Sheet{},
			err:   true,
			panic: false,
		},
		{
			in: strings.NewReader(`Name: Someone
Origin: Somewhere
Background: Something
Role: Warmonger
Tarot: XXI
`),
			out: Sheet{
				Header: Header{
					Name:       StringP("Someone"),
					Origin:     StringP("Somewhere"),
					Background: StringP("Something"),
					Role:       StringP("Warmonger"),
					Tarot:      StringP("XXI"),
				},
				Sessions: []Session{},
			},
			err:   false,
			panic: false,
		},
		{
			in: strings.NewReader(`Name: Someone
Origin: Somewhere
Background: Something
Role: Warmonger
Tarot: XXI

2015/06/01 Creation [500]
	* BULLSHIT +5 [250]
	- Awesomeskill
`),
			out: Sheet{
				Header: Header{
					Name:       StringP("Someone"),
					Origin:     StringP("Somewhere"),
					Background: StringP("Something"),
					Role:       StringP("Warmonger"),
					Tarot:      StringP("XXI"),
				},
				Sessions: []Session{
					Session {
						Date:   time.Date(2015, time.June, 01, 0, 0, 0, 0, time.UTC),
						Title:  "Creation",
						Reward: IntP(500),
						Upgrades: []Upgrade{
							Upgrade{
								Mark: "*",
								Name: "BULLSHIT +5",
								Cost: IntP(250),
							},
							Upgrade{
								Mark: "-",
								Name: "Awesomeskill",
								Cost: nil,
							},
						},
					},
				},
			},
			err:   false,
			panic: false,
		},
	}

	for i, c := range cases {
		out, err, panic := func() (out Sheet, err error, panic bool) {
			defer func() {
				if e := recover(); e != nil {
					panic = true
				}
			}()

			out, err = ParseSheet(c.in)
			return
		}()

		if (err != nil) != c.err {
			if err == nil {
				t.Logf("Expected error on case %d", i+1)
			} else {
				t.Logf("Unexpected error on case %d: %s", i+1, err)
			}
			t.Fail()
		}

		if !reflect.DeepEqual(out, c.out) {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}

		if panic != c.panic {
			if panic {
				t.Logf("Unexpected panic on case %d", i+1)
			} else {
				t.Logf("Should panic on case %d", i+1)
			}
			t.Fail()
		}
	}

}
