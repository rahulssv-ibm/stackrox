import React, { useEffect, useState } from 'react';
import { Split } from '@patternfly/react-core';

import {
    CompoundSearchFilterConfig,
    SearchFilterAttributeName,
    SearchFilterEntityName,
} from '../types';
import { getDefaultAttribute, getDefaultEntity } from '../utils/utils';

import EntitySelector, { SelectedEntity } from './EntitySelector';
import AttributeSelector, { SelectedAttribute } from './AttributeSelector';
import CompoundSearchFilterInputField, { InputFieldValue } from './CompoundSearchFilterInputField';

export type CompoundSearchFilterProps = {
    config: Partial<CompoundSearchFilterConfig>;
    defaultEntity?: SearchFilterEntityName;
    defaultAttribute?: SearchFilterAttributeName;
};

function CompoundSearchFilter({
    config,
    defaultEntity,
    defaultAttribute,
}: CompoundSearchFilterProps) {
    const [selectedEntity, setSelectedEntity] = useState<SelectedEntity>(() => {
        if (defaultEntity) {
            return defaultEntity;
        }
        return getDefaultEntity(config);
    });

    const [selectedAttribute, setSelectedAttribute] = useState<SelectedAttribute>(() => {
        if (defaultAttribute) {
            return defaultAttribute;
        }
        const defaultEntity = getDefaultEntity(config);
        if (!defaultEntity) {
            return undefined;
        }
        return getDefaultAttribute(defaultEntity, config);
    });

    const [inputValue, setInputValue] = useState<InputFieldValue>('');

    useEffect(() => {
        if (defaultEntity) {
            setSelectedEntity(defaultEntity);
        }
    }, [defaultEntity]);

    useEffect(() => {
        if (defaultAttribute) {
            setSelectedAttribute(defaultAttribute);
        }
    }, [defaultAttribute]);

    return (
        <Split>
            <EntitySelector
                selectedEntity={selectedEntity}
                onChange={(value) => {
                    setSelectedEntity(value as SearchFilterEntityName);
                    const defaultAttribute = getDefaultAttribute(
                        value as SearchFilterEntityName,
                        config
                    );
                    setSelectedAttribute(defaultAttribute);
                    setInputValue('');
                }}
                config={config}
            />
            <AttributeSelector
                selectedEntity={selectedEntity}
                selectedAttribute={selectedAttribute}
                onChange={(value) => {
                    setSelectedAttribute(value as SearchFilterAttributeName);
                    setInputValue('');
                }}
                config={config}
            />
            <CompoundSearchFilterInputField
                selectedEntity={selectedEntity}
                selectedAttribute={selectedAttribute}
                value={inputValue}
                onChange={(value) => {
                    setInputValue(value);
                }}
                onSearch={(value) => {
                    // @TODO: Add search filter value to URL search
                    // eslint-disable-next-line no-alert
                    alert(value);
                }}
                config={config}
            />
        </Split>
    );
}

export default CompoundSearchFilter;
