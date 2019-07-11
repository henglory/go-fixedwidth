package fixedwidth

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html/charset"
)

var (
	nilFloat64 *float64
	nilFloat32 *float32
	nilInt     *int
	nilString  *string
)

func float64p(v float64) *float64 { return &v }
func float32p(v float32) *float32 { return &v }
func intp(v int) *int             { return &v }
func int64p(v int64) *int64       { return &v }
func int32p(v int32) *int32       { return &v }
func int16p(v int16) *int16       { return &v }
func int8p(v int8) *int8          { return &v }
func stringp(v string) *string    { return &v }

// EncodableString is a string that implements the encoding TextUnmarshaler and TextMarshaler interface.
// This is useful for testing.
type EncodableString struct {
	S   string
	Err error
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *EncodableString) UnmarshalText(text []byte) error {
	s.S = string(text)
	return s.Err
}

// MarshalText implements encoding.TextUnmarshaler.
func (s EncodableString) MarshalText() ([]byte, error) {
	return []byte(s.S), s.Err
}

func TestThaiMarshal(t *testing.T) {
	people := []struct {
		ID        int     `fixed:"1,5"`
		FirstName string  `fixed:"6,21" cp:"Windows-874"`
		LastName  string  `fixed:"22,30"`
		Grade     float64 `fixed:"31,39"`
	}{
		{1, "นายทดสอบ ทดสอบ", "Lopshire", 99.5},
	}

	data, err := Marshal(people)

	ms874Reader := bytes.NewBuffer(data)
	reader, err := charset.NewReaderLabel("Windows-874", ms874Reader)
	if err != nil {
		t.Error(err)
	}
	nb, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}
	str := []rune(string(nb))
	if string(str[0:5]) != "1    " {
		t.Errorf("%s", string(str[0:5]))
	}
	if string(str[5:21]) != "นายทดสอบ ทดสอบ  " {
		t.Errorf("%s", string(str[5:21]))
	}
	if string(str[21:30]) != "Lopshire " {
		t.Errorf("%s", string(str[21:30]))
	}

}

type peopleStruct struct {
	ID        int     `fixed:"1,5"`
	FirstName string  `fixed:"6,21" cp:"Windows-874"`
	LastName  string  `fixed:"22,30"`
	Grade     float64 `fixed:"31,39"`
}

func TestThaiUnMarshal(t *testing.T) {
	people := peopleStruct{1, "นายทดสอบ ทดสอบ", "Lopshire", 99.5}

	data, err := Marshal(people)

	if err != nil {
		t.Error(err)
	}

	var newPeople peopleStruct

	err = Unmarshal(data, &newPeople)
	if err != nil {
		t.Error(err)
	}
	if newPeople.FirstName != "นายทดสอบ ทดสอบ" {
		t.Errorf("%s", newPeople.FirstName)
	}
}
