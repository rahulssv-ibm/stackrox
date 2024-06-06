import { gql, useQuery } from '@apollo/client';
import { getPaginationParams } from 'utils/searchUtils';
import { ClientPagination } from 'services/types';
import { getRegexScopedQueryString } from '../../utils/searchUtils';
import { QuerySearchFilter } from '../../types';

const nodeListQuery = gql`
    query getNodes($query: String, $pagination: Pagination) {
        nodes(query: $query, pagination: $pagination) {
            id
            name
            nodeCVECountBySeverity {
                critical {
                    total
                }
                important {
                    total
                }
                moderate {
                    total
                }
                low {
                    total
                }
            }
            cluster {
                name
            }
            operatingSystem
            scanTime
        }
    }
`;

// TODO - Verify these types once the BE is implemented
type Node = {
    id: string;
    name: string;
    nodeCVECountBySeverity: {
        critical: {
            total: number;
        };
        important: {
            total: number;
        };
        moderate: {
            total: number;
        };
        low: {
            total: number;
        };
    };
    cluster: {
        name: string;
    };
    // TODO Swap this to the osImage field
    operatingSystem: string;
    scanTime: string;
};

export default function useNodes({
    querySearchFilter,
    page,
    perPage,
    sortOption,
}: { querySearchFilter: QuerySearchFilter } & ClientPagination) {
    return useQuery<{ nodes: Node[] }>(nodeListQuery, {
        variables: {
            query: getRegexScopedQueryString(querySearchFilter),
            pagination: getPaginationParams({ page, perPage, sortOption }),
        },
    });
}
