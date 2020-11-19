package handlers

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func Login(c *fiber.Ctx) error {
// 	p := new(Auth)
// 	argon2ID := utils.NewArgon2ID()

// 	if err := c.BodyParser(p); err != nil {
// 		return err
// 	}

// 	// hash, _ := utils.HashPassword(p.Password)

// 	collection := databases.Db.Collection("user")

// 	var result models.User

// 	err := collection.FindOne(c.Context(), bson.D{{Key: "email", Value: p.Email}}).Decode(&result)

// 	if err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}

// 	hashedPassword := result.Password

// 	ok1, err := argon2ID.Verify(p.Password, hashedPassword)

// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	if ok1 == false {
// 		return c.Status(500).SendString("Error pw")
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["sub"] = result.ID
// 	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // a week

// 	s, err := token.SignedString([]byte(middlewares.JwtSecret))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"token": s,
// 		"user": struct {
// 			Id    primitive.ObjectID `json:"id"`
// 			Email string             `json:"email"`
// 		}{
// 			Id:    result.ID,
// 			Email: result.Email,
// 		},
// 	})
// }
