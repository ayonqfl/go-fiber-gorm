package routes

import (
	"time"

	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	ID                 uint      `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	Photo              string    `json:"photo"`
	UsersRoles         string    `json:"users_roles"`
	AccType            string    `json:"acc_type"`
	UserID             string    `json:"user_id"`
	Branch             string    `json:"branch"`
	Name               string    `json:"name"`
	EmailStatus        string    `json:"email_status"`
	PhoneStatus        string    `json:"phone_status"`
	Phone              string    `json:"phone"`
	AccountStatus      string    `json:"account_status"`
	Exchange           string    `json:"exchange"`
	DealerGroupID      string    `json:"dealer_group_id"`
	MaxLogin           int       `json:"max_login"`
	LoggedIn           int       `json:"logged_in"`
	LastLogin          string    `json:"last_login"`
	LoginIP            string    `json:"login_ip"`
	Premium            bool      `json:"premium"`
	PremiumStartDate   string    `json:"premium_start_date"`
	PremiumEndDate     string    `json:"premium_end_date"`
	MaxLoginMobile     int       `json:"max_login_mobile"`
	LoggedInMobile     int       `json:"logged_in_mobile"`
	TotalMaxLogin      int       `json:"total_max_login"`
	TotalLoggedIn      int       `json:"total_logged_in"`
	MarginAllowed      bool      `json:"margin_allowed"`
	AcStatusUpdateBy   string    `json:"ac_status_update_by"`
	AcStatusUpdateTime string    `json:"ac_status_update_time"`
	ParkingEnabled     bool      `json:"parking_enabled"`
	IsBulkOrder        bool      `json:"is_bulk_order"`
	FirstLogin         bool      `json:"first_login"`
	LoginOTP           bool      `json:"login_otp"`
}

func CreateResponseUser(user models.User) UserSerializer {
	return UserSerializer{
		ID:                 user.ID,
		CreatedAt:          user.CreatedAt,
		Username:           user.Username,
		Email:              user.Email,
		Photo:              user.Photo,
		UsersRoles:         user.UsersRoles,
		AccType:            user.AccType,
		UserID:             user.UserID,
		Branch:             user.Branch,
		Name:               user.Name,
		EmailStatus:        user.EmailStatus,
		PhoneStatus:        user.PhoneStatus,
		Phone:              user.Phone,
		AccountStatus:      user.AccountStatus,
		Exchange:           user.Exchange,
		DealerGroupID:      user.DealerGroupID,
		MaxLogin:           user.MaxLogin,
		LoggedIn:           user.LoggedIn,
		LastLogin:          user.LastLogin,
		LoginIP:            user.LoginIP,
		Premium:            user.Premium,
		PremiumStartDate:   user.PremiumStartDate,
		PremiumEndDate:     user.PremiumEndDate,
		MaxLoginMobile:     user.MaxLoginMobile,
		LoggedInMobile:     user.LoggedInMobile,
		TotalMaxLogin:      user.TotalMaxLogin,
		TotalLoggedIn:      user.TotalLoggedIn,
		MarginAllowed:      user.MarginAllowed,
		AcStatusUpdateBy:   user.AcStatusUpdateBy,
		AcStatusUpdateTime: user.AcStatusUpdateTime,
		ParkingEnabled:     user.ParkingEnabled,
		IsBulkOrder:        user.IsBulkOrder,
		FirstLogin:         user.FirstLogin,
		LoginOTP:           user.LoginOTP,
	}
}

func UserHandlers(route fiber.Router) {
	// Define users list API function
	route.Get("/list", func(C *fiber.Ctx) error {
		users := []models.User{}
		responseUsers := []UserSerializer{}

		database.Database.Db.Order("id DESC").Limit(100).Find(&users)
		for _, user := range users {
			responseUser := CreateResponseUser(user)
			responseUsers = append(responseUsers, responseUser)
		}

		return C.Status(200).JSON(fiber.Map{
			"message": "Success",
			"data":    responseUsers,
		})
	})

	// Define users create API function
	route.Post("/create", func(C *fiber.Ctx) error {
		var user models.User

		if err := C.BodyParser(&user); err != nil {
			return C.Status(400).JSON(err.Error())
		}

		database.Database.Db.Create(&user)
		responseUser := CreateResponseUser(user)
		return C.Status(200).JSON(fiber.Map{
			"message": "Success",
			"data":    responseUser,
		})
	})
}