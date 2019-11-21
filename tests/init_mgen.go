package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/internal/build"
)

// This file holds variable and type relating specifically
// to the task of generating tests.

var (
	MG_GENERATE_STATE_TESTS_KEY      = "MULTIGETH_TESTS_GENERATE_STATE_TESTS"
	MG_GENERATE_DIFFICULTY_TESTS_KEY = "MULTIGETH_TESTS_GENERATE_DIFFICULTY_TESTS"
	MG_CHAINCONFIG_FEATURE_EQ_KEY    = "MULTIGETH_TESTS_CHAINCONFIG_FEATURE_EQUIVALANCE"
	MG_CHAINCONFIG_CHAINSPEC_KEY     = "MULTIGETH_TESTS_CHAINCONFIG_PARITY_SPECS"
)

type chainspecRefsT map[string]chainspecRef

var chainspecRefsState = chainspecRefsT{}
var chainspecRefsDifficulty = chainspecRefsT{}

type chainspecRef struct {
	Filename string `json:"filename"`
	Sha1Sum  []byte `json:"sha1sum"`
}

func (c chainspecRef) String() string {
	return fmt.Sprintf("file: %s, file.sha1sum: %x", c.Filename, c.Sha1Sum)
}

func (c *chainspecRef) UnmarshalJSON(input []byte) error {
	type xT struct {
		F string `json:"filename"`
		S string `json:"sha1sum"`
	}
	var x = xT{}
	err := json.Unmarshal(input, &x)
	if err != nil {
		return err
	}
	c.Filename = x.F
	c.Sha1Sum = common.Hex2Bytes(x.S)
	return nil
}

func (c chainspecRef) MarshalJSON() ([]byte, error) {
	var x = struct {
		F string `json:"filename"`
		S string `json:"sha1sum"`
	}{
		F: c.Filename,
		S: common.Bytes2Hex(c.Sha1Sum[:]),
	}

	return json.MarshalIndent(x, "", "    ")
}

// submoduleParentRef captures the current git status of the tests submodule.
// This is used for reference when writing tests.
var submoduleParentRef = func() string {
	subModOut := build.RunGit("submodule", "status")
	subModOut = strings.ReplaceAll(strings.TrimSpace(subModOut), " ", "_")
	return subModOut
}()

