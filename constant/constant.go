package constant

const (
	RoleVoiceRecognize = "voice_recognize"
	RoleVoiceGenerate  = "voice_generate"
	RoleCore           = "core"
	RoleAI             = "ai"
	RoleClient         = "client"
	RolePlugin         = "plugin"

	GetAvailablePortRandomTimes = 3
)

var (
	RoleCollection = []string{RoleAI, RoleClient, RoleCore, RolePlugin, RoleVoiceGenerate, RoleVoiceRecognize}
)

const (
	GetModuleByRoleURL       = "/module/role"
	GetModuleByRoleQueryRole = "role"

	VoiceRecognizeURL      = "/voice"
	AICompletionsURL       = "/chat/completions"
	FrameWorkPullCenterURL = "/module/pull"
	CorePullCenterURL      = "/pull"
)
