import React from 'react';
import { Toolbar, ToolbarGroup, ToolbarContent, ToolbarItem } from '@patternfly/react-core';

import useURLSearch from 'hooks/useURLSearch';
import { getFilteredConfig } from 'Components/CompoundSearchFilter/utils/searchFilterConfig';
import {
    makeFilterChipDescriptors,
    onURLSearch,
} from 'Components/CompoundSearchFilter/utils/utils';
import {
    OnSearchPayload,
    clusterSearchFilterConfig,
    deploymentSearchFilterConfig,
    namespaceSearchFilterConfig,
} from 'Components/CompoundSearchFilter/types';
import SearchFilterChips from 'Components/PatternFly/SearchFilterChips';
import CompoundSearchFilter from 'Components/CompoundSearchFilter/components/CompoundSearchFilter';

type ViolationsTableSearchFilterProps = {};

const searchFilterConfig = {
    Cluster: getFilteredConfig(clusterSearchFilterConfig, ['Name']),
    Namespace: getFilteredConfig(namespaceSearchFilterConfig, ['Name']),
    Deployment: getFilteredConfig(deploymentSearchFilterConfig, ['Name']),
};

function ViolationsTableSearchFilter({}: ViolationsTableSearchFilterProps) {
    const { searchFilter, setSearchFilter } = useURLSearch();

    const filterChipGroupDescriptors = makeFilterChipDescriptors(searchFilterConfig);

    const onSearch = (payload: OnSearchPayload) => {
        onURLSearch(searchFilter, setSearchFilter, payload);
    };

    return (
        <Toolbar>
            <ToolbarContent>
                <ToolbarGroup className="pf-v5-u-w-100">
                    <ToolbarItem className="pf-v5-u-flex-1">
                        <CompoundSearchFilter
                            config={searchFilterConfig}
                            searchFilter={searchFilter}
                            onSearch={onSearch}
                        />
                    </ToolbarItem>
                </ToolbarGroup>
                <ToolbarGroup className="pf-v5-u-w-100">
                    <SearchFilterChips filterChipGroupDescriptors={filterChipGroupDescriptors} />
                </ToolbarGroup>
            </ToolbarContent>
        </Toolbar>
    );
}

export default ViolationsTableSearchFilter;
