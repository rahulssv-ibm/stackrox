import React from 'react';
import PropTypes from 'prop-types';
import getEntityName from 'utils/getEntityName';
import { entityNameQueryMap } from 'utils/queryMap';
import entityLabels from 'messages/entity';
import entityTypes from 'constants/entityTypes';

import Query from 'Components/ThrowingQuery';
import PageHeader from 'Components/PageHeader';
import startCase from 'lodash/startCase';
import ExportButton from 'Components/ExportButton';

const getEntityVariables = (type, id) => {
    if (type === entityTypes.SUBJECT) {
        return { name: id };
    }
    return { id };
};

const getQueryAndVariables = (entityType, entityId) => {
    const query = entityNameQueryMap[entityType] || null;
    const safeEntityId = decodeURIComponent(entityId); // fix bug  ROX-4543-fix-bad-encoding-in-config-mgt-API-request
    return {
        query,
        variables: getEntityVariables(entityType, safeEntityId),
    };
};

const EntityPageHeader = ({ entityType, entityId, urlParams }) => {
    const { query, variables } = getQueryAndVariables(entityType, entityId);
    if (!query) return null;

    return (
        <Query query={query} variables={variables}>
            {({ data }) => {
                const header = getEntityName(entityType, data, entityId) || '-';
                const subHeader = entityLabels[entityType];
                const exportFilename = `${startCase(subHeader)} Report: "${header}"`;

                let pdfId = 'capture-dashboard-stretch';
                if (urlParams && urlParams.entityListType1) {
                    pdfId = 'capture-list';
                }
                return (
                    <PageHeader classes="z-1" header={header} subHeader={subHeader}>
                        <div className="flex flex-1 justify-end">
                            <div className="flex">
                                <div className="flex items-center">
                                    <ExportButton
                                        fileName={exportFilename}
                                        type={entityType}
                                        page="configManagement"
                                        pdfId={pdfId}
                                    />
                                </div>
                            </div>
                        </div>
                    </PageHeader>
                );
            }}
        </Query>
    );
};

EntityPageHeader.propTypes = {
    entityType: PropTypes.string.isRequired,
    entityId: PropTypes.string.isRequired,
    urlParams: PropTypes.shape({
        entityListType1: PropTypes.string,
    }),
};

EntityPageHeader.defaultProps = {
    urlParams: null,
};

export default EntityPageHeader;
