package database

type FindOption interface {
	apply(*option)
}

type option struct {
	query        []Query
	order        any
	offset       int
	limit        int
	preloads     []string
	joins        []string
	selectFields string
	groupBy      string
	having       string
	havingArgs   []interface{}
}

type optionFn func(*option)

func (f optionFn) apply(opt *option) {
	f(opt)
}

func WithQuery(query ...Query) FindOption {
	return optionFn(func(opt *option) {
		opt.query = query
	})
}

func WithOffset(offset int) FindOption {
	return optionFn(func(opt *option) {
		opt.offset = offset
	})
}

func WithLimit(limit int) FindOption {
	return optionFn(func(opt *option) {
		opt.limit = limit
	})
}

func WithOrder(order interface{}) FindOption {
	return optionFn(func(opt *option) {
		opt.order = order
	})
}

func WithPreload(preloads []string) FindOption {
	return optionFn(func(opt *option) {
		opt.preloads = preloads
	})
}

func WithJoin(joins ...string) FindOption {
	return optionFn(func(opt *option) {
		opt.joins = append(opt.joins, joins...)
	})
}

func WithSelect(selectFields string) FindOption {
	return optionFn(func(opt *option) {
		opt.selectFields = selectFields
	})
}

func WithGroupBy(groupBy string) FindOption {
	return optionFn(func(opt *option) {
		opt.groupBy = groupBy
	})
}

func WithHaving(having string, args ...interface{}) FindOption {
	return optionFn(func(opt *option) {
		opt.having = having
		opt.havingArgs = args
	})
}

func getOption(opts ...FindOption) option {
	opt := option{
		query:  []Query{},
		offset: 0,
		limit:  1000,
		order:  "id",
	}

	for _, o := range opts {
		o.apply(&opt)
	}

	return opt
}
