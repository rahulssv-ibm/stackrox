import React from 'react';
import { Button } from '@patternfly/react-core';
import { SearchFilter } from 'types/search';

export type SnoozeCveToggleButtonProps = {
    searchFilter: SearchFilter;
    setSearchFilter: (searchFilter: SearchFilter) => void;
};

function SnoozeCveToggleButton({ searchFilter, setSearchFilter }: SnoozeCveToggleButtonProps) {
    const isSnoozeFilterActive = searchFilter['CVE Snoozed']?.[0] === 'true';
    const buttonText = isSnoozeFilterActive ? 'Show observed CVEs' : 'Show snoozed CVEs';

    function toggleSnoozeFilter() {
        const nextFilter = { ...searchFilter };
        if (isSnoozeFilterActive) {
            delete nextFilter['CVE Snoozed'];
        } else {
            nextFilter['CVE Snoozed'] = ['true'];
        }
        setSearchFilter(nextFilter);
    }

    return (
        <Button variant="secondary" onClick={toggleSnoozeFilter}>
            {buttonText}
        </Button>
    );
}

export default SnoozeCveToggleButton;
