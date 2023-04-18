package deepcopy

type options struct {
	// MaxDepth is the maximum depth to traverse.
	// If MaxDepth is 0, it will be treated as no limit.
	MaxDepth int
	// TagName is the tag name to use.
	// If TagName is empty, it will be treated as no tag.
	TagName string
	// // OnlyField is the field name to copy.
	// // If OnlyField is empty, it will be treated as no field.
	// OnlyField string
}

type Option func(*options)

func WithMaxDepth(maxDepth int) Option {
	return func(o *options) {
		o.MaxDepth = maxDepth
	}
}

func WithTagName(tagName string) Option {
	return func(o *options) {
		o.TagName = tagName
	}
}
