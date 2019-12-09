import React from 'react';
import PropTypes from 'prop-types';
import { useTheme } from 'Containers/ThemeProvider';

import PageHeader from 'Components/PageHeader';
import ScanButton from 'Containers/Compliance/ScanButton';
import ExportButton from 'Components/ExportButton';
import entityTypes from 'constants/entityTypes';
import useCaseTypes from 'constants/useCaseTypes';
import Tile from './Tile';

const ComplianceDashboardHeader = ({ classes, bgStyle }) => {
    const { isDarkMode } = useTheme();
    const darkModeClasses = `${
        isDarkMode ? 'text-base-600 hover:bg-primary-200' : 'text-base-100 hover:bg-primary-800'
    }`;

    return (
        <PageHeader classes={classes} bgStyle={bgStyle} header="Compliance" subHeader="Dashboard">
            <div className="flex flex-1 justify-end">
                <Tile entityType={entityTypes.CLUSTER} position="first" />
                <Tile entityType={entityTypes.NAMESPACE} position="middle" />
                <Tile entityType={entityTypes.NODE} position="middle" />
                <Tile entityType={entityTypes.DEPLOYMENT} position="last" />
                <div className="ml-3 border-l border-base-100 mr-3 opacity-50" />
                <div className="flex">
                    <div className="flex items-center">
                        <ScanButton
                            className={`flex items-center justify-center border-2 border-primary-400 rounded px-2 uppercase lg:min-w-32 xl:min-w-43 h-10 ${darkModeClasses}`}
                            text="Scan environment"
                            textClass="hidden lg:block"
                            textCondensed="Scan all"
                            clusterId="*"
                            standardId="*"
                        />
                    </div>
                    <div className="flex items-center">
                        <ExportButton
                            className={`flex items-center border-2 border-primary-400 rounded p-2 uppercase h-10 ${darkModeClasses}`}
                            fileName="Compliance Dashboard Report"
                            textClass="hidden lg:block"
                            type="ALL"
                            page={useCaseTypes.COMPLIANCE}
                            pdfId="capture-dashboard"
                        />
                    </div>
                </div>
            </div>
        </PageHeader>
    );
};

ComplianceDashboardHeader.propTypes = {
    classes: PropTypes.string,
    bgStyle: PropTypes.shape({})
};

ComplianceDashboardHeader.defaultProps = {
    classes: null,
    bgStyle: null
};

export default ComplianceDashboardHeader;
