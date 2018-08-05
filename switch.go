// Package mystrom provides a convinent way to access your myStrom Switch device
// (https://mystrom.ch/de/wifi-switch/) via your own application.
package mystrom

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	apiOn     = "relay?state=1"
	apiOff    = "relay?state=0"
	apiToggle = "toggle"
	apiReport = "report"
	apiTemp   = "temp"
)

// Switch holds all info and logic to talk your myStrom Switch device.
type Switch struct {
	url string
}

// NewSwitch requires an IP or Hostname of your Switch device and returns a Client.
func NewSwitch(host string) (*Switch, error) {
	s := &Switch{}

	if host == "" {
		return s, errors.New("Hostname or IP must be provided")
	}
	s.url = "http://" + host + "/"

	return s, nil
}

// Toggle toggles the power state of the Switch.
func (s Switch) Toggle() error {
	_, err := http.Get(s.url + apiToggle)
	return err
}

// On turns the power of the Switch on.
func (s Switch) On() error {
	_, err := http.Get(s.url + apiOn)
	return err
}

// Off turns the power of the Switch off.
func (s Switch) Off() error {
	_, err := http.Get(s.url + apiOff)
	return err
}

// SwitchReport represets the content of a report of the Switch.
type SwitchReport struct {
	Power float64 `json:"power"` // current power consumption in watts
	Relay bool    `json:"relay"` // state of the Switch, true is on, false is off
}

// Report returns a report of the current statut of the Switch.
func (s Switch) Report() (*SwitchReport, error) {
	r := &SwitchReport{}

	resp, err := http.Get(s.url + apiReport)
	if err != nil {
		return r, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(contents, r)
	return r, err
}

// SwitchTemperature represets the content of a temperature response of
// the Switch. All temperatures are provided in °C.
type SwitchTemperature struct {
	Measured     float64 `json:"measured"`     // measured temp
	Compensation float64 `json:"compensation"` // assumed gap
	Compensated  float64 `json:"compensated"`  // real temp
}

// Temperature returns the current temperature in °C.
func (s Switch) Temperature() (*SwitchTemperature, error) {
	t := &SwitchTemperature{}

	resp, err := http.Get(s.url + apiTemp)
	if err != nil {
		return t, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(contents, t)
	return t, err
}
