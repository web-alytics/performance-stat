package pgsql

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validation(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		exp  bool
	}{
		{
			desc: "valid string inp",
			in:   t.Name(),
			exp:  false,
		},
		{
			desc: "in valid string inp",
			in:   "",
			exp:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := isEmpty(tC.in)
			if got != tC.exp {
				t.Fatalf("got:%v\n,exp:%v\n", got, tC.exp)
			}

		})
	}
}

func Test_conn_string(t *testing.T) {
	got := getConnFromEnv()
	assert.NotEmpty(t, got)
}

func Test_new_connection(t *testing.T) {
	testCases := []struct {
		desc     string
		setupEnv bool
		expErr   bool
	}{
		{
			desc:     "in-valid conn string",
			setupEnv: true,
			expErr:   true,
		},
		{
			desc:     "valid connection",
			setupEnv: false,
			expErr:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.setupEnv {
				setupEnv(t)
				defer unsetEnv(t)
			}
			got, err := New()
			assert.NotNil(t, got)
			if tC.expErr {
				t.Logf("err received:%v\n", err)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

func unsetEnv(t *testing.T) {
	if err := os.Unsetenv(USERNAME); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv(PASSWORD); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv(HOST); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv(DBNAME); err != nil {
		t.Fatal(err)
	}

}

func setupEnv(t *testing.T) {
	os.Setenv(USERNAME, t.Name())
	os.Setenv(PASSWORD, t.Name())
	os.Setenv(HOST, t.Name())
	os.Setenv(DBNAME, t.Name())

}
