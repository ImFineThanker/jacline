package main

import (
	"time"
)

type ActivateCode struct {
	Code       string        `json:"code"`
	DeviceCode string        `json:"deviceId"`
	ValidUntil time.Duration `json:"validUntil"`
	UsedDate   time.Time     `json:"usedDate"`
}

type Copilot struct {
	jetDir            string
	vscodeSettingsDir string
	vscodePluginsDir  string
	jwtToken          string
	DeviceCode        string
}

type Jetbrains struct {
	GithubCom GithubCom `json:"github.com"`
}

type GithubCom struct {
	User        string      `json:"user"`
	OauthToken  string      `json:"oauth_token"`
	DevOverride DevOverride `json:"dev_override"`
}

type DevOverride struct {
	CopilotTokenUrl string `json:"copilot_token_url"`
}

type LocalActivateCode struct {
	Codes []string `json:"codes"`
}
