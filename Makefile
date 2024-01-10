GO_ASSET_DIR := ./assets

get-models:
	@ git clone https://github.com/irdaislakhuafa/go-identiface-testdata ${GO_ASSET_DIR}

get-all-models:
	@ git clone https://github.com/davisking/dlib-models.git ${GO_ASSET_DIR}/models

clean-models:
	@ rm -rfv ./assets/models

clean-assets:
	@ rm -rfv ./assets