package hw10programoptimization

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	scanner := bufio.NewScanner(r)
	var user User
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return nil, fmt.Errorf("unmarshal error: %w", err)
		}
		countDomains(&user, domain, result)
	}
	return result, nil
}

func countDomains(u *User, domain string, result DomainStat) {
	atIndex := strings.Index(u.Email, "@")
	if atIndex == -1 {
		return
	}
	emailDomain := strings.ToLower(u.Email[atIndex+1:])
	if strings.HasSuffix(emailDomain, domain) {
		result[emailDomain]++
	}
}
