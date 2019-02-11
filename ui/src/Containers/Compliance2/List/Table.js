import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { standardTypes } from 'constants/entityTypes';
import pluralize from 'pluralize';
import { CLIENT_SIDE_SEARCH_OPTIONS as SEARCH_OPTIONS } from 'constants/searchOptions';

import Table from 'Components/Table';
import Panel from 'Components/Panel';
import Loader from 'Components/Loader';

import TablePagination from 'Components/TablePagination';
import TableGroup from 'Components/TableGroup';
import entityToColumns from 'constants/tableColumns';
import componentTypes from 'constants/componentTypes';
import AppQuery from 'Components/AppQuery';
import NoResultsMessage from 'Components/NoResultsMessage';

const standardTypeValues = Object.values(standardTypes);

class ListTable extends Component {
    static propTypes = {
        params: PropTypes.shape({}).isRequired,
        selectedRow: PropTypes.shape({}),
        updateSelectedRow: PropTypes.func.isRequired,
        pollInterval: PropTypes.number
    };

    static defaultProps = {
        selectedRow: null,
        pollInterval: 0
    };

    constructor(props) {
        super(props);
        this.state = {
            page: 0
        };
    }

    setTablePage = page => this.setState({ page });

    // This is a client-side implementation of filtering by the "Compliance State" Search Option
    filterByComplianceState = (data, params) => {
        const searchKey = SEARCH_OPTIONS.COMPLIANCE.STATE;
        if (!params.query[searchKey]) return data.results;
        const isPassing = params.query[searchKey].toLowerCase() === 'passing';
        const isFailing = params.query[searchKey].toLowerCase() === 'failing';
        const { results } = data;
        return results.filter(result => {
            const { id, name, ...standards } = result;
            return Object.values(standards).reduce((acc, strValue) => {
                const intValue = parseInt(strValue, 10); // strValue comes in the format "100.00%"
                if (isPassing) {
                    if (acc === false) return acc;
                    return intValue === 100;
                }
                if (isFailing) {
                    if (acc === true) return acc;
                    return intValue !== 100;
                }
                return acc;
            }, null);
        });
    };

    render() {
        const { params, selectedRow, updateSelectedRow, pollInterval } = this.props;
        const { page } = this.state;
        return (
            <AppQuery
                params={params}
                componentType={componentTypes.LIST_TABLE}
                pollInterval={pollInterval}
            >
                {({ loading, data }) => {
                    const isStandard = standardTypeValues.includes(params.entityType);
                    let tableData;
                    let contents = <Loader />;
                    let paginationComponent;
                    let headerText;
                    if (!loading || (data && data.results)) {
                        if (!data)
                            return (
                                <NoResultsMessage message="No compliance data available. Please run a scan." />
                            );
                        tableData = this.filterByComplianceState(data, params);
                        const total = tableData.length;
                        const groupedByText = params.query.groupBy
                            ? `across ${tableData.length} ${pluralize(params.query.groupBy, total)}`
                            : '';
                        headerText = isStandard
                            ? `${data.totalRows} ${pluralize(
                                  'control',
                                  data.totalRows
                              )} ${groupedByText}`
                            : `${total} ${pluralize(params.entityType, total)} ${groupedByText}`;
                        contents = isStandard ? (
                            <TableGroup
                                groups={tableData}
                                totalRows={data.totalRows}
                                tableColumns={entityToColumns[params.entityType]}
                                onRowClick={updateSelectedRow}
                                entityType={params.query.groupBy || 'control'}
                                idAttribute="id"
                                selectedRowId={selectedRow ? selectedRow.id : null}
                            />
                        ) : (
                            <Table
                                rows={tableData}
                                columns={entityToColumns[params.entityType]}
                                onRowClick={updateSelectedRow}
                                idAttribute="id"
                                selectedRowId={selectedRow ? selectedRow.id : null}
                                noDataText="No results found. Please refine your search."
                                page={page}
                                defaultSorted={[
                                    {
                                        id: 'name',
                                        desc: false
                                    }
                                ]}
                            />
                        );
                        paginationComponent = (
                            <TablePagination
                                page={page}
                                dataLength={total}
                                setPage={this.setTablePage}
                            />
                        );
                    }
                    return (
                        <Panel header={headerText} headerComponents={paginationComponent}>
                            {contents}
                        </Panel>
                    );
                }}
            </AppQuery>
        );
    }
}

export default ListTable;
