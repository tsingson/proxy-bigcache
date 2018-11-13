package main

type (
	UmsConfig struct {
		ActiveAuthURI   string `json:"ActiveAuthUri"`
		RegisterAuthURI string `json:"RegisterAuthUri"`
		PlayAuthURI     string `json:"PlayAuthUri"`
	}
	AaaConfig struct {
		ServerPort        string `json:"ServerPort"`
		LogFilenamePrefix string `json:"LogFilenamePrefix"`
	}
	CdnConfig struct {
		VodGslb  []string `json:"VodGslb"`
		LiveGslb []string `json:"LiveGslb"`   
	}
	ApkEpgConfig struct {
		ApkName   string   `json:"ApkName"`
		ApkType   string   `json:"ApkType"`
		ApkEpgURI []string `json:"ApkEpgURI"`
	}

	Config struct {
		Name         string         `json:"Name"`
		Version      string         `json:"Version"`
		Debug        bool           `json:"Debug"`
		UmsConfig    UmsConfig      `json:"UmsConfig"`
		AaaConfig    AaaConfig      `json:"AaaConfig"`
		CdnConfig    CdnConfig      `json:"CdnConfig"`
		ApkEpgConfig []ApkEpgConfig `json:"ApkEpgConfig"`
	}
)

// design and code by tsingson
