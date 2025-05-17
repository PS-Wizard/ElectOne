package auth

import (
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type IncomingRequest struct {
	VoterCardID   string `json:"voter_card_id"`
	CitizenshipID string `json:"citizenship_id"`
	TOTP          string `json:"totp"`
	Password      string `json:"password"`
}

func HandleDeleteUser(c *fiber.Ctx) error {
	var request IncomingRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request Body")
	}
	user, err := operations.GetUserByCitizenAndVoterID(request.CitizenshipID, request.VoterCardID)
	if err != nil {
		return fmt.Errorf("Encountered Error: %+v", err)
	}

	if err := validateTOTP(request.TOTP, user.TotpSecret); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return fmt.Errorf("Incorrect Password")
	}

	if err = operations.DeleteUser(user.UserID); err != nil {
		return fmt.Errorf("Failed To Delete User: %+v", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Your Account Has Been Successfully Deleted",
	})
}

func HandleForgotPassword(c *fiber.Ctx) error {
	var request IncomingRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request Body")
	}
	user, err := operations.GetUserByCitizenAndVoterID(request.CitizenshipID, request.VoterCardID)
	if err != nil {
		return fmt.Errorf("Encountered Error: %+v", err)
	}

	if user.TotpSecret == "" {
		return fmt.Errorf("The user is yet to setup TOTP: %+v", err)
	}

	if err := validateTOTP(request.TOTP, user.TotpSecret); err != nil {
		return err
	}

	if err = operations.UpdateUserPassword(user.UserID, request.Password); err != nil {
		return fmt.Errorf("Failed To Update User Password: %+v", err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Updated Password",
	})
}

func validateTOTP(incomingTOTP string, secret string) error {

	if !totp.Validate(incomingTOTP, secret) {
		return fmt.Errorf("Invalid TOTP")
	}

	return nil
}
