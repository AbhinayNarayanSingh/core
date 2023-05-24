package locals

const (
	BadRequest          = "Bad Request."
	InternalServerError = "Ops! something has gone wrong on the website's server, engineers are notified."

	AccountCreated = ", Congratulations! Your account has been successfully created."

	InvalidPassword           = "Invalid password, remember that passwords are case-sensitive."
	EmailNotRegistered        = "Email address isn't associated with this account."
	EmailAssociateWithAccount = "Email is already associated with an account."
	PhoneNotRegistered        = "Phone number isn't associated with this account."
	PhoneAssociateWithAccount = "Phone number is already associated with an account."
	AccountNotActivated       = "Account activation required to continue."
	AccountActivated          = "Your account has been activated successfully, proceed with signing."
	OTPSendOnPhone            = "A OTP (One Time Password) has been sent on your phone number."
	OTPSendOnEmail            = "A OTP (One Time Password) has been sent on your email address."
	OTPSendOnPhoneForReset    = "A security code has been sent on your phone number to reset password."
	OTPSendOnEmailForReset    = "A security code has been sent on your email address to reset password."
	OTPNotGenerated           = "Repeated failures user is advised to generate new OTP."
	OTPInvalid                = "Invalid OTP."
	UserNotFound              = "User not found."
)
