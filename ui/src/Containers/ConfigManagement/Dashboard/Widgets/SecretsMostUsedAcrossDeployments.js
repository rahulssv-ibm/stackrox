import React from 'react';
import Loader from 'Components/Loader';
import { Link, withRouter } from 'react-router-dom';
import { Tooltip } from 'react-tippy';
import URLService from 'modules/URLService';
import gql from 'graphql-tag';
import entityTypes from 'constants/entityTypes';
import Query from 'Components/ThrowingQuery';
import Widget from 'Components/Widget';
import pluralize from 'pluralize';
import ReactRouterPropTypes from 'react-router-prop-types';
import dateFns from 'date-fns';

const QUERY = gql`
    query secrets {
        secrets {
            id
            name
            clusterName
            namespace
            files {
                name
                type
                metadata {
                    __typename
                    ... on Cert {
                        endDate
                        startDate
                    }
                    ... on ImagePullSecret {
                        registries {
                            name
                            username
                        }
                    }
                }
            }
            deployments {
                id
            }
        }
    }
`;

const getCertificateStatus = files => {
    let status = 'no';
    files.forEach(file => {
        if (file.metadata) {
            const { startDate, endDate } = file.metadata;
            if (!startDate && !endDate) return;

            const isUpcoming = dateFns.isAfter(new Date(startDate), new Date());
            const hasExpired = dateFns.isAfter(new Date(endDate), new Date());

            if (isUpcoming) {
                status = 'upcoming';
            }

            if (hasExpired) {
                status = 'expired';
            }
        }
    });

    return `has ${status} certs`;
};

const SecretsMostUsedAcrossDeployments = ({ match, location }) => {
    function processData(data) {
        if (!data || !data.secrets) return [];

        return data.secrets
            .filter(secret => secret.deployments.length !== 0)
            .sort((a, b) => b.deployments.length - a.deployments.length)
            .slice(0, 10);
    }
    return (
        <Query query={QUERY}>
            {({ loading, data }) => {
                let contents = <Loader />;
                const viewAllURL = URLService.getURL(match, location)
                    .base(entityTypes.SECRET)
                    .url();

                const viewAllLink = (
                    <Link to={viewAllURL} className="no-underline">
                        <button className="btn-sm btn-base" type="button">
                            View All
                        </button>
                    </Link>
                );

                if (!loading && data) {
                    const results = processData(data);

                    contents = (
                        <ul
                            className="list-reset w-full columns-2 columns-gap-0"
                            style={{ columnRule: '1px solid var(--base-300)' }}
                        >
                            {results.map((item, index) => {
                                const linkTo = URLService.getURL(match, location)
                                    .base(entityTypes.SECRET)
                                    .push(item.id)
                                    .url();
                                return (
                                    <Link
                                        key={`${item.id}-${index}`}
                                        to={linkTo}
                                        className={`no-underline text-base-600 hover:bg-base-400 hover:bg-base-200 inline-block border-base-300 w-full ${
                                            index !== 4 || index !== 9 ? 'border-b' : ''
                                        }`}
                                    >
                                        <li key={`${item.name}-${index}`}>
                                            <div className="flex flex-row">
                                                <div className="self-center text-2xl tracking-widest pl-4 pr-4">
                                                    {index + 1}
                                                </div>
                                                <div className="flex flex-col truncate pr-4 pb-4 pt-4 text-sm">
                                                    <span className="text-xs pb-1 italic text-base-500">
                                                        {item.clusterName}/{item.namespace}
                                                    </span>
                                                    <span className="pb-2">{item.name}</span>
                                                    <Tooltip
                                                        position="top"
                                                        trigger="mouseenter"
                                                        animation="none"
                                                        duration={0}
                                                        arrow
                                                        distance={20}
                                                        html={
                                                            <div className="text-sm italic">
                                                                {`${
                                                                    item.deployments.length
                                                                } ${pluralize(
                                                                    'Deployment',
                                                                    item.deployments.length
                                                                )}, `}
                                                                {getCertificateStatus(item.files)}
                                                            </div>
                                                        }
                                                        unmountHTMLWhenHide
                                                    >
                                                        {item.deployments.length > 0 && (
                                                            <div className="truncate italic">
                                                                {`${
                                                                    item.deployments.length
                                                                } ${pluralize(
                                                                    'Deployment',
                                                                    item.deployments.length
                                                                )}, `}
                                                                {getCertificateStatus(item.files)}
                                                            </div>
                                                        )}
                                                    </Tooltip>
                                                </div>
                                            </div>
                                        </li>
                                    </Link>
                                );
                            })}
                        </ul>
                    );
                }
                return (
                    <Widget
                        className="s-2 overflow-hidden pdf-page"
                        header="Secrets Most Used Across Deployments"
                        headerComponents={viewAllLink}
                    >
                        {contents}
                    </Widget>
                );
            }}
        </Query>
    );
};

SecretsMostUsedAcrossDeployments.propTypes = {
    match: ReactRouterPropTypes.match.isRequired,
    location: ReactRouterPropTypes.location.isRequired
};

export default withRouter(SecretsMostUsedAcrossDeployments);
