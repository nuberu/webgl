package types

type ActiveInfo struct {
	name     string
	size     int
	infoType GLEnum
}

func NewActiveInfo(name string, size int, infoType GLEnum) *ActiveInfo {
	return &ActiveInfo{
		name:     name,
		size:     size,
		infoType: infoType,
	}
}

func (ai *ActiveInfo) GetName() string {
	return ai.name
}

func (ai *ActiveInfo) GetSize() int {
	return ai.size
}

func (ai *ActiveInfo) GetType() GLEnum {
	return ai.infoType
}
