package config

func WebsiteURL() string {
	return Get().WebsiteProtocol+Get().WebsiteDomain
}