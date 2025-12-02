package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"github.com/salmantaghooni/golang-car-web-api/config"
	"github.com/salmantaghooni/golang-car-web-api/constants"
	"github.com/salmantaghooni/golang-car-web-api/data/cache"
	"github.com/salmantaghooni/golang-car-web-api/pkg/logging"
	"github.com/salmantaghooni/golang-car-web-api/pkg/service_errors"
)

type OtpService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, redisClient: redis, cfg: cfg}
}

func (s *OtpService) SetOtp(mobile_number string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOTPDefaultKey, mobile_number)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}
	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OTPExists}
	} else if err != nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OTPUsed}
	}
	err = cache.Set[*OtpDto](s.redisClient, key, val, s.cfg.OTP.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) ValidateOtp(mobile_number string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOTPDefaultKey, mobile_number)
	res, err := cache.Get[*OtpDto](s.redisClient, key)
	if err != nil {
		return err
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OTPUsed}
	} else if err != nil && !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OTPNotValid}
	} else if err != nil && !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set[*OtpDto](s.redisClient, key, res, s.cfg.OTP.ExpireTime*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
