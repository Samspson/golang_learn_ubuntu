package search

// 实现默认匹配器
type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// 实现了默认的匹配器的行为
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
