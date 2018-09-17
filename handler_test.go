package main

import "testing"

func TestCutmName(t *testing.T) {
	s := "ops-test-4033419859-b6w75"

	name, err := cutName(s)
	if err != nil {
		t.Error("transform err", err)
	}
	if name != "ops-test" {
		t.Error("transform failed")
	}
}

func TestIsNamespaceOK(t *testing.T) {
	tests := []struct {
		ns  string
		nss string
		ok  bool
	}{
		{
			ns:  "aa",
			nss: "",
			ok:  true,
		},
		{
			ns:  "pre",
			nss: "pre,default",
			ok:  true,
		},
	}

	for _, v := range tests {
		namespaces := getNS(v.nss)
		if got := IsNamespaceOK(v.ns, namespaces); got != v.ok {
			t.Errorf("%v err, got %v, want %v\n", v.ns, got, v.ok)
		}
	}
}
