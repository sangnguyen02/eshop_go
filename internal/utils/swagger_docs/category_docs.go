package docs

// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.Category true "Category data"
// @Success 201 {object} model.Category
// @Failure 400 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/categories [post]
func CreateCategoryDoc() {}

// @Summary Get category by ID
// @Description Get details of a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Category
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/categories/{id} [get]
func GetCategoryByIDDoc() {}

// @Summary Get all categories
// @Description Get all categories with pagination and search
// @Tags categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Param name query string false "Search by name"
// @Success 200 {object} model.SuccessResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/categories [get]
func GetAllCategoriesDoc() {}

// @Summary Update an existing category
// @Description Update a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.Category true "Category data"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/categories/{id} [put]
func UpdateCategoryDoc() {}

// @Summary Delete a category
// @Description Delete a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/categories/{id} [delete]
func DeleteCategoryDoc() {}
