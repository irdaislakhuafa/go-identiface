package identiface

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_Identiface(t *testing.T) {

	type Mode int
	const (
		MODE_CLASSIFY_SINGLE = Mode(iota + 1)
		MODE_RECOGNIZE_SINGLE
		MODE_ADD_DATASET_SINGLE
	)

	const (
		assetsDir = "./assets"
		imagesDir = "images"
		modelsDir = "models"
	)

	type image struct {
		id       string
		name     string
		fileName string
	}

	type params struct {
		targetImage string
		images      []image
		assetsDir   string
		imagesDir   string
		modelsDir   string
	}

	type want struct {
		id string
	}

	type wantErr struct {
		code codes.Code
	}

	type test struct {
		mode          Mode
		beforeFunc    func(i Identiface[string], test test)
		name          string
		params        params
		want          want
		wantErr       wantErr
		isWantInitErr bool
		isWantTestErr bool
	}

	tests := []test{
		{
			mode:       MODE_ADD_DATASET_SINGLE,
			name:       "success add single datasets from bytes",
			beforeFunc: nil,
			params: params{
				targetImage: "", // empty because we don't neet target image for add datasets, just need images
				images: []image{
					{id: "x1", name: "this is tzuyu", fileName: "tzuyu.jpg"},
					{id: "x2", name: "this is jimin", fileName: "jimin.jpg"},
				},
				imagesDir: imagesDir, assetsDir: assetsDir, modelsDir: modelsDir,
			},
			want:          want{id: ""},
			isWantInitErr: false,
			isWantTestErr: false,
			wantErr:       wantErr{code: codes.NoCode},
		},
		{
			mode:       MODE_ADD_DATASET_SINGLE,
			name:       "failed add single datasets from bytes because no face detected",
			beforeFunc: nil,
			params: params{
				targetImage: "", // empty because we don't neet target image for add datasets, just need images
				images: []image{
					{id: "x0", name: "this is wall without face", fileName: "wall-no-face.jpeg"},
				},
				imagesDir: imagesDir, assetsDir: assetsDir, modelsDir: modelsDir,
			},
			want:          want{id: ""},
			isWantInitErr: false,
			isWantTestErr: true,
			wantErr:       wantErr{code: codes.CodeIdentiface},
		},
		{
			mode: MODE_CLASSIFY_SINGLE,
			name: "success classify single in bytes",
			beforeFunc: func(i Identiface[string], test test) {
				i.SetTolerance(0.2)

				for _, img := range test.params.images {
					fileBytes, err := os.ReadFile(filepath.Join(test.params.assetsDir, test.params.imagesDir, img.fileName))
					if err != nil {
						panic(err)
					}

					if err := i.AddSingleDatasetFromBytes(img.id, fileBytes); err != nil {
						panic(err)
					}
				}

				i.LoadDatasets()
			},
			params: params{
				targetImage: "tzuyu2.jpg",
				images: []image{
					{id: "x1", name: "this is tzuyu", fileName: "tzuyu.jpg"},
					{id: "x2", name: "this is jimin", fileName: "jimin.jpg"},
				},
				imagesDir: imagesDir, assetsDir: assetsDir, modelsDir: modelsDir,
			},
			want:          want{id: "x1"},
			isWantInitErr: false,
			isWantTestErr: false,
		},
		{
			mode:       MODE_RECOGNIZE_SINGLE,
			name:       "success recognize single face",
			beforeFunc: nil,
			params: params{
				targetImage: "tzuyu.jpg",
				imagesDir:   imagesDir, assetsDir: assetsDir, modelsDir: modelsDir,
			},
			want:          want{},
			isWantInitErr: false,
			isWantTestErr: false,
		},
		{
			mode:       MODE_RECOGNIZE_SINGLE,
			name:       "failed recognize single face because got multiple face",
			beforeFunc: nil,
			params: params{
				targetImage: "pristin.jpg",
				imagesDir:   imagesDir, assetsDir: assetsDir, modelsDir: modelsDir,
			},
			want:          want{},
			isWantInitErr: false,
			isWantTestErr: true,
			wantErr:       wantErr{code: codes.CodeIdentiface},
		},
		{
			mode:       MODE_RECOGNIZE_SINGLE,
			name:       "failed recognize single face because not face detected",
			beforeFunc: nil,
			params: params{
				targetImage: "wall-no-face.jpeg",
				imagesDir:   imagesDir, assetsDir: assetsDir, modelsDir: modelsDir,
			},
			want:          want{},
			isWantInitErr: false,
			isWantTestErr: true,
			wantErr:       wantErr{code: codes.CodeIdentiface},
		},
		{
			name: "failed classify and datasets loaded",
			mode: MODE_CLASSIFY_SINGLE,
			beforeFunc: func(i Identiface[string], test test) {
				i.SetTolerance(0.2)

				for _, img := range test.params.images {
					fileBytes, err := os.ReadFile(filepath.Join(test.params.assetsDir, test.params.imagesDir, img.fileName))
					if err != nil {
						panic(err)
					}

					if err := i.AddSingleDatasetFromBytes(img.id, fileBytes); err != nil {
						panic(err)
					}
				}

				i.LoadDatasets()
			},
			params: params{
				targetImage: "rena.jpg",
				images: []image{
					{id: "x1", name: "this is tzuyu", fileName: "tzuyu.jpg"},
					{id: "x2", name: "this is jimin", fileName: "jimin.jpg"},
				},
				imagesDir: imagesDir, assetsDir: assetsDir, modelsDir: modelsDir,
			},
			want:          want{id: ""},
			isWantInitErr: false,
			isWantTestErr: true,
			wantErr:       wantErr{code: codes.CodeIdentiface},
		},
		{
			name: "failed classify and datasets not loaded",
			mode: MODE_CLASSIFY_SINGLE,
			beforeFunc: func(i Identiface[string], test test) {
				i.SetTolerance(0.2)

				for _, img := range test.params.images {
					fileBytes, err := os.ReadFile(filepath.Join(test.params.assetsDir, test.params.imagesDir, img.fileName))
					if err != nil {
						panic(err)
					}

					if err := i.AddSingleDatasetFromBytes(img.id, fileBytes); err != nil {
						panic(err)
					}
				}

				// i.LoadDatasets()
			},
			params: params{
				targetImage: "rena.jpg",
				images: []image{
					{id: "x1", name: "this is tzuyu", fileName: "tzuyu.jpg"},
					{id: "x2", name: "this is jimin", fileName: "jimin.jpg"},
				},
				imagesDir: imagesDir, assetsDir: assetsDir, modelsDir: modelsDir,
			},
			want:          want{id: ""},
			isWantInitErr: false,
			isWantTestErr: true,
			wantErr:       wantErr{code: codes.CodeIdentiface},
		},
	}

	f := files.GetCurrentMethodName()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", f, tt.name), func(t *testing.T) {
			iFace, err := Init[string](filepath.Join(tt.params.assetsDir, tt.params.modelsDir))

			if tt.beforeFunc != nil {
				tt.beforeFunc(iFace, tt)
			}

			if tt.isWantInitErr {
				if err == nil {
					t.Fatalf("want init err is %#v but got err %#v", tt.isWantInitErr, err.Error())
				}

				if code := errors.GetCode(err); code != tt.wantErr.code {
					t.Fatalf("want init err code is %#v but got err code %#v", tt.wantErr.code, code)
				} else {
					t.Logf("want init err code %#v is equals with err code %#v", tt.wantErr.code, code)
				}
			} else {
				if err != nil {
					t.Fatalf("want init err is %#v but got err with msg %#v", tt.isWantInitErr, err.Error())
				}
			}

			switch tt.mode {
			case MODE_ADD_DATASET_SINGLE:
				{
					for _, img := range tt.params.images {
						filePath := filepath.Join(tt.params.assetsDir, tt.params.imagesDir, img.fileName)
						fileBytes, err := os.ReadFile(filePath)
						if err != nil {
							t.Fatalf("cannot read file %#v, %#v", filePath, err.Error())
						}

						if err := iFace.AddSingleDatasetFromBytes(img.id, fileBytes); tt.isWantTestErr {
							if err == nil {
								t.Fatalf("want test err is %#v but got err %#v for img id %#v", tt.isWantTestErr, err, img.id)
							} else if code := errors.GetCode(err); code != tt.wantErr.code {
								t.Fatalf("want test err code is %#v but got err code %#v", tt.wantErr.code, code)
							} else {
								t.Logf("want test err code is %#v equals with result err code %#v", tt.wantErr.code, code)
							}
						} else {
							if err != nil {
								t.Fatalf("want test err is %#v but got err with msg %#v", tt.isWantTestErr, err.Error())
							}
						}
					}
				}
			case MODE_CLASSIFY_SINGLE:
				{
					filePath := filepath.Join(tt.params.assetsDir, tt.params.imagesDir, tt.params.targetImage)
					targetImageBytes, err := os.ReadFile(filePath)
					if err != nil {
						t.Fatalf("cannot read file %#v, %#v", filePath, err.Error())
					}

					data, err := iFace.ClassifySingleFromBytes(targetImageBytes)
					if tt.isWantTestErr {
						if err == nil {
							t.Fatalf("want test err is %#v but got err %#v with result %+v", tt.isWantTestErr, err, data)
						} else if code := errors.GetCode(err); code != tt.wantErr.code {
							t.Fatalf("want test err code is %#v but got err code %#v", tt.wantErr.code, code)
						} else {
							t.Logf("want test err code is %#v equals with result err code %#v", tt.wantErr.code, code)
						}
					} else {
						if err != nil {
							t.Fatalf("want test err is %#v but got err with msg %#v", tt.isWantTestErr, err.Error())
						}
					}

					if data.ID != tt.want.id {
						t.Fatalf("want id %#v but got %#v", tt.want.id, data.ID)
					} else {
						t.Logf("want id %#v is equals with result id %#v for image %#v", tt.want.id, data.ID, tt.params.targetImage)
					}
				}
			case MODE_RECOGNIZE_SINGLE:
				{
					filePath := filepath.Join(tt.params.assetsDir, tt.params.imagesDir, tt.params.targetImage)
					targetImageBytes, err := os.ReadFile(filePath)
					if err != nil {
						t.Fatalf("cannot read file %#v, %#v", filePath, err.Error())
					}

					_, err = iFace.RecognizeSingleFromBytes(targetImageBytes)
					if tt.isWantTestErr {
						if err == nil {
							t.Fatalf("want test err is %#v but got err %#v", tt.isWantTestErr, err)
						} else if code := errors.GetCode(err); code != tt.wantErr.code {
							t.Fatalf("want test err code is %#v but got err code %#v", tt.wantErr.code, code)
						} else {
							t.Logf("want test err code is %#v equals with result err code %#v", tt.wantErr.code, code)
						}
					} else {
						if err != nil {
							t.Fatalf("want test err is %#v but got err with msg %#v", tt.isWantTestErr, err.Error())
						}
					}
				}
			}
		})
		fmt.Println("")
	}
}
