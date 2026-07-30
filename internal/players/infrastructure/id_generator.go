package infrastructure

type UUIDGenerator struct{}

func (UUIDGenerator) NewId() string {
	return "asdf"
}
