package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(true, IsEmpty(""))
	assertions.Equal(false, IsEmpty(" "))
	assertions.Equal(false, IsEmpty("bob"))
	assertions.Equal(false, IsEmpty(" bob "))
}

func TestIsNotEmpty(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(false, IsNotEmpty(""))
	assertions.Equal(true, IsNotEmpty(" "))
	assertions.Equal(true, IsNotEmpty("bob"))
	assertions.Equal(true, IsNotEmpty(" bob "))
}

func TestIsBlank(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(true, IsBlank(""))
	assertions.Equal(true, IsBlank(" "))
	assertions.Equal(false, IsBlank("bob"))
	assertions.Equal(false, IsBlank(" bob "))
}

func TestIsNotBlank(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(false, IsNotBlank(""))
	assertions.Equal(false, IsNotBlank(" "))
	assertions.Equal(true, IsNotBlank("bob"))
	assertions.Equal(true, IsNotBlank(" bob "))
}

func TestClean(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal("", Clean(""))
	assertions.Equal("abc", Clean("abc"))
	assertions.Equal("abc", Clean("    abc    "))
	assertions.Equal("", Clean("     "))
}

func TestTrim(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal("", Trim(""))
	assertions.Equal("", Trim("     "))
	assertions.Equal("abc", Trim("abc"))
	assertions.Equal("abc", Trim("    abc    "))
	assertions.Equal("ab c", Trim(" ab c "))
}

func TestEquals(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(true, Equals("abc", "abc"))
	assertions.Equal(false, Equals("abc", "ABC"))
	assertions.Equal(false, Equals("", "ABC"))
	assertions.Equal(false, Equals("", " "))
	assertions.Equal(true, Equals("", ""))
}


