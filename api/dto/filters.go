package dto

import (
	"fmt"
	"slices"
	"strings"
)

type OrderFilter struct {
	OrderBy string `json:"order_by" query:"order_by"`
}

func (f OrderFilter) ParseOrderFilter(safeFilters []string) string {
	orderMap := make([]string, 0)

	items := strings.Split(f.OrderBy, ",")
	for _, item := range items {
		parts := strings.Split(strings.TrimSpace(item), ":")
		if len(parts) == 2 && slices.Contains(safeFilters, parts[0]) && slices.Contains([]string{"asc", "desc"}, parts[1]) {
			orderMap = append(orderMap, fmt.Sprintf("%s %s", parts[0], parts[1]))
		}
	}

	result := strings.Join(orderMap, ",")
	return result
}

type SearchFilter struct {
	Q string `json:"q" query:"q"`
}

func (f SearchFilter) GetSearch(search []string) string {
	if f.Q != "" {
		var conditions []string
		for _, s := range search {
			conditions = append(
				conditions,
				fmt.Sprintf("(to_tsvector('simple', %s) @@ plainto_tsquery('simple', '%s'))", s, f.Q),
			)
		}
		combinedConditions := fmt.Sprintf("(%s)", strings.Join(conditions, " OR "))
		return combinedConditions
	}
	return ""
}
