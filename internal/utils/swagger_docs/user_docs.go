package docs

// @Summary Sign up a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body handler.SignUpForm true "User data"
// @Success 201 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/signup [post]
func SignUpDoc() {}

// @Summary Sign in a user
// @Description Login a user and return a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param user body handler.LoginForm true "User credentials"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/signin [post]
func SignInDoc() {}

// @Summary Search users by name
// @Description Search users with pagination
// @Tags users
// @Accept json
// @Produce json
// @Param name query string true "User name"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/search [get]
func SearchUsersDoc() {}

// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/{id} [get]
func GetUserByIDDoc() {}

// @Summary Update an existing user
// @Description Update a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body handler.UpdateForm true "User data"
// @Success 200 {object} model.User
// @Failure 400 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/{id} [put]
func UpdateUserDoc() {}

// @Summary Update user password
// @Description Update the password of a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param password body handler.PasswordUpdateForm true "New password"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/pass/{id} [put]
func UpdatePasswordDoc() {}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/{id} [delete]
func DeleteUserDoc() {}
