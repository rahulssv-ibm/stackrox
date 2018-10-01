package matcher

import (
	"fmt"

	"github.com/stackrox/rox/central/searchbasedpolicies"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/search"
)

type matcherImpl struct {
	q                *v1.Query
	policyName       string
	violationPrinter searchbasedpolicies.ViolationPrinter
}

func (m *matcherImpl) errorPrefixForMatchOne(fieldLabel search.FieldLabel, id string) string {
	return fmt.Sprintf("matching policy %s against %s %s", m.policyName, fieldLabel, id)
}

func (m *matcherImpl) MatchOne(searcher searchbasedpolicies.Searcher, fieldLabel search.FieldLabel, id string) ([]*v1.Alert_Violation, error) {
	q := search.ConjunctionQuery(search.NewQueryBuilder().AddStrings(fieldLabel, id).ProtoQuery(), m.q)
	results, err := searcher.Search(q)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	if len(results) > 1 {
		return nil, fmt.Errorf("%s: got more than one result: %+v", m.errorPrefixForMatchOne(fieldLabel, id), results)
	}
	result := results[0]
	if result.ID != id {
		return nil, fmt.Errorf("%s: id of result %+v did not match passed id", m.errorPrefixForMatchOne(fieldLabel, id), result)
	}

	violations := m.violationPrinter(result)
	if len(violations) == 0 {
		return nil, fmt.Errorf("%s: result matched query but couldn't find any violation messages: %+v", m.errorPrefixForMatchOne(fieldLabel, id), result)
	}
	return violations, nil
}

func (m *matcherImpl) Match(searcher searchbasedpolicies.Searcher) (map[string][]*v1.Alert_Violation, error) {
	results, err := searcher.Search(m.q)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	violationsMap := make(map[string][]*v1.Alert_Violation, len(results))
	for _, result := range results {
		if result.ID == "" {
			return nil, fmt.Errorf("matching policy %s: got empty result id: %+v", m.policyName, result)
		}

		violations := m.violationPrinter(result)
		if len(violations) == 0 {
			return nil, fmt.Errorf("matching policy %s: result matched query but couldn't find any violation messages: %+v", m.policyName, result)
		}
		violationsMap[result.ID] = violations
	}
	return violationsMap, nil
}
