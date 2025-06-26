package docs

// @Summary Create a new brand
// @Description Create a new brand
// @Tags brands
// @Accept json
// @Produce json
// @Param brand body model.Brand true "Brand data"
// @Success 201 {object} model.Brand
// @Failure 400 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/brands [post]
func CreateBrandDoc() {}

// @Summary Get brand by ID
// @Description Get details of a brand by ID
// @Tags brands
// @Accept json
// @Produce json
// @Param id path int true "Brand ID"
// @Success 200 {object} model.Brand
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/brands/{id} [get]
func GetBrandByIDDoc() {}

// @Summary Get all brands
// @Description Get all brands with pagination and search
// @Tags brands
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Param name query string false "Search by name"
// @Success 200 {object} model.SuccessResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/brands [get]
func GetAllBrandsDoc() {}

// @Summary Update an existing brand
// @Description Update a brand by ID
// @Tags brands
// @Accept json
// @Produce json
// @Param id path int true "Brand ID"
// @Param brand body model.Brand true "Brand data"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/brands/{id} [put]
func UpdateBrandDoc() {}

// @Summary Delete a brand
// @Description Delete a brand by ID
// @Tags brands
// @Accept json
// @Produce json
// @Param id path int true "Brand ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/brands/{id} [delete]
func DeleteBrandDoc() {}
