package freebox

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Device :
type Device struct {
	Name           string
	BoxModel       string
	BoxModelName   string
	Host           string
	IP             string
	IPv6           string
	PortHTTP       int
	PortHTTPS      int
	HTTPSAvailable bool
	APIVersion     string
	APIDomain      string
	APIBaseURL     string
	UID            string
}

func (d Device) majorAPIVersion() (int, error) {
	re := regexp.MustCompile(`([\d]+)\.[\d]+`)
	res := re.FindStringSubmatch(d.APIVersion)
	if len(res) == 2 {
		major, _ := strconv.Atoi(res[1])
		return major, nil
	}
	return 0, errors.New("Major API version not found")
}

// URL :
func (d Device) Url(scheme string) string {
	if scheme == "" {
		scheme = "http"
	}

	version := ""
	if major, err := d.majorAPIVersion(); err == nil {
		version = fmt.Sprintf("v%d/", major)
	}

	/*
	if d.HTTPSAvailable {
		return fmt.Sprintf("https://%s%s%s", d.Host, d.APIBaseURL, version)
	}
	*/

	return fmt.Sprintf("%s://%s%s%s", scheme, d.Host, d.APIBaseURL, version)
}
