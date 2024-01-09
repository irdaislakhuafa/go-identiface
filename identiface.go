package identiface

import "github.com/Kagami/go-face"

// A human face data descriptor
type Data[ID any] struct {
	ID   ID
	Data face.Face
}

type Identiface[ID any] interface {
	SetDatasets(datasets []Data[ID]) Identiface[ID]
	GetDatasets() []Data[ID]
	ClearDatasets()
	LoadDatasets()
	AddSingleDatasetFromBytes(id ID, dataset []byte) error
	ClassifySingleFromBytes(dataset []byte) (Data[ID], error)
	RecognizeSingleFromBytes(dataset []byte) (face.Face, error)
	SetTolerance(tolerance float32) Identiface[ID]
	GetTolerance() float32
	SetGrey(grey bool) Identiface[ID]
	IsGrey() bool
	SetCNN(cnn bool) Identiface[ID]
	IsCNN() bool
}
