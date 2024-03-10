package constant

const (
	RoleVoiceRecognize = "voice_recognize"
	RoleVoiceGenerate  = "voice_generate"
	RoleCore           = "core"
	RoleAI             = "ai"
	RoleClient         = "client"
	RolePlugin         = "plugin"
	RoleDesktop        = "desktop"

	GetAvailablePortRandomTimes = 3
)

var (
	RoleCollection = []string{RoleAI, RoleClient, RoleCore, RolePlugin, RoleVoiceGenerate, RoleVoiceRecognize}
)

const (
	GetModuleByRoleURL       = "/module/role"
	GetModuleByRoleQueryRole = "role"

	VoiceRecognizeURL          = "/voice"
	AICompletionsURL           = "/chat/completions"
	FrameWorkPullCenterURL     = "/module/pull"
	CorePullCenterURL          = "/pull"
	DesktopNewWindowURL        = "/window"
	DesktopOpenFileSelectorURL = "/fileselector"
	DesktopRegisterHotKeyURL   = "/hotkey"
	DesktopWindowExposeURL     = "/window"
	DesktopCursorPossition     = "/info/cursorposition"
	DesktopWorkArea            = "/info/workarea"

	CoreVoiceCommand         = "/command/voice"
	CoreVoiceCommandJSONPath = "path"

	ProxyAICompletionsURL             = "/proxy/ai/chat/completions"
	ProxyAICompletionsJSONUserCommand = "usercommand"
	ProxyAICompletionJSONContent      = "content"
	ProxyVoiceRecognizeVoiceURL       = "/proxy/voice_recognize/voice"
	ProxyVoiceRecognizeVoiceJSONPath  = "path"
	ProxyPluginURL                    = "/proxy/plugin"
	ProxyDesktopFileSelectorURL       = "/proxy/desktop/fileselector"
	ProxyDesktopNewWindowURL          = "/proxy/desktop/window"
	ProxyDesktopRegisterHotKeyURL     = "/proxy/desktop/hotkey"
	ProxyDesktopWindowExposeURL       = "/proxy/desktop/window"
	ProxyDesktopCursorPossition       = "/proxy/desktop/info/cursorposition"
	ProxyDesktopWorkArea              = "/proxy/desktop/info/workarea"
)
