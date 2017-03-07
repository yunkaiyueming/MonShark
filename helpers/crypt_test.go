package helpers

import (
	"testing"
)

func TestSha1(t *testing.T) {
	originStr := []string{"Rret3s1As", "Ns2Md5s", "fni02Mwg5"}
	expectedStr := []string{"d3c2ce42ab46c9b8747e43efa32bf1d23ddc6c45", "bfec75bf8c88b27e30cebc41936ee21093f9b21d", "21b753529d58b3fc8198c518904dde1140920936"}
	for i, str := range originStr {
		if expectedStr[i] != Sha1(str) {
			t.Error("not equal")
		}
	}
}

func TestMd5(t *testing.T) {
	originStr := []string{"Rret3s1As", "Ns2Md5s", "fni02Mwg5"}
	expectedStr := []string{"a5cb5b8fd34e45ade62380f00f66acc0", "59e710261320ad259526329577dbd916", "6fffabe89bd3b7dd821619a9c269b669"}
	for i, str := range originStr {
		if expectedStr[i] != Md5(str) {
			t.Error("not equal")
		}
	}
}
