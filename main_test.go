package main

import (
	"flag"
	"os"
	"testing"
)

const (
	GITHUB_SECRETS = `{
		"github_token": "ghs_nlYySwvlS0eEo2gArV1uyAHb8EQ6hK3kaooJ",
		"SECRET_SINGLE_LINE": "dummySecretValue",
		"SECRET_MULTI_LINE": "this\nis\na multiline\nsecret"
	}`
	SECRET_ID = "github_token"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Setenv(INPUT_GITHUB_SECRETS, GITHUB_SECRETS)
	os.Exit(m.Run())
}

func TestMainFunc(t *testing.T) {
	// Redirect standard out to null
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
		os.Args = os.Args[:len(os.Args)-1]
	}()
	os.Stdout = os.NewFile(0, os.DevNull)

	os.Args = append(os.Args, SECRET_ID)
	main()
}

func TestVersionFlag(t *testing.T) {
	// Redirect standard out to null
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
		os.Args = os.Args[:len(os.Args)-1]
	}()
	os.Stdout = os.NewFile(0, os.DevNull)

	os.Args = append(os.Args, "-V")
	main()
}

func TestFindSecretVal(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"SECRET_MULTI_LINE", "this\nis\na multiline\nsecret"},
		{"SECRET_SINGLE_LINE", "dummySecretValue"},
		{SECRET_ID, "ghs_nlYySwvlS0eEo2gArV1uyAHb8EQ6hK3kaooJ"},
	}

	for _, test := range tests {
		got, _ := findSecret(test.input)
		if got != test.want {
			t.Errorf("findSecret(%q) = %q, wanted %q", test.input, got, test.want)
		}
	}
}

func TestFindSecretErr(t *testing.T) {
	var tests = []struct {
		input            string
		want             string
		envGithubSecrets string
	}{
		{SECRET_ID, "", GITHUB_SECRETS},
		{SECRET_ID, "entry 'github_token' not found", "{}"},
		{SECRET_ID, "environment variable GITHUB_SECRETS is empty", ""},
		{SECRET_ID, "unexpected end of JSON input", "{"},
	}

	for _, test := range tests {
		os.Setenv(INPUT_GITHUB_SECRETS, test.envGithubSecrets)

		_, got := findSecret(test.input)
		if got != nil && got.Error() != test.want {
			t.Errorf("findSecret(%q) = %q, wanted %q", test.input, got, test.want)
		}
	}
}
