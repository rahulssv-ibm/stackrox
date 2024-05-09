import React, { ReactElement } from 'react';
import { TextInput, PageSection, Form, Checkbox } from '@patternfly/react-core';
import * as yup from 'yup';

import { ImageIntegrationBase } from 'services/ImageIntegrationsService';

import usePageState from 'Containers/Integrations/hooks/usePageState';
import FormMessage from 'Components/PatternFly/FormMessage';
import FormTestButton from 'Components/PatternFly/FormTestButton';
import FormSaveButton from 'Components/PatternFly/FormSaveButton';
import FormCancelButton from 'Components/PatternFly/FormCancelButton';
import useCentralCapabilities from 'hooks/useCentralCapabilities';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormLabelGroup from '../FormLabelGroup';

export type EcrIntegration = {
    categories: 'REGISTRY'[];
    ecr: {
        registryId: string;
        endpoint: string;
        region: string;
        useIam: boolean;
        accessKeyId: string;
        secretAccessKey: string;
        useAssumeRole: boolean;
        assumeRoleId: string;
        assumeRoleExternalId: string;
    };
    type: 'ecr';
} & ImageIntegrationBase;

export type EcrIntegrationFormValues = {
    config: EcrIntegration;
    updatePassword: boolean;
};

export const validationSchema = yup.object().shape({
    config: yup.object().shape({
        name: yup.string().trim().required('An integration name is required'),
        categories: yup
            .array()
            .of(yup.string().trim().oneOf(['REGISTRY']))
            .min(1, 'Must have at least one type selected')
            .required('A category is required'),
        ecr: yup.object().shape({
            registryId: yup.string().trim().required('A 12-digit AWS ID is required'),
            endpoint: yup.string().trim(),
            region: yup.string().trim().required('An AWS region is required'),
            useIam: yup.bool(),
            accessKeyId: yup.string().when('useIam', {
                is: false,
                then: (accessKeyIdSchema) =>
                    accessKeyIdSchema.test(
                        'acessKeyId-test',
                        'An access key ID is required',
                        (value, context: yup.TestContext) => {
                            const requirePasswordField =
                                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                                // @ts-ignore
                                context?.from[2]?.value?.updatePassword || false;

                            if (!requirePasswordField) {
                                return true;
                            }

                            const trimmedValue = value?.trim();
                            return !!trimmedValue;
                        }
                    ),
            }),
            secretAccessKey: yup.string().when('useIam', {
                is: false,
                then: (secretAccessKeySchema) =>
                    secretAccessKeySchema.test(
                        'secretAccessKey-test',
                        'A secret access key is required',
                        (value, context: yup.TestContext) => {
                            const requirePasswordField =
                                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                                // @ts-ignore
                                context?.from[2]?.value?.updatePassword || false;

                            if (!requirePasswordField) {
                                return true;
                            }

                            const trimmedValue = value?.trim();
                            return !!trimmedValue;
                        }
                    ),
            }),
            useAssumeRole: yup.bool(),
            assumeRoleId: yup.string().when('useAssumeRole', {
                is: true,
                then: (assumeRoleIdSchema) =>
                    assumeRoleIdSchema.trim().required('A Role ID is required'),
            }),
            assumeRoleExternalId: yup.string().when('useAssumeRole', {
                is: true,
                then: (assumeRoleExternalIdSchema) => assumeRoleExternalIdSchema.trim(),
            }),
        }),
        skipTestIntegration: yup.bool(),
        type: yup.string().matches(/ecr/),
    }),
    updatePassword: yup.bool(),
});

export const defaultValues: EcrIntegrationFormValues = {
    config: {
        id: '',
        name: '',
        categories: ['REGISTRY'],
        ecr: {
            registryId: '',
            endpoint: '',
            region: '',
            useIam: false,
            accessKeyId: '',
            secretAccessKey: '',
            useAssumeRole: false,
            assumeRoleId: '',
            assumeRoleExternalId: '',
        },
        autogenerated: false,
        clusterId: '',
        clusters: [],
        skipTestIntegration: false,
        type: 'ecr',
    },
    updatePassword: true,
};

function EcrIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<EcrIntegration>): ReactElement {
    const formInitialValues = { ...defaultValues, ...initialValues };

    const { isCentralCapabilityAvailable } = useCentralCapabilities();
    const canUseContainerIamRoleForEcr = isCentralCapabilityAvailable(
        'centralScanningCanUseContainerIamRoleForEcr'
    );

    if (initialValues) {
        formInitialValues.config = { ...formInitialValues.config, ...initialValues };
        // We want to clear the password because backend returns '******' to represent that there
        // are currently stored credentials
        formInitialValues.config.ecr.accessKeyId = '';
        formInitialValues.config.ecr.secretAccessKey = '';

        // Don't assume user wants to change password; that has caused confusing UX.
        formInitialValues.updatePassword = false;
    }
    const {
        values,
        touched,
        errors,
        dirty,
        isValid,
        setFieldValue,
        handleBlur,
        isSubmitting,
        isTesting,
        onSave,
        onTest,
        onCancel,
        message,
    } = useIntegrationForm<EcrIntegrationFormValues>({
        initialValues: formInitialValues,
        validationSchema,
    });
    const { isCreating } = usePageState();

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    function onUpdateCredentialsChange(value, event) {
        setFieldValue('config.ecr.accessKeyId', '');
        setFieldValue('config.ecr.secretAccessKey', '');
        return setFieldValue(event.target.id, value);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                <FormMessage message={message} />
                <Form isWidthLimited>
                    <FormLabelGroup
                        label="Integration name"
                        isRequired
                        fieldId="config.name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="config.name"
                            value={values.config.name}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="12-digit AWS ID"
                        isRequired
                        fieldId="config.ecr.registryId"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            type="text"
                            id="config.ecr.registryId"
                            value={values.config.ecr.registryId}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Endpoint"
                        fieldId="config.ecr.endpoint"
                        touched={touched}
                        errors={errors}
                        helperText={
                            values.config.ecr.useAssumeRole
                                ? 'Endpoint disabled when AssumeRole is set'
                                : ''
                        }
                    >
                        <TextInput
                            type="text"
                            id="config.ecr.endpoint"
                            value={values.config.ecr.endpoint}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable || values.config.ecr.useAssumeRole}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Region"
                        isRequired
                        fieldId="config.ecr.region"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="config.ecr.region"
                            value={values.config.ecr.region}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    {!isCreating && isEditable && (
                        <FormLabelGroup
                            fieldId="updatePassword"
                            helperText="Enable this option to replace currently stored credentials (if any)"
                            errors={errors}
                        >
                            <Checkbox
                                label="Update stored credentials"
                                id="updatePassword"
                                isChecked={values.updatePassword}
                                onChange={(event, value) => onUpdateCredentialsChange(value, event)}
                                onBlur={handleBlur}
                                isDisabled={!isEditable}
                            />
                        </FormLabelGroup>
                    )}
                    {canUseContainerIamRoleForEcr && (
                        <FormLabelGroup
                            fieldId="config.ecr.useIam"
                            touched={touched}
                            errors={errors}
                        >
                            <Checkbox
                                label="Use container IAM role"
                                id="config.ecr.useIam"
                                aria-label="use container iam role"
                                isChecked={values.config.ecr.useIam}
                                onChange={(event, value) => onChange(value, event)}
                                onBlur={handleBlur}
                                isDisabled={!isEditable}
                            />
                        </FormLabelGroup>
                    )}
                    {!values.config.ecr.useIam && (
                        <>
                            <FormLabelGroup
                                isRequired={values.updatePassword}
                                label="Access key ID"
                                fieldId="config.ecr.accessKeyId"
                                touched={touched}
                                errors={errors}
                            >
                                <TextInput
                                    isRequired={values.updatePassword}
                                    type="password"
                                    id="config.ecr.accessKeyId"
                                    value={values.config.ecr.accessKeyId}
                                    onChange={(event, value) => onChange(value, event)}
                                    onBlur={handleBlur}
                                    isDisabled={!isEditable || !values.updatePassword}
                                />
                            </FormLabelGroup>
                            <FormLabelGroup
                                isRequired={values.updatePassword}
                                label="Secret access key"
                                fieldId="config.ecr.secretAccessKey"
                                touched={touched}
                                errors={errors}
                            >
                                <TextInput
                                    isRequired={values.updatePassword}
                                    type="password"
                                    id="config.ecr.secretAccessKey"
                                    value={values.config.ecr.secretAccessKey}
                                    onChange={(event, value) => onChange(value, event)}
                                    onBlur={handleBlur}
                                    isDisabled={!isEditable || !values.updatePassword}
                                />
                            </FormLabelGroup>
                        </>
                    )}
                    <FormLabelGroup
                        fieldId="config.ecr.useAssumeRole"
                        touched={touched}
                        errors={errors}
                        helperText={
                            !isEditable || values.config.ecr.endpoint !== ''
                                ? 'AssumeRole disabled when Endpoint is set'
                                : ''
                        }
                    >
                        <Checkbox
                            label="Use AssumeRole"
                            id="config.ecr.useAssumeRole"
                            aria-label="use assume role"
                            isChecked={values.config.ecr.useAssumeRole}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable || values.config.ecr.endpoint !== ''}
                        />
                    </FormLabelGroup>
                    {values.config.ecr.useAssumeRole && (
                        <>
                            <FormLabelGroup
                                label="AssumeRole ID"
                                fieldId="config.ecr.assumeRoleId"
                                touched={touched}
                                errors={errors}
                            >
                                <TextInput
                                    type="text"
                                    id="config.ecr.assumeRoleId"
                                    value={values.config.ecr.assumeRoleId}
                                    onChange={(event, value) => onChange(value, event)}
                                    onBlur={handleBlur}
                                    isDisabled={!isEditable}
                                />
                            </FormLabelGroup>
                            <FormLabelGroup
                                label="AssumeRole External ID"
                                fieldId="config.ecr.assumeRoleExternalId"
                                touched={touched}
                                errors={errors}
                            >
                                <TextInput
                                    type="text"
                                    id="config.ecr.assumeRoleExternalId"
                                    value={values.config.ecr.assumeRoleExternalId}
                                    onChange={(event, value) => onChange(value, event)}
                                    onBlur={handleBlur}
                                    isDisabled={!isEditable}
                                />
                            </FormLabelGroup>
                        </>
                    )}
                    <FormLabelGroup
                        fieldId="config.skipTestIntegration"
                        touched={touched}
                        errors={errors}
                    >
                        <Checkbox
                            label="Create integration without testing"
                            id="config.skipTestIntegration"
                            aria-label="skip test integration"
                            isChecked={values.config.skipTestIntegration}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                </Form>
            </PageSection>
            {isEditable && (
                <IntegrationFormActions>
                    <FormSaveButton
                        onSave={onSave}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!dirty || !isValid}
                    >
                        Save
                    </FormSaveButton>
                    <FormTestButton
                        onTest={onTest}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!isValid}
                    >
                        Test
                    </FormTestButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default EcrIntegrationForm;
