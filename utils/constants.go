package utils



// config
const (
	CONFIG_FILE_NAME = "hypermind.config"
	DEFAULT_SERVER_PORT = 9090
	DEFAULT_REDIS_SERVER_IP = "127.0.0.1"
	DEFAULT_REDIS_SERVER_PORT = "6379"
	DEFAULT_REDIS_SERVER_PASSWORD = ""
	REDIS_USER_KEY = "user"
)

// web
const (
	COOKIE_LIFE_CYCLE_MINUTES int = 60
)

// user
const (
	LOGIN_NAME_KEY string = "loginName"
	PASSWORD_KEY string = "password"
	CN_NAME_KEY string = "cnName"
	EMAIL_KEY string = "email"
	MOBILE_PHONE_KEY string = "mobilePhone"

	ROOT_USER_NAME = "root"
)

// page parameter
const (
	HOME_PAGE_KEY = "homePage"
	HOME_PAGE = "home"
	ABOUT_ME_PAGE_KEY = "aboutMePage"
	ABOUT_ME_PAGE = "about-me"
	ABOUT_WEBSITE_PAGE_KEY = "aboutWebsitePage"
	ABOUT_WEBSITE_PAGE = "about-website"
	MEETING_KANBAN_PAGE_KEY = "meetingKanbanPage"
	MEETING_KANBAN_PAGE = "meeting-kanban"
	HASH_RING_PAGE_KEY = "hashRingPage"
	HASH_RING_PAGE = "hash-ring"
)