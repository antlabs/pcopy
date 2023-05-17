// Copyright [2020-2023] [guonaihong]
package pcopy

type options struct {
	// // MaxDepth is the maximum depth to traverse.
	// // If MaxDepth is 0, it will be treated as no limit.
	// maxDepth int
	// // TagName is the tag name to use.
	// // If TagName is empty, it will be treated as no tag.
	// tagName string
	// // OnlyField is the field name to copy.
	// // If OnlyField is empty, it will be treated as no field.
	// OnlyField string
	preheat bool

	// 使用预热cache
	usePreheat bool
}

type Option func(*options)

// func WithMaxDepth(maxDepth int) Option {
// 	return func(o *options) {
// 		o.maxDepth = maxDepth
// 	}
// }

// func WithTagName(tagName string) Option {
// 	return func(o *options) {
// 		o.tagName = tagName
// 	}
// }

func WithUsePreheat() Option {
	return func(o *options) {
		o.usePreheat = true
	}
}

// 内部函数
func withPreheat() Option {
	return func(o *options) {
		o.preheat = true
	}
}
