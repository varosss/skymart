package fakes

type FakePasswordVerifier struct {
	Ok bool
}

func (f *FakePasswordVerifier) Compare(hash, password string) bool {
	return f.Ok
}
