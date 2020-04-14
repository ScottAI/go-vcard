package go_vcard

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

var testCard = Card{
	"VERSION": []*Property{{Value:[][]string{{"4.0"}}}},
	"UID":     []*Property{{Value: [][]string{{"urn","uuid","4fbe8971-0bc3-424c-9c26-36c3e1eff6b1"}}}},
	"FN": []*Property{{
		Value:  [][]string{{"J. Doe"}},
		Params: map[string][]string{"PID": {"1.1"}},
	}},
	"N": []*Property{{Value: [][]string{{"Doe"},{"J."},{""},{""},{""}}}},
	"EMAIL": []*Property{{
		Value:  [][]string{{"jdoe@example.com"}},
		Params: map[string][]string{"PID": {"1.1"}},
	}},
	"CLIENTPIDMAP": []*Property{{Value: [][]string{{"1"},{"urn","uuid","53e374d9-337e-4727-8803-a1e9c14e0556"}}}},
}

var testCardHandmade = Card{
	"VERSION": []*Property{{Value: [][]string{{"4.0"}}}},
	"N":       []*Property{{Value: [][]string{{"Bloggs"},{"Joe"},{""},{""},{""}}}},
	"FN":      []*Property{{Value: [][]string{{"Joe Bloggs"}}}},
	"EMAIL": []*Property{{
		Value:  [][]string{{"me@joebloggs.com"}},
		Params: map[string][]string{"TYPE":{"home"},"PREF":{"1"}},
	}},
	"TEL": []*Property{{
		Value:  [][]string{{"tel","+44 20 1234 5678"}},
		Params: map[string][]string{"TYPE": {"cell", "home"}, "PREF": {"1"}},
	}},
	"ADR": []*Property{{
		Value:  [][]string{{""},{""},{"1 Trafalgar Square"},{"London"},{""},{"WC2N"},{"United Kingdom"}},
		Params: map[string][]string{"TYPE": {"home"}, "PREF": {"1"}},
	}},
	"URL": []*Property{{
		Value:  [][]string{{"http://joebloggs.com"}},
		Params: map[string][]string{"TYPE": {"home"}, "PREF": {"1"}},
	}},
	"IMPP": []*Property{{
		Value:  [][]string{{"skype:joe.bloggs"}},
		Params: map[string][]string{"TYPE": {"home"}, "PREF": {"1"}},
	}},
	"X-SOCIALPROFILE": []*Property{{
		Value:  [][]string{{"twitter:https://twitter.com/joebloggs"}},
		Params: map[string][]string{"TYPE": {"home"}, "PREF": {"1"}},
	}},
}

var testCardGoogle = Card{
	"VERSION": []*Property{{Value: [][]string{{"3.0"}}}},
	"N":       []*Property{{Value: [][]string{{"Bloggs"},{"Joe"},{""},{""}}}},
	"FN":      []*Property{{Value: [][]string{{"Joe Bloggs"}}}},
	"EMAIL": []*Property{{
		Value:  [][]string{{"me@joebloggs.com"}},
		Params: map[string][]string{"TYPE": {"INTERNET", "HOME"}},
	}},
	"TEL": []*Property{{
		Value:  [][]string{{"+44 20 1234 5678"}},
		Params: map[string][]string{"TYPE": {"CELL"}},
	}},
	"ADR": []*Property{{
		Value:  [][]string{{""},{""},{"1 Trafalgar Square"},{"London"},{""},{"WC2N"},{"United Kingdom"}},
		Params: map[string][]string{"TYPE": {"HOME"}},
	}},
	"URL": []*Property{
		{Value: [][]string{{"http\\://joebloggs.com"}}, Group: "item1"},
		{Value: [][]string{{"http\\://twitter.com/test"}}, Group: "item2"},
	},
	"X-SKYPE": []*Property{{Value: [][]string{{"joe.bloggs"}}}},
	"X-ABLABEL": []*Property{
		{Value: [][]string{{"_$!<HomePage>!$_"}}, Group: "item1"},
		{Value: [][]string{{"Twitter"}}, Group: "item2"},
	},
}

var testCardApple = Card{
	"VERSION": []*Property{{Value: [][]string{{"3.0"}}}},
	"N":       []*Property{{Value: [][]string{{"Bloggs"},{"Joe"},{""},{""}}}},
	"FN":      []*Property{{Value: [][]string{{"Joe Bloggs"}}}},
	"EMAIL": []*Property{{
		Value:  [][]string{{"me@joebloggs.com"}},
		Params: map[string][]string{"TYPE": {"INTERNET", "HOME", "pref"}},
	}},
	"TEL": []*Property{{
		Value:  [][]string{{"+44 20 1234 5678"}},
		Params: map[string][]string{"TYPE": {"CELL", "VOICE", "pref"}},
	}},
	"ADR": []*Property{{
		Value:  [][]string{{""},{""},{"1 Trafalgar Square"},{"London"},{},{"WC2N"},{"United Kingdom"}},
		Params: map[string][]string{"TYPE": {"HOME", "pref"}},
	}},
	"URL": []*Property{{
		Value:  [][]string{{"http://joebloggs.com"}},
		Params: map[string][]string{"TYPE": {"pref"}},
		Group:  "item1",
	}},
	"X-ABLABEL": []*Property{
		{Value: [][]string{{"_$!<HomePage>!$_"}}, Group: "item1"},
	},
	"IMPP": []*Property{{
		Value:  [][]string{{"skype:joe.bloggs"}},
		Params: map[string][]string{"X-SERVICE-TYPE": {"Skype"}, "TYPE": {"HOME", "pref"}},
	}},
	"X-SOCIALPROFILE": []*Property{{
		Value:  [][]string{{"https://twitter.com/joebloggs"}},
		Params: map[string][]string{"TYPE": {"twitter"}},
	}},
}



func TestCard(t *testing.T) {
	testCardFullName := testCard["FN"][0]
	if field := testCard.Get(PropFN); testCardFullName != field {
		t.Errorf("Expected card FN field to be %+v but got %+v", testCardFullName, field)
	}
	if v := testCard.Value(PropFN); MatrixToString(v) != MatrixToString(testCardFullName.Value) {
		t.Errorf("Expected card FN field to be %q but got %q", MatrixToString(testCardFullName.Value), MatrixToString(v))
	}

	if field := testCard.Get("X-IDONTEXIST"); field != nil {
		t.Errorf("Expected card X-IDONTEXIST field to be %+v but got %+v", nil, field)
	}
	if v := testCard.Value("X-IDONTEXIST"); MatrixToString(v) != "" {
		t.Errorf("Expected card X-IDONTEXIST field value to be %q but got %q", "", v)
	}

	cardMultipleValues := Card{
		"EMAIL": []*Property{
			{Value: [][]string{{"me@example.org"}}, Params: map[string][]string{"TYPE": {"home"}}},
			{Value: [][]string{{"me@example.com"}}, Params: map[string][]string{"TYPE": {"work"}}},
		},
	}
	expected := [][][]string{{{"me@example.org"}}, {{"me@example.com"}}}
	if values := cardMultipleValues.Values(PropEmail); !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected card emails to be %+v but got %+v", expected, values)
	}
	if values := cardMultipleValues.Values("X-IDONTEXIST"); values != nil {
		t.Errorf("Expected card X-IDONTEXIST values to be %+v but got %+v", nil, values)
	}
}

func TestCard_AddValue(t *testing.T) {
	card := make(Card)

	name1 := "Akiyama Mio"
	card.AddValue("FN", [][]string{{name1}})
	if values := card.Values("FN")[0][0][0]; values != name1 {
		t.Errorf("Expected one FN value, got %v", values)
	}

	name2 := "Mio Akiyama"
	card.AddValue("FN", [][]string{{name2}})
	if values := card.Values("FN"); len(values) != 2 || values[0][0][0] != name1 || values[1][0][0] != name2 {
		t.Errorf("Expected two FN values, got %v", values)
	}
}

func TestCard_Preferred(t *testing.T) {
	if pref := testCard.Pref("X-IDONTEXIST"); pref != nil {
		t.Errorf("Expected card preferred X-IDONTEXIST field to be %+v but got %+v", nil, pref)
	}
	if v := testCard.PrefValue("X-IDONTEXIST"); MatrixToString(v) != "" {
		t.Errorf("Expected card preferred X-IDONTEXIST field value to be %q but got %q", "", v)
	}

	cards := []Card{
		{
			"EMAIL": []*Property{
				{Value: [][]string{{"me@example.org"}}, Params: map[string][]string{"TYPE": {"home"}}},
				{Value: [][]string{{"me@example.com"}}, Params: map[string][]string{"TYPE": {"work"}, "PREF": {"1"}}},
			},
		},
		{
			"EMAIL": []*Property{
				{Value: [][]string{{"me@example.org"}}, Params: map[string][]string{"TYPE": {"home"}, "PREF": {"25"}}},
				{Value: [][]string{{"me@example.com"}}, Params: map[string][]string{"TYPE": {"work"}, "PREF": {"50"}}},
			},
		},
		// v3.0
		{
			"EMAIL": []*Property{
				{Value: [][]string{{"me@example.org"}}, Params: map[string][]string{"TYPE": {"home"}}},
				{Value: [][]string{{"me@example.com"}}, Params: map[string][]string{"TYPE": {"work", "pref"}}},
			},
		},
	}

	for _, card := range cards {
		if pref := card.Pref(PropEmail); pref != card["EMAIL"][1] {
			t.Errorf("Expected card preferred email to be %+v but got %+v", card["EMAIL"][1], pref)
		}
		if v := card.PrefValue(PropEmail); !strings.EqualFold(MatrixToString(v), "me@example.com") {
			t.Errorf("Expected card preferred email to be %q but got %q", "me@example.com", MatrixToString(v))
		}
	}
}

func TestCard_Name(t *testing.T) {
	card := make(Card)
	if name := card.Name(); name != nil {
		t.Errorf("Expected empty card name to be %+v but got %+v", nil, name)
	}
	if names := card.Names(); names != nil {
		t.Errorf("Expected empty card names to be %+v but got %+v", nil, names)
	}

	expectedName := &Name{
		FamilyName: "Doe",
		GivenName:  "J.",
	}
	expectedNames := []*Name{expectedName}
	card.AddName(expectedName)
	if name := card.Name(); !reflect.DeepEqual(expectedName, name) {
		t.Errorf("Expected populated card name to be %+v but got %+v", expectedName, name)
	}
	if names := card.Names(); !reflect.DeepEqual(expectedNames, names) {
		t.Errorf("Expected populated card names to be %+v but got %+v", expectedNames, names)
	}
}

func TestCard_Kind(t *testing.T) {
	card := make(Card)

	if kind := card.Kind(); kind != KindIndividual {
		t.Errorf("Expected kind of empty card to be %q but got %q", KindIndividual, kind)
	}

	card.SetKind(KindOrg)
	if kind := card.Kind(); kind != KindOrg {
		t.Errorf("Expected kind of populated card to be %q but got %q", KindOrg, kind)
	}
}

func TestCard_FormattedNames(t *testing.T) {
	card := make(Card)

	expectedNames := []*Property{{Value: [][]string{{""}}}}
	if names := card.FormattedNames(); !reflect.DeepEqual(expectedNames, names) {
		t.Errorf("Expected empty card formatted names to be %+v but got %+v", expectedNames, names)
	}

	expectedNames = []*Property{{Value: [][]string{{"Akiyama Mio"}}}}
	card.SetValue(PropFN, expectedNames[0].Value)
	if names := card.FormattedNames(); !reflect.DeepEqual(expectedNames, names) {
		t.Errorf("Expected populated card formatted names to be %+v but got %+v", expectedNames, names)
	}
}

func TestCard_Gender(t *testing.T) {
	card := make(Card)

	var expectedSex string
	var expectedIdentity string
	if sex, identity := card.Gender(); sex != expectedSex || identity != expectedIdentity {
		t.Errorf("Expected gender to be (%q %q) but got (%q %q)", expectedSex, expectedIdentity, sex, identity)
	}

	expectedSex = SexFemale
	card.SetGender(expectedSex, expectedIdentity)
	if sex, identity := card.Gender(); sex != expectedSex || identity != expectedIdentity {
		t.Errorf("Expected gender to be (%q %q) but got (%q %q)", expectedSex, expectedIdentity, sex, identity)
	}

	expectedSex = SexOther
	expectedIdentity = "<3"
	card.SetGender(expectedSex, expectedIdentity)
	if sex, identity := card.Gender(); sex != expectedSex || identity != expectedIdentity {
		t.Errorf("Expected gender to be (%q %q) but got (%q %q)", expectedSex, expectedIdentity, sex, identity)
	}
}

func TestCard_Address(t *testing.T) {
	card := make(Card)

	if address := card.Address(); address != nil {
		t.Errorf("Expected empty card address to be nil, got %v", address)
	}
	if addresses := card.Addresses(); addresses != nil {
		t.Errorf("Expected empty card addresses to be nil, got %v", addresses)
	}

	added := &Address{
		StreetAddress: "1 Trafalgar Square",
		Locality:      "London",
		PostalCode:    "WC2N",
		Country:       "United Kingdom",
	}
	card.AddAdress(added)

	equal := func(a, b *Address) bool {
		if (a == nil && b != nil) || (b == nil && a != nil) {
			return false
		}
		a.Property, b.Property = nil, nil
		return reflect.DeepEqual(a, b)
	}

	if address := card.Address(); !equal(added, address) {
		t.Errorf("Expected address to be %+v but got %+v", added, address)
	}
	if addresses := card.Addresses(); len(addresses) != 1 || !equal(added, addresses[0]) {
		t.Errorf("Expected addresses to be %+v, got %+v", []*Address{added}, addresses)
	}
}

func TestCard_Revision(t *testing.T) {
	card := make(Card)

	if rev, err := card.Revision(); err != nil {
		t.Fatal("Expected no error when getting revision of an empty card, got:", err)
	} else if !rev.IsZero() {
		t.Error("Expected a zero time when getting revision of an empty card, got:", rev)
	}

	expected := time.Date(1984, time.November, 4, 0, 0, 0, 0, time.UTC)
	card.SetRevision(expected)
	if rev, err := card.Revision(); err != nil {
		t.Fatal("Expected no error when getting revision of a populated card, got:", err)
	} else if !rev.Equal(rev) {
		t.Errorf("Expected revision to be %v but got %v", expected, rev)
	}
}

