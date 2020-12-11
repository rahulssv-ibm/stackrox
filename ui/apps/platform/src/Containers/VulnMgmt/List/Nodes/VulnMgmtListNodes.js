/* eslint-disable react/display-name */
import React from 'react';
import { gql } from '@apollo/client';

import queryService from 'utils/queryService';
import TopCvssLabel from 'Components/TopCvssLabel';
import TableCellLink from 'Components/TableCellLink';
import CVEStackedPill from 'Components/CVEStackedPill';
import DateTimeField from 'Components/DateTimeField';
import { defaultHeaderClassName, defaultColumnClassName } from 'Components/Table';
import entityTypes from 'constants/entityTypes';
import { LIST_PAGE_SIZE } from 'constants/workflowPages.constants';
import WorkflowListPage from 'Containers/Workflow/WorkflowListPage';
import { NODE_LIST_FRAGMENT } from 'Containers/VulnMgmt/VulnMgmt.fragments';
import { workflowListPropTypes, workflowListDefaultProps } from 'constants/entityPageProps';
import removeEntityContextColumns from 'utils/tableUtils';
import { nodeSortFields } from 'constants/sortFields';

// TODO: need to get default node sort
export const defaultNodeSort = [
    {
        id: nodeSortFields.NODE,
        desc: false,
    },
];

// TODO: need to get node table columns
// Node | CVE (both total # / #non fixable) | Top CVSS | Scan Time | OS | Runtime | Node Status | Cluster | Risk Priority |
export function getImageTableColumns(workflowState) {
    const tableColumns = [
        {
            Header: 'Id',
            headerClassName: 'hidden',
            className: 'hidden',
            accessor: 'id',
        },
        {
            Header: `Node`,
            headerClassName: `w-1/6 ${defaultHeaderClassName}`,
            className: `w-1/6 word-break-all ${defaultColumnClassName}`,
            id: nodeSortFields.NODE,
            accessor: 'name',
            sortField: nodeSortFields.NODE,
        },
        {
            Header: `CVEs`,
            entityType: entityTypes.CVE,
            headerClassName: `w-1/6 ${defaultHeaderClassName}`,
            className: `w-1/6 ${defaultColumnClassName}`,
            Cell: ({ original, pdf }) => {
                const { vulnCounter, id, scan, notes } = original;

                const newState = workflowState.pushListItem(id).pushList(entityTypes.CVE);
                const url = newState.toUrl();
                const fixableUrl = newState.setSearch({ Fixable: true }).toUrl();

                return (
                    <CVEStackedPill
                        vulnCounter={vulnCounter}
                        url={url}
                        fixableUrl={fixableUrl}
                        hideLink={pdf}
                        imageNotes={notes}
                        scan={scan}
                    />
                );
            },
            id: nodeSortFields.CVE_COUNT,
            accessor: 'vulnCounter.all.total',
            sortField: nodeSortFields.CVE_COUNT,
        },
        {
            Header: `Top CVSS`,
            headerClassName: `w-1/12 text-center ${defaultHeaderClassName}`,
            className: `w-1/12 ${defaultColumnClassName}`,
            Cell: ({ original }) => {
                const { topVuln } = original;
                if (!topVuln) {
                    return (
                        <div className="mx-auto flex flex-col">
                            <span>–</span>
                        </div>
                    );
                }
                const { cvss, scoreVersion } = topVuln;
                return <TopCvssLabel cvss={cvss} version={scoreVersion} />;
            },
            id: nodeSortFields.TOP_CVSS,
            accessor: 'topVuln.cvss',
            sortField: nodeSortFields.TOP_CVSS,
        },
        {
            Header: `Cluster`,
            entityType: entityTypes.CLUSTER,
            headerClassName: `w-1/12 ${defaultHeaderClassName}`,
            className: `w-1/12 ${defaultColumnClassName}`,
            Cell: ({ original, pdf }) => {
                const { clusterName, clusterId, id } = original;
                const url = workflowState
                    .pushListItem(id)
                    .pushRelatedEntity(entityTypes.CLUSTER, clusterId)
                    .toUrl();
                return <TableCellLink pdf={pdf} url={url} text={clusterName} />;
            },
            id: nodeSortFields.CLUSTER,
            accessor: 'clusterName',
            sortField: nodeSortFields.CLUSTER,
        },
        {
            Header: `Operating System`,
            headerClassName: `w-1/12 ${defaultHeaderClassName}`,
            className: `w-1/12 ${defaultColumnClassName}`,
            id: nodeSortFields.OPERATING_SYSTEM,
            accessor: 'osImage',
            sortField: nodeSortFields.OPERATING_SYSTEM,
        },

        {
            Header: `Container Runtime`,
            headerClassName: `w-1/12 ${defaultHeaderClassName}`,
            className: `w-1/12 ${defaultColumnClassName}`,
            id: nodeSortFields.CONTAINER_RUNTIME,
            accessor: 'containerRuntimeVersion',
            sortField: nodeSortFields.CONTAINER_RUNTIME,
        },
        // {
        //     Header: 'Node Status',
        //     headerClassName: `w-1/12 ${nonSortableHeaderClassName}`,
        //     className: `w-1/12 ${defaultColumnClassName}`,
        //     Cell: ({ original, pdf }) => {
        //         const { deploymentCount } = original;
        //         const imageStatus = deploymentCount === 0 ? 'inactive' : 'active';
        //         return <StatusChip status={imageStatus} asString={pdf} />;
        //     },
        //     id: imageSortFields.IMAGE_STATUS,
        //     accessor: 'deploymentCount',
        //     sortField: imageSortFields.IMAGE_STATUS,
        //     sortable: false,
        // },
        {
            Header: `Join Time`,
            headerClassName: `w-1/12 ${defaultHeaderClassName}`,
            className: `w-1/12 ${defaultColumnClassName}`,
            Cell: ({ original, pdf }) => {
                const { joinedAt } = original;
                if (!joinedAt) {
                    return '–';
                }
                return <DateTimeField date={joinedAt} asString={pdf} />;
            },
            id: nodeSortFields.NODE_JOIN_TIME,
            accessor: 'joinedAt',
            sortField: nodeSortFields.NODE_JOIN_TIME,
        },
        {
            Header: `Scan Time`,
            headerClassName: `w-1/12 ${defaultHeaderClassName}`,
            className: `w-1/12 ${defaultColumnClassName}`,
            Cell: ({ original, pdf }) => {
                const { scan } = original;
                if (!scan) {
                    return '–';
                }
                return <DateTimeField date={scan.scanTime} asString={pdf} />;
            },
            id: nodeSortFields.SCAN_TIME,
            accessor: 'scan.scanTime',
            sortField: nodeSortFields.SCAN_TIME,
        },
        {
            Header: `Risk Priority`,
            headerClassName: `w-1/12 ${defaultHeaderClassName}`,
            className: `w-1/12 ${defaultColumnClassName}`,
            id: nodeSortFields.PRIORITY,
            accessor: 'priority',
            sortField: nodeSortFields.PRIORITY,
        },
    ];
    return removeEntityContextColumns(tableColumns, workflowState);
}

// TODO: set getNodes query to get real nodes list
const VulnMgmtNodes = ({ selectedRowId, search, sort, page, data, totalResults }) => {
    const query = gql`
        query getNodes($query: String, $pagination: Pagination) {
            results: nodes(query: $query, pagination: $pagination) {
                ...nodeFields
            }
            count: nodeCount(query: $query)
        }
        ${NODE_LIST_FRAGMENT}
    `;

    const tableSort = sort || defaultNodeSort;
    const queryOptions = {
        variables: {
            query: queryService.objectToWhereClause(search),
            pagination: queryService.getPagination(tableSort, page, LIST_PAGE_SIZE),
        },
    };

    return (
        <WorkflowListPage
            data={data}
            totalResults={totalResults}
            query={query}
            queryOptions={queryOptions}
            entityListType={entityTypes.NODE}
            getTableColumns={getImageTableColumns}
            selectedRowId={selectedRowId}
            search={search}
            sort={tableSort}
            page={page}
        />
    );
};

VulnMgmtNodes.propTypes = workflowListPropTypes;
VulnMgmtNodes.defaultProps = workflowListDefaultProps;

export default VulnMgmtNodes;
