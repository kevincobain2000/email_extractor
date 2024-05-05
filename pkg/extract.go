package pkg

type Extract struct {
}

func NewExtract() *Extract {
	return &Extract{}
}

func (r *Extract) GetTypes(req *ExtractorRequest) error {
	return nil
}
