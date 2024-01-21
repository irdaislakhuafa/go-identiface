package identiface

import (
	"github.com/Kagami/go-face"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

// A human face data descriptor
type Data[ID any] struct {
	ID   ID
	Face face.Face
}

type Identiface[ID any] interface {
	// This method will override existing datasets. Call `LoadDatasets()` to load new datasets after overriding.
	SetDatasets(datasets []Data[ID]) Identiface[ID]

	// Get used datasets from `Indetiface`
	GetDatasets() []Data[ID]

	// Clear used datasets. The datasets will be empty after using this method
	ClearDatasets()

	// Load datasets from `Identiface`. By default not loaded for performance reasons
	LoadDatasets()

	// Adding single datasets for face recognization from bytes dataset
	AddSingleDatasetFromBytes(id ID, datasetBytes []byte) error

	// Classify or Identify single datasets. Will return `error` if dataset from parameter is not recognized based on used datasets
	ClassifySingleFromBytes(datasetBytes []byte) (Data[ID], error)

	// Recognize single face from `datasetBytes`. Return `error` if face more than one or `datasetBytes` does not contain face
	RecognizeSingleFromBytes(datasetBytes []byte) (face.Face, error)

	// Set custom recognizer with `https://github.com/Kagami/go-face.git`. By default using `face.NewRecognizer()`
	SetRecognizer(rec *face.Recognizer) Identiface[ID]

	// Get used recognizer of `Identiface`
	GetRecognizer() *face.Recognizer

	// Set custom tolerance. By default using `0.4`. Smaller value, the stricter the facial pattern recognition
	SetTolerance(tolerance float32) Identiface[ID]

	// Get used tolerance value
	GetTolerance() float32

	// Set value to `true` if want use grey image. By default `false`
	SetGrey(isGrey bool) Identiface[ID]

	// Get status grey image config
	IsGrey() bool

	// Set value to `true` if want use CNN image. By default `false`
	SetCNN(isCNN bool) Identiface[ID]

	// Get status CNN image config
	IsCNN() bool

	// Close recognizer of `Identiface`
	Close()
}

type identiface[ID any] struct {
	tolerance float32
	rec       *face.Recognizer
	isGrey    bool
	isCNN     bool
	datasets  []Data[ID]
}

// Initialize Identiface
func Init[ID any](modelsDir string) (Identiface[ID], error) {
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeIdentiface, "failed to create recognizer, %v", err)
	}

	i := &identiface[ID]{
		tolerance: 0.4,
		rec:       rec,
		isGrey:    false,
		isCNN:     false,
		datasets:  []Data[ID]{},
	}

	return i, nil
}

// This method will override existing datasets. Call `LoadDatasets()` to load new datasets after overriding.
func (i *identiface[ID]) SetDatasets(datasets []Data[ID]) Identiface[ID] {
	i.datasets = datasets
	return i
}

// Get used datasets from `Indetiface`
func (i *identiface[ID]) GetDatasets() []Data[ID] {
	return i.datasets
}

// Clear used datasets. The datasets will be empty after using this method
func (i *identiface[ID]) ClearDatasets() {
	clear(i.datasets)
	i.rec.SetSamples(nil, nil)
}

// Load datasets from `Identiface`. By default not loaded for performance reasons
func (i *identiface[ID]) LoadDatasets() {
	listSample := []face.Descriptor{}
	listID := []int32{}

	for i, v := range i.datasets {
		listSample = append(listSample, v.Face.Descriptor)
		listID = append(listID, int32(i))
	}

	i.rec.SetSamples(listSample, listID)
}

// Adding single datasets for face recognization from bytes dataset
func (i *identiface[ID]) AddSingleDatasetFromBytes(id ID, datasetBytes []byte) error {
	var (
		err   error
		faces []face.Face
	)

	if i.isCNN {
		faces, err = i.rec.RecognizeCNN(datasetBytes)
	} else if i.isGrey {
		// TODO: handle grey image recognization here
	} else {
		faces, err = i.rec.Recognize(datasetBytes)
	}

	if err != nil {
		return errors.NewWithCode(codes.CodeIdentifaceFaceNotRecognized, "cannot recognize face from bytes dataset, %v", err)
	}

	if lf := len(faces); lf <= 0 {
		return errors.NewWithCode(codes.CodeIdentifaceNoFaceDetected, "no face detected from bytes dataset")
	} else if lf > 1 {
		return errors.NewWithCode(codes.CodeIdentifaceMultipleFaceDetected, "there is more than one face from bytes dataset")
	}

	data := Data[ID]{
		ID:   id,
		Face: faces[0],
	}

	i.datasets = append(i.datasets, data)
	return nil
}

// Classify or Identify single datasets. Will return `error` if dataset from parameter is not recognized based on used datasets
func (i *identiface[ID]) ClassifySingleFromBytes(datasetBytes []byte) (Data[ID], error) {
	faceID, err := i.RecognizeSingleFromBytes(datasetBytes)
	if err != nil {
		return Data[ID]{}, err
	}

	personID := i.rec.ClassifyThreshold(faceID.Descriptor, i.tolerance)
	if personID < 0 {
		return Data[ID]{}, errors.NewWithCode(codes.CodeIdentifaceFaceNotRegistered, "can't classify bytes dataset, may dataset not registered or existing datasets not loaded")
	}

	data := i.datasets[personID]
	return data, nil
}

// Recognize single face from `datasetBytes`. Return `error` if face more than one or `datasetBytes` does not contain face
func (i *identiface[ID]) RecognizeSingleFromBytes(datasetBytes []byte) (face.Face, error) {
	var (
		err    error
		faceID *face.Face
	)

	if i.isCNN {
		faceID, err = i.rec.RecognizeSingleCNN(datasetBytes)
	} else if i.isGrey {
		// TODO: handle grey image recognization here
	} else {
		faceID, err = i.rec.RecognizeSingle(datasetBytes)
	}

	if err != nil {
		return face.Face{}, errors.NewWithCode(codes.CodeIdentifaceFaceNotRecognized, "cannot recognize single face, %v", err)
	}

	if faceID == nil {
		return face.Face{}, errors.NewWithCode(codes.CodeIdentifaceNoFaceDetected, "the bytes dataset does not contain a single face")
	}

	return *faceID, nil
}

// Set custom recognizer with `https://github.com/Kagami/go-face.git`. By default using `face.NewRecognizer()`
func (i *identiface[ID]) SetRecognizer(rec *face.Recognizer) Identiface[ID] {
	i.rec = rec
	return i
}

// Get used recognizer of `Identiface`
func (i *identiface[ID]) GetRecognizer() *face.Recognizer {
	return i.rec
}

// Set custom tolerance. By default using `0.4`
func (i *identiface[ID]) SetTolerance(tolerance float32) Identiface[ID] {
	i.tolerance = tolerance
	return i
}

// Get used tolerance value
func (i *identiface[ID]) GetTolerance() float32 {
	return i.tolerance
}

// Set value to `true` if want use grey image
func (i *identiface[ID]) SetGrey(isGrey bool) Identiface[ID] {
	i.isGrey = isGrey
	return i
}

// Get status grey image config
func (i *identiface[ID]) IsGrey() bool {
	return i.isGrey
}

// Set value to `true` if want use CNN image
func (i *identiface[ID]) SetCNN(isCNN bool) Identiface[ID] {
	i.isCNN = isCNN
	return i
}

// Get status CNN image config
func (i *identiface[ID]) IsCNN() bool {
	return i.isCNN
}

// Close recognizer of `Identiface`
func (i *identiface[ID]) Close() {
	i.rec.Close()
}
