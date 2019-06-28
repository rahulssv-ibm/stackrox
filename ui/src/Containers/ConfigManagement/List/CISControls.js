import React from 'react';
import PropTypes from 'prop-types';
import entityTypes from 'constants/entityTypes';
import { standardLabels } from 'messages/standards';
import { LIST_STANDARD as QUERY } from 'queries/standard';
import queryService from 'modules/queryService';
import { sortVersion } from 'sorters/sorters';

import { defaultHeaderClassName, defaultColumnClassName } from 'Components/Table';
import LabelChip from 'Components/LabelChip';
import List from './List';

const tableColumns = [
    {
        Header: 'Id',
        headerClassName: 'hidden',
        className: 'hidden',
        accessor: 'id'
    },
    {
        Header: `Standard`,
        headerClassName: `w-1/8 ${defaultHeaderClassName}`,
        className: `w-1/8 ${defaultColumnClassName}`,
        accessor: 'standard'
    },
    {
        Header: `Control`,
        headerClassName: `w-1/2 ${defaultHeaderClassName}`,
        className: `w-1/2 ${defaultColumnClassName}`,
        accessor: 'control',
        sortMethod: sortVersion
    },
    {
        Header: `Control Status`,
        headerClassName: `w-1/8 ${defaultHeaderClassName}`,
        className: `w-1/8 ${defaultColumnClassName}`,
        // eslint-disable-next-line
        Cell: ({ original }) => {
            return !original.passing ? <LabelChip text="Fail" type="alert" /> : 'Pass';
        }
    },
    {
        Header: `Nodes`,
        headerClassName: `w-1/8 ${defaultHeaderClassName}`,
        className: `w-1/8 ${defaultColumnClassName}`,
        Cell: ({ original }) => {
            const { length } = original.nodes;
            if (length > 1) return `${length} Nodes`;
            return original.nodes[0];
        }
    }
];

const createTableRows = data => {
    if (!data || !data.results || !data.results.results.length) return [];

    let standardKeyIndex = 0;
    let controlKeyIndex = 0;
    let nodeKeyIndex = 0;
    data.results.results[0].aggregationKeys.forEach(({ scope }, idx) => {
        if (scope === entityTypes.STANDARD) standardKeyIndex = idx;
        if (scope === entityTypes.CONTROL) controlKeyIndex = idx;
        if (scope === entityTypes.NODE) nodeKeyIndex = idx;
    });
    const controls = {};
    data.results.results.forEach(({ keys, numFailing }) => {
        const controlId = keys[controlKeyIndex].id;
        if (controls[controlId]) {
            controls[controlId].nodes.push(keys[nodeKeyIndex].name);
            if (numFailing) {
                controls[controlId].passing = false;
            }
        } else {
            controls[controlId] = {
                id: controlId,
                standard: standardLabels[keys[standardKeyIndex].id],
                control: `${keys[controlKeyIndex].name} - ${keys[controlKeyIndex].description}`,
                passing: !numFailing,
                nodes: [keys[nodeKeyIndex].name]
            };
        }
    });
    return Object.values(controls);
};
const variables = {
    where: queryService.objectToWhereClause({
        Standard: 'CIS'
    }),
    groupBy: [entityTypes.STANDARD, entityTypes.CONTROL, entityTypes.NODE]
};

const CISControls = ({ className, selectedRowId, onRowClick }) => (
    <List
        className={className}
        query={QUERY}
        headerText="CIS Controls"
        variables={variables}
        entityType={entityTypes.CONTROL}
        tableColumns={tableColumns}
        createTableRows={createTableRows}
        onRowClick={onRowClick}
        selectedRowId={selectedRowId}
        idAttribute="id"
    />
);

CISControls.propTypes = {
    className: PropTypes.string,
    selectedRowId: PropTypes.string,
    onRowClick: PropTypes.func.isRequired
};

CISControls.defaultProps = {
    className: '',
    selectedRowId: null
};

export default CISControls;
