import React from 'react';
import PropTypes from 'prop-types';

import * as Icon from 'react-feather';
import find from 'lodash/find';
import { Tooltip } from 'react-tippy';
import { sortValue, sortDate } from 'sorters/sorters';
import dateFns from 'date-fns';
import dateTimeFormat from 'constants/dateTimeFormat';

function DeploymentNameColumn({ original }) {
    const isSuspicious = find(original.whitelistStatuses, {
        anomalousProcessesExecuted: true
    });
    return (
        <div className="flex">
            <span className="pr-1">
                {isSuspicious && (
                    <Tooltip
                        useContext
                        position="top"
                        trigger="mouseenter"
                        arrow
                        html={<span className="text-sm">Abnormal processes discovered</span>}
                        unmountHTMLWhenHide
                    >
                        <Icon.Circle className="h-2 w-2 text-alert-400" fill="#ffebf1" />
                    </Tooltip>
                )}
                {!isSuspicious && <Icon.Circle className="h-2 w-2" />}
            </span>
            {original.deployment.name}
        </div>
    );
}

DeploymentNameColumn.propTypes = {
    original: PropTypes.shape({
        deployment: PropTypes.shape({
            name: PropTypes.string.isRequired
        }).isRequired,
        whitelistStatuses: PropTypes.arrayOf(PropTypes.object).isRequired
    }).isRequired
};

const columns = [
    {
        Header: 'Name',
        accessor: 'deployment.name',
        searchField: 'Deployment',
        // eslint-disable-next-line react/prop-types
        Cell: DeploymentNameColumn
    },
    {
        Header: 'Updated',
        accessor: 'deployment.updatedAt',
        searchField: 'Updated At',
        // eslint-disable-next-line react/prop-types
        Cell: ({ value }) => <span>{dateFns.format(value, dateTimeFormat)}</span>,
        sortMethod: sortDate
    },
    {
        Header: 'Cluster',
        searchField: 'Cluster',
        accessor: 'deployment.cluster'
    },
    {
        Header: 'Namespace',
        searchField: 'Namespace',
        accessor: 'deployment.namespace'
    },
    {
        Header: 'Priority',
        searchField: 'Priority',
        accessor: 'deployment.priority',
        sortMethod: sortValue
    }
];

export default columns;
