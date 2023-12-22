package provider

type Healthcheck func(name string, delay uint16)

var HealthcheckHook Healthcheck

func (pp *proxySetProvider) Count() int {
	return len(pp.proxies)
}
