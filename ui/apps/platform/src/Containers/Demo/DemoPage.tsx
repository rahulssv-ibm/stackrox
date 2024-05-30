import React from 'react';
import {
    CodeBlock,
    CodeBlockCode,
    Divider,
    Flex,
    FlexItem,
    PageSection,
    Title,
} from '@patternfly/react-core';

import PageTitle from 'Components/PageTitle';
import CompoundSearchFilter from 'Components/CompoundSearchFilter/components/CompoundSearchFilter';
import { compoundSearchFilter } from 'Components/CompoundSearchFilter/types';
import useURLSearch from 'hooks/useURLSearch';

function DemoPage() {
    const { searchFilter, setSearchFilter } = useURLSearch();

    const historyHeader = '//// Search History ////';
    const historyText = `\n\n${Object.keys(searchFilter)
        .map((searchKey) => {
            const searchValue = searchFilter[searchKey];
            return `${searchKey}: ${Array.isArray(searchValue) ? searchValue.join(',') : searchValue}`;
        })
        .join('\n')}`;
    const content = `${historyHeader}${history.length !== 0 ? historyText : ''}`;

    return (
        <>
            <PageTitle title="Demo - Advanced Filters" />
            <PageSection variant="light">
                <Flex>
                    <Flex direction={{ default: 'column' }} flex={{ default: 'flex_1' }}>
                        <Title headingLevel="h1">Demo - Advanced Filters</Title>
                        <FlexItem>
                            This section will demo the capabilities of advanced filters. NOT A REAL
                            PAGE
                        </FlexItem>
                    </Flex>
                </Flex>
            </PageSection>
            <Divider component="div" />
            <PageSection>
                <PageSection variant="light">
                    <Flex
                        direction={{ default: 'column' }}
                        spaceItems={{ default: 'spaceItemsLg' }}
                    >
                        <CompoundSearchFilter
                            config={compoundSearchFilter}
                            onSearch={(searchKey, searchValue) => {
                                setSearchFilter({
                                    ...searchFilter,
                                    [searchKey]: searchValue,
                                });
                            }}
                        />
                        <CodeBlock>
                            <CodeBlockCode id="code-content">{content}</CodeBlockCode>
                        </CodeBlock>
                    </Flex>
                </PageSection>
            </PageSection>
        </>
    );
}

export default DemoPage;
