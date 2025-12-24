package utils

import "github.com/gofiber/fiber/v2"

// APIResponse represents the standard API response structure
type APIResponse struct {
	Message     string      `json:"message"`
	Errors      interface{} `json:"errors,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Code        int         `json:"code"`
	RedirectURL string      `json:"redirect_url,omitempty"`
}

// ResponseOptions holds optional parameters for SendResponse
type ResponseOptions struct {
	Message  string
	Data     interface{}
	Errors   interface{}
	Redirect string
}

// SendResponse sends a standardized JSON response
func SendResponse(c *fiber.Ctx, code int, opts ResponseOptions) error {

	response := APIResponse{
		Message:     opts.Message,
		Errors:      opts.Errors,
		Data:        opts.Data,
		Code:        code,
		RedirectURL: opts.Redirect,
	}

	return c.Status(code).JSON(response)
}