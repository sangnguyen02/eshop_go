package docs

// @Summary Upload a single image file
// @Description Upload a single image (jpg, jpeg, or png) to the server
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "Image file to upload"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/v1/upload [post]
func UploadFileSingleDoc() {}

// @Summary Upload multiple image files
// @Description Upload multiple image files (jpg, jpeg, or png) to the server
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "Image files to upload"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/v1/upload/multiple [post]
func UploadFileMultipleDoc() {}
