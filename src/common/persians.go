package common

import (
	"regexp"

	"github.com/Mehrshadshaabani/Web-API-Clean/config"
	"github.com/Mehrshadshaabani/Web-API-Clean/pkg/logging"
)

const iranianMobileNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianMobileNumberValidate(mobileNumber string) bool {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	res, err := regexp.MatchString(iranianMobileNumberPattern, mobileNumber)
	if err != nil {
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return res
}
