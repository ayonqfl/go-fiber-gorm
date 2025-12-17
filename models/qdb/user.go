package models

import "time"

type User struct {
    ID                    uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    CreatedAt             time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
    Username              string    `json:"username"`
    Password              string    `json:"password"`
    Email                 string    `json:"email"`
    Photo                 string    `json:"photo"`
    UsersRoles            string    `json:"users_roles"`
    AccType               string    `json:"acc_type"`
    UserID                string    `json:"user_id"`
    Branch                string    `json:"branch"`
    Name                  string    `json:"name"`
    EmailStatus           string    `json:"email_status"`
    PhoneStatus           string    `json:"phone_status"`
    Phone                 string    `json:"phone"`
    AccountStatus         string    `json:"account_status"`
    Exchange              string    `json:"exchange"`
    DealerGroupID         string    `json:"dealer_group_id"`
    MaxLogin              int       `json:"max_login"`
    LoggedIn              int       `json:"logged_in"`
    LastLogin             string    `json:"last_login"`
    LoginIP               string    `json:"login_ip"`
    Premium               bool      `json:"premium"`
    PremiumStartDate      string    `json:"premium_start_date"`
    PremiumEndDate        string    `json:"premium_end_date"`
    MaxLoginMobile        int       `json:"max_login_mobile"`
    LoggedInMobile        int       `json:"logged_in_mobile"`
    TotalMaxLogin         int       `json:"total_max_login"`
    TotalLoggedIn         int       `json:"total_logged_in"`
    MarginAllowed         bool      `json:"margin_allowed" gorm:"default:false"`
    AcStatusUpdateBy      string    `json:"ac_status_update_by"`
    AcStatusUpdateTime    string    `json:"ac_status_update_time"`
    ParkingEnabled        bool      `json:"parking_enabled" gorm:"default:false"`
    IsBulkOrder           bool      `json:"is_bulk_order" gorm:"default:false"`
    FirstLogin            bool      `json:"first_login" gorm:"default:false"`
    LoginOTP              bool      `json:"login_otp" gorm:"default:false"`
}

// TableName overrides the table name
func (User) TableName() string {
    return "users"
}