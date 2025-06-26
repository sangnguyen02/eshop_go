package docs

// @Summary Search products by name
// @Description Search products with pagination
// @Tags products
// @Accept json
// @Produce json
// @Param name query string false "Product name"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} model.SuccessResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/products/search [get]
func SearchProductsDoc() {}

// @Summary Get product by ID
// @Description Get details of a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/products/{id} [get]
func GetProductByIDDoc() {}

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product data"
// @Success 201 {object} model.Product
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/products [post]
func CreateProductDoc() {}

// @Summary Update an existing product
// @Description Update an existing product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Product true "Product data"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/products/{id} [put]
func UpdateProductDoc() {}

// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/products/{id} [delete]
func DeleteProductDoc() {}
