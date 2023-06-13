package any

var toStringTest = []struct {
	s        any
	expected string
}{
	{
		1,
		"1",
	},
	{
		1.1,
		"1.1",
	},
	{
		"123",
		"123",
	},
	{
		'A',
		"65",
	},
	{
		map[string]string{"1": "1"},
		"{\"1\":\"1\"}",
	},
}

func TestToString(t *testing.T) {
	for _, test := range toStringTest {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, ToString(test.s))
		})
	}
}
