package app

import (
	"net/http"
)

type ResourceResolver interface {
	Resolve(string) string
}

func LocalDir(directory string) ResourceResolver {
	_ = "STUB: not implemented"
	return *new(ResourceResolver)
}

type localResourceResolver struct {
	http.Handler
	directory string
}

func (r localResourceResolver) Resolve(location string) string {
	_ = "STUB: not implemented"
	return ""
}

func RemoteBucket(url string) ResourceResolver {
	_ = "STUB: not implemented"
	return *new(ResourceResolver)
}

type remoteResourceResolver struct {
	url string
}

func (r remoteResourceResolver) Resolve(location string) string {
	_ = "STUB: not implemented"
	return ""
}

func PrefixedLocation(prefix string) ResourceResolver {
	_ = "STUB: not implemented"
	return *new(ResourceResolver)
}

type prefixedResourceResolver struct {
	localResourceResolver
	prefix string
}

func (r prefixedResourceResolver) Resolve(location string) string {
	_ = "STUB: not implemented"
	return ""
}

func GitHubPages(repositoryName string) ResourceResolver {
	_ = "STUB: not implemented"
	return *new(ResourceResolver)
}

func clientResourceResolver(resourcesLocation string) func(string) string {
	_ = "STUB: not implemented"
	return nil
}

func resolveOGResource(domain string, location string) string { _ = "STUB: not implemented"; return "" }

func remoteLocation(location string) bool { _ = "STUB: not implemented"; return false }

func webLocation(location string) bool { _ = "STUB: not implemented"; return false }

type ProxyResource struct {
	Path string

	ResourcePath string
}
