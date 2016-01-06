package tldr

type PageFetcher interface {
	Fetch(page, platform string) (Page, error)
}
