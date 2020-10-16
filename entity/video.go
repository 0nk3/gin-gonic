package entity

// Person ..
type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=45"` // greater than equal, less than equal
	Email     string `json:"email" validate:"required, email"`
}

// Video describes video data
type Video struct {
	Title       string `json:"title" binding:"min=2,max=11"` // validate
	Description string `json:"description" binding:"max=15"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
