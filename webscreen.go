package webscreen

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/tebeka/selenium/firefox"
	"log"
)

type Engine struct {
	service     *selenium.Service
	drivers     []selenium.WebDriver
	driverType  SeleniumDriver
	seleniumURL string
}

func NewEngineFromConfig(config *SeleniumConfig) (*Engine, error) {
	e := &Engine{}
	opts := make([]selenium.ServiceOption, 0)
	switch config.DriverType {
	case ChromeDriver:
		opts = append(opts, selenium.ChromeDriver(config.DriverPath))
	case GeckoDriver:
		opts = append(opts, selenium.GeckoDriver(config.DriverPath))
	case HTMLUnitDriver:
		opts = append(opts, selenium.HTMLUnit(config.DriverPath))
	}
	service, err := selenium.NewSeleniumService(config.ServerJarPath, 8080, opts...)
	if err != nil {
		return nil, err
	}
	e.service = service
	e.seleniumURL = "http://localhost:8080/wd/hub"
	e.driverType = config.DriverType
	return e, nil
}

func (e *Engine) Stop() {
	log.Println("shutting down engine")
	for _, wd := range e.drivers {
		_ = wd.Quit()
	}
	if e.service != nil {
		_ = e.service.Stop()
	}
	log.Println("engine stopped")
}

func (e *Engine) NewRunner() (*Runner, error) {
	caps := selenium.Capabilities{}
	switch e.driverType {
	case ChromeDriver:
		caps.AddChrome(chrome.Capabilities{})
	case GeckoDriver:
		caps.AddFirefox(firefox.Capabilities{})
	}
	wd, err := selenium.NewRemote(caps, e.seleniumURL)
	if err != nil {
		return nil, err
	}
	e.drivers = append(e.drivers, wd)
	return &Runner{wd: wd}, nil
}
