package client

import (
	"encoding/json"
	"time"
)

type Info struct {
	Name           string `json:"name"`
	IP             string `json:"ip"`
	CurrentTime    string `json:"current_time"`
	Uptime         string `json:"uptime"`
	CPUTemperature string `json:"cpu_temperature"`
	Version        string `json:"version"`
}

type Capabilities struct {
	DevMode       bool `json:"dev_mode"`
	Dashboard     bool `json:"dashboard"`
	HealthCheck   bool `json:"health_check"`
	Equipment     bool `json:"equipment"`
	Timers        bool `json:"timers"`
	Lighting      bool `json:"lighting"`
	Temperature   bool `json:"temperature"`
	ATO           bool `json:"ato"`
	Camera        bool `json:"camera"`
	Doser         bool `json:"doser"`
	Ph            bool `json:"ph"`
	Macro         bool `json:"macro"`
	Configuration bool `json:"configuration"`
}

type HealthCheckNotify struct {
	Enable    bool    `json:"enable"`
	MaxMemory float64 `json:"max_memory"`
	MaxCPU    float64 `json:"max_cpu"`
}

type Settings struct {
	Name         string            `json:"name"`
	Interface    string            `json:"interface"`
	Address      string            `json:"address"`
	Display      bool              `json:"display"`
	Notification bool              `json:"notification"`
	Capabilities Capabilities      `json:"capabilities"`
	HealthCheck  HealthCheckNotify `json:"health_check"`
	HTTPS        bool              `json:"https"`
	Pprof        bool              `json:"pprof"`
	RPI_PWMFreq  int               `json:"rpi_pwm_freq"`
}

type AdafruitIO struct {
	Enable bool   `json:"enable"`
	Token  string `json:"token"`
	User   string `json:"user"`
	Prefix string `json:"prefix"`
}

type TelemetryConfig struct {
	AdafruitIO      AdafruitIO   `json:"adafruitio"`
	Mailer          MailerConfig `json:"mailer"`
	Notify          bool         `json:"notify"`
	Throttle        int          `json:"throttle"`
	HistoricalLimit int          `json:"historical_limit"`
	CurrentLimit    int          `json:"current_limit"`
}

type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Error struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	Count   int    `json:"count"`
	ID      string `json:"id"`
}

type StatsResponse struct {
	Current    []interface{} `json:"current"`
	Historical []interface{} `json:"historical"`
}

type Dashboard struct {
	Column      int       `json:"column"`
	Row         int       `json:"row"`
	Width       int       `json:"width"`
	Height      int       `json:"height"`
	GridDetails [][]Chart `json:"grid_details"`
}

type Chart struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type DisplayState struct {
	On         bool `json:"on"`
	Brightness int  `json:"brightness"`
}

type DisplayConfig struct {
	Enable     bool `json:"enable"`
	Brightness int  `json:"brightness"`
}

type MailerConfig struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	From     string `json:"from"`
	Password string `json:"password"`
	To       string `json:"to"`
}

type AnalogInput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Pin    int    `json:"pin"`
	Driver string `json:"driver"`
}

type ATO struct {
	ID             string        `json:"id"`
	Inlet          string        `json:"inlet"`
	Pump           string        `json:"pump"`
	Period         time.Duration `json:"period"`
	Control        bool          `json:"control"`
	Enable         bool          `json:"enable"`
	Notify         Notify        `json:"notify"`
	Name           string        `json:"name"`
	DisableOnAlert bool          `json:"disable_on_alert"`
}
type Pump struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Jack     string         `json:"jack"`
	Pin      int            `json:"pin"`
	Regiment DosingRegiment `json:"regiment"`
}
type DosingRegiment struct {
	Enable   bool     `json:"enable"`
	Schedule Schedule `json:"schedule"`
	Duration float64  `json:"duration"`
	Speed    float64  `json:"speed"`
}

type CalibrationDetails struct {
	Speed    float64 `json:"speed"`
	Duration float64 `json:"duration"`
}

type Schedule struct {
	Day    string `json:"day"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
	Second string `json:"second"`
}
type Driver struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}
type Equipment struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Outlet string `json:"outlet"`
	On     bool   `json:"on"`
}
type Inlet struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Pin       int    `json:"pin"`
	Equipment string `json:"equipment"`
	Reverse   bool   `json:"reverse"`
	Driver    string `json:"driver"`
}
type Jack struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Pins   []int  `json:"pins"`
	Driver string `json:"driver"` // can be either hal or pca9685
}
type Channel struct {
	Name     string  `json:"name"`
	Min      int     `json:"min"`
	StartMin int     `json:"start_min"`
	Max      int     `json:"max"`
	Reverse  bool    `json:"reverse"`
	Pin      int     `json:"pin"`
	Color    string  `json:"color"`
	Profile  Profile `json:"profile"`
}
type Light struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Channels map[int]Channel `json:"channels"`
	Jack     string          `json:"jack"`
}

type Profile struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

type AutoConfig struct {
	Values []int `json:"values"` // 12 ticks after every 2 hours
}

type FixedConfig struct {
	Value int `json:"value"`
}

type DiurnalConfig struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Outlet struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Pin       int    `json:"pin"`
	Equipment string `json:"equipment"`
	Reverse   bool   `json:"reverse"`
	Driver    string `json:"driver"`
}
type Probe struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Enable      bool          `json:"enable"`
	Period      time.Duration `json:"period"`
	AnalogInput string        `json:"analog_input"`
	Control     bool          `json:"control"`
	Notify      Notify        `json:"notify"`
}
type TC struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Max        float64       `json:"max"`
	Min        float64       `json:"min"`
	Heater     string        `json:"heater"`
	Cooler     string        `json:"cooler"`
	Period     time.Duration `json:"period"`
	Control    bool          `json:"control"`
	Enable     bool          `json:"enable"`
	Notify     Notify        `json:"notify"`
	Sensor     string        `json:"sensor"`
	Fahrenheit bool          `json:"fahrenheit"`
	ChartMin   float64       `json:"chart_min"`
	ChartMax   float64       `json:"chart_max"`
}

type Notify struct {
	Enable bool    `json:"enable"`
	Max    float64 `json:"max"`
	Min    float64 `json:"min"`
}
type Reminder struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type UpdateEquipment struct {
	Revert   bool          `json:"revert"`
	ID       string        `json:"id"`
	On       bool          `json:"on"`
	Duration time.Duration `json:"duration"`
}
type Job struct {
	ID        string          `json:"id"`
	Minute    string          `json:"minute"`
	Day       string          `json:"day"`
	Hour      string          `json:"hour"`
	Second    string          `json:"second"`
	Name      string          `json:"name"`
	Type      string          `json:"type"`
	Reminder  Reminder        `json:"reminder"`
	Equipment UpdateEquipment `json:"equipment"`
	Enable    bool            `json:"enable"`
}

var StepTypes = []string{
	"wait",
	"subsystem",
	"macro",
	"equipment",
	"ato",
	"tc",
	"lighting",
	"ph",
	"doser",
	"timer",
}

type GenericStep struct {
	ID string `json:"id"`
	On bool   `json:"on"`
}

type WaitStep struct {
	Duration time.Duration `json:"duration"`
}

type Step struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

type Macro struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Steps  []Step `json:"steps"`
	Enable bool   `json:"enable"`
}

type MotionConfig struct {
	Enable bool   `json:"enable"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type CameraConfig struct {
	Enable         bool          `json:"enable"`
	ImageDirectory string        `json:"image_directory"`
	CaptureFlags   string        `json:"capture_flags"`
	TickInterval   time.Duration `json:"tick_interval"`
	Upload         bool          `json:"upload"`
	Motion         MotionConfig  `json:"motion"`
}
type ImageItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PinValues map[int]float64
