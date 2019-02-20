import React from 'react';
import Widget from 'Components/Widget';
import Query from 'Components/ThrowingQuery';
import Loader from 'Components/Loader';
import PropTypes from 'prop-types';
import Gauge from 'Components/visuals/GaugeWithDetail';
import NoResultsMessage from 'Components/NoResultsMessage';
import { standardBaseTypes } from 'constants/entityTypes';
import { standardShortLabels } from 'messages/standards';
import { resourceLabels } from 'messages/common';
import { AGGREGATED_RESULTS } from 'queries/controls';
import URLService from 'modules/URLService';
import contextTypes from 'constants/contextTypes';
import pageTypes from 'constants/pageTypes';

const isStandard = type => !!standardBaseTypes[type];

const sortByTitle = (a, b) => {
    if (a.title < b.title) return -1;
    if (a.title > b.title) return 1;
    return 0;
};

function processData({ entityType, query }, { results, complianceStandards }) {
    let filteredResults;
    if (standardBaseTypes[entityType]) {
        filteredResults = results.results.filter(result =>
            result.aggregationKeys[0].id.includes(entityType)
        );
    } else {
        filteredResults = results.results;
    }
    if (!filteredResults.length)
        return [
            {
                title: entityType,
                passing: { value: 0, link: '' },
                failing: { value: 0, link: '' }
            }
        ];
    const standardDataMapping = filteredResults
        .filter(datum => !(datum.passing === 0 && datum.failing === 0))
        .reduce((accMapping, currValue) => {
            const newMapping = { ...accMapping };
            const { id: standardId } = currValue.aggregationKeys[0];
            const standard = complianceStandards.find(cs => cs.id === standardId);
            let { numPassing: totalPassing, numFailing: totalFailing } = currValue;
            if (newMapping[standardId]) {
                totalPassing += newMapping[standardId].passing.value;
                totalFailing += newMapping[standardId].failing.value;
            }
            const newQuery = { ...query };
            newQuery['Compliance State'] = 'Passing';
            newQuery.Standard = standard.name;
            const passingLink = URLService.getLinkTo(contextTypes.COMPLIANCE, pageTypes.LIST, {
                entityType,
                query: newQuery
            });
            newQuery['Compliance State'] = 'Failing';
            newQuery.Standard = standard.name;
            const failingLink = URLService.getLinkTo(contextTypes.COMPLIANCE, pageTypes.LIST, {
                entityType,
                query: newQuery
            });
            delete newQuery['Compliance State'];
            delete newQuery.Standard;
            const defaultLink = URLService.getLinkTo(contextTypes.COMPLIANCE, pageTypes.LIST, {
                entityType,
                query: newQuery
            });
            newMapping[standardId] = {
                title: standardShortLabels[standard.id],
                passing: {
                    value: totalPassing,
                    link: passingLink.url
                },
                failing: {
                    value: totalFailing,
                    link: failingLink.url
                },
                defaultLink: defaultLink.url
            };
            return newMapping;
        }, {});
    return Object.values(standardDataMapping).sort(sortByTitle);
}

const getQueryVariables = params => {
    const groupBy = ['STANDARD'];
    let unit = 'CONTROL';

    if (params.query && params.query.groupBy) {
        groupBy.push(params.query.groupBy);
    } else if (!isStandard(params.entityType)) {
        groupBy.push(params.entityType);
        unit = params.entityType;
    }
    const { query } = params;
    return {
        groupBy,
        unit,
        query
    };
};

const ComplianceAcrossEntities = ({ params }) => {
    const variables = getQueryVariables(params);
    return (
        <Query query={AGGREGATED_RESULTS} variables={variables}>
            {({ loading, data }) => {
                let contents = <Loader />;
                const headerText = standardBaseTypes[params.entityType]
                    ? `Controls in Compliance`
                    : `${resourceLabels[params.entityType]}s in Compliance`;
                if (!loading && data) {
                    const results = processData(params, data);
                    if (!results.length) {
                        contents = (
                            <NoResultsMessage message="No data available. Please run a scan." />
                        );
                    } else {
                        contents = <Gauge data={results} />;
                    }
                }
                return (
                    <Widget header={headerText} bodyClassName="p-2">
                        {contents}
                    </Widget>
                );
            }}
        </Query>
    );
};

ComplianceAcrossEntities.propTypes = {
    params: PropTypes.shape({}).isRequired
};

export default ComplianceAcrossEntities;
