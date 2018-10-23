package toyaml_test

import (
	"bytes"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/ktr0731/toyaml"
	toml "github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type example struct {
	Foo string
	Bar struct {
		HogeFuga int
	}
}

func TestToYAML(t *testing.T) {
	val := example{
		Foo: "foo",
		Bar: struct {
			HogeFuga int
		}{100},
	}

	b, err := toml.Marshal(val)
	require.NoError(t, err)

	assertValue := func(t *testing.T, b []byte) {
		var actual example
		err := yaml.Unmarshal(b, &actual)
		require.NoError(t, err)
		assert.EqualValues(t, val, actual)
	}

	t.Run("[]byte", func(t *testing.T) {
		b, err := toyaml.ToYAML(b)
		require.NoError(t, err)
		assertValue(t, b)
	})

	t.Run("string", func(t *testing.T) {
		b, err := toyaml.ToYAML(string(b))
		require.NoError(t, err)
		assertValue(t, b)
	})

	t.Run("io.Reader", func(t *testing.T) {
		b, err := toyaml.ToYAML(bytes.NewBuffer(b))
		require.NoError(t, err)
		assertValue(t, b)
	})
}
