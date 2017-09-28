package config

func (c cfg) WebsiteURL() string {
	return c.WebsiteProtocol + c.WebsiteDomain
}
