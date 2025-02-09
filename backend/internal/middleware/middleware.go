package middleware

import (
	"fmt"
	"regexp"
	
	"backend/models"
)

func isValidIP(ip string) bool {
	ipRegex := `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	re := regexp.MustCompile(ipRegex)
	return re.MatchString(ip)
}

func isValidPingTime(pingTime string) bool {
	pingTimeRegex := `^\d+(\.\d+)?(ns|µs|ms|s)$`
	re := regexp.MustCompile(pingTimeRegex)
	return re.MatchString(pingTime)
}

func isValidDate(date string) bool {
	timestampRegex := `^\d{2}:\d{2}:\d{4}:\d{2}:\d{2}:\d{2}\.\d{3}$`
	re := regexp.MustCompile(timestampRegex)
	return re.MatchString(date)
}

func ValidatePingResults(pingResults []models.PingResult) error {
	for _, result := range pingResults {
		if !isValidIP(result.IP) {
			return fmt.Errorf("невалидный IP адрес: %s", result.IP)
		}
		if !isValidPingTime(result.PingTime) {
			return fmt.Errorf("невалидное время пинга: %s", result.PingTime)
		}
		if !isValidDate(result.Date) {
			return fmt.Errorf("невалидная дата: %s", result.Date)
		}
	}
	return nil
}
