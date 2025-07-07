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

var WhiteList = [][]string{
	browsers,
}
