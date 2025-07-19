package filters

import (
	"github.com/doug-martin/goqu/v9"
)

// SelectFilterBuilder is a function that builds a select filter
type SelectFilterBuilder func(query *goqu.SelectDataset) *goqu.SelectDataset

// UpdateFilterBuilder is a function that builds an update filter
type UpdateFilterBuilder func(query *goqu.UpdateDataset) *goqu.UpdateDataset

// DeleteFilterBuilder is a function that builds a delete filter
type DeleteFilterBuilder func(query *goqu.DeleteDataset) *goqu.DeleteDataset

// ApplyFilters applies all filters to the query
func ApplyFilters(q *goqu.SelectDataset, filters ...SelectFilterBuilder) *goqu.SelectDataset {
	for _, filter := range filters {
		q = filter(q)
	}
	return q
}

// IsSelectFilter creates a filter that checks if a column equals a value
func IsSelectFilter(column string, value interface{}) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Where(goqu.Ex{column: value})
	}
}

// IsUpdateFilter creates a filter that checks if a column equals a value
func IsUpdateFilter(column string, value interface{}) UpdateFilterBuilder {
	return func(query *goqu.UpdateDataset) *goqu.UpdateDataset {
		return query.Where(goqu.Ex{column: value})
	}
}

// IsDeleteFilter creates a filter that checks if a column equals a value
func IsDeleteFilter(column string, value interface{}) DeleteFilterBuilder {
	return func(query *goqu.DeleteDataset) *goqu.DeleteDataset {
		return query.Where(goqu.Ex{column: value})
	}
}

// NotSelectFilter creates a filter that checks if a column does not equal a value
func NotSelectFilter(column string, value interface{}) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Where(goqu.Ex{column: goqu.Op{"neq": value}})
	}
}

// InSelectFilter creates a filter that checks if a column is in a list of values
func InSelectFilter(column string, values ...interface{}) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Where(goqu.Ex{column: values})
	}
}

// LikeSelectFilter creates a filter that checks if a column matches a pattern
func LikeSelectFilter(column string, pattern string) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Where(goqu.Ex{column: goqu.Op{"like": pattern}})
	}
}

// OrderByFilter creates a filter that orders the results by a column
func OrderByFilter(column string, direction string) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		if direction == "desc" {
			return query.Order(goqu.I(column).Desc())
		}
		return query.Order(goqu.I(column).Asc())
	}
}

// LimitFilter creates a filter that limits the number of results
func LimitFilter(limit uint) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Limit(limit)
	}
}

// OffsetFilter creates a filter that skips a number of results
func OffsetFilter(offset uint) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Offset(offset)
	}
}

// ApplySelectFilters applies all filters to a query
func ApplySelectFilters(query *goqu.SelectDataset, filters ...SelectFilterBuilder) *goqu.SelectDataset {
	for _, filter := range filters {
		query = filter(query)
	}
	return query
}

// ApplyUpdateFilters applies all filters to a query
func ApplyUpdateFilters(query *goqu.UpdateDataset, filters ...UpdateFilterBuilder) *goqu.UpdateDataset {
	for _, filter := range filters {
		query = filter(query)
	}
	return query
}

// ApplyDeleteFilters applies all filters to a query
func ApplyDeleteFilters(query *goqu.DeleteDataset, filters ...DeleteFilterBuilder) *goqu.DeleteDataset {
	for _, filter := range filters {
		query = filter(query)
	}
	return query
}

// WithJoin creates a filter that joins a table
func WithJoin(table, leftKey, rightKey string) SelectFilterBuilder {
	return func(query *goqu.SelectDataset) *goqu.SelectDataset {
		return query.Join(
			goqu.T(table),
			goqu.On(goqu.I(table+"."+rightKey), goqu.I(leftKey)),
		)
	}
}
