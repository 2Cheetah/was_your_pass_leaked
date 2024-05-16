package helper

import "testing"

func TestSha1HashFromString(t *testing.T) {
	testCases := []struct {
		password string
		expected string
	}{
		{"password", "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8"},
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"p@ssw0rd! ", "113b8d0b7b8a30eb5380e120b696dc58c3f4852d"},
	}

	for i, testCase := range testCases {
		t.Logf("Running test case %d", i+1)
		result := Sha1HashFromString(testCase.password)
		if result != testCase.expected {
			t.Errorf("sha-1 sum didn't match expected for password %s",
				testCase.password)
		}
	}
}
