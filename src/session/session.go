package session

import (
	"net/http"
	"sync"
)

type Session interface {
	GetUserAgent() string
	AddCookiesToRequest(req *http.Request)
	UpdateCookies(cookies []*http.Cookie) error
	GetCookies() []*http.Cookie
	FindCookie(key string) string
}

type session struct {
	sync.Mutex
	UserAgent string
	Cookies   []*http.Cookie
}

func NewSession() Session {
	return &session{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Safari/537.36",
	}
}

// FindCookie implements Session.
func (s *session) FindCookie(key string) string {
	for _, cookie := range s.Cookies {
		if cookie.Name == key {
			return cookie.Value
		}
	}

	return ""
}

// GetCookies implements Session.
func (s *session) GetCookies() []*http.Cookie {
	return s.Cookies
}

// GetUserAgent implements Session.
func (s *session) GetUserAgent() string {
	return s.UserAgent
}

// AddCookiesToRequest implements Session.
func (s *session) AddCookiesToRequest(req *http.Request) {
	for _, cookie := range s.Cookies {
		req.AddCookie(cookie)
	}
}

// UpdateCookies implements Session.
func (s *session) UpdateCookies(cookies []*http.Cookie) error {
	s.Lock()
	defer s.Unlock()

	for _, cookie := range cookies {
		err := s.updateCookies(cookie)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *session) updateCookies(cookie *http.Cookie) error {
	fixCookies := []*http.Cookie{}

	isExist := false
	for _, oldCookie := range s.Cookies {
		if oldCookie.Name == cookie.Name {
			fixCookies = append(fixCookies, cookie)
			isExist = true
		} else {
			fixCookies = append(fixCookies, oldCookie)
		}
	}

	if !isExist {
		fixCookies = append(fixCookies, cookie)
	}

	s.Cookies = fixCookies

	return nil
}
