import { useEffect, useState } from 'react';

import { fetchNetworkBaselines } from 'services/NetworkService';
import { filterLabels } from 'constants/networkFilterModes';
import { FlattenedNetworkBaseline } from 'Containers/Network/networkTypes';
import { networkFlowStatus, nodeTypes } from 'constants/networkGraph';

type Result = { isLoading: boolean; data: FlattenedNetworkBaseline[]; error: string | null };

export function getPeerEntityName(peer): string {
    switch (peer.entity.info.type) {
        case nodeTypes.EXTERNAL_ENTITIES:
            return 'External Entities';
        case nodeTypes.CIDR_BLOCK:
            return peer.entity.info.externalSource.name;
        default:
            return peer.entity.info.deployment.name;
    }
}

/*
 * This hook does an API call to the baseline status API to get the baseline status
 * of the supplied peers
 */
function useFetchNetworkBaselines({ deploymentId, filterState }): Result {
    const [result, setResult] = useState<Result>({ data: [], error: null, isLoading: true });

    useEffect(() => {
        const networkBaselinesPromise = fetchNetworkBaselines({ deploymentId });

        networkBaselinesPromise
            .then((response) => {
                const { namespace, peers } = response;
                const data = peers.reduce((acc, currPeer) => {
                    currPeer.properties.forEach((property) => {
                        const name = getPeerEntityName(currPeer);
                        const peer = {
                            entity: {
                                id: currPeer.entity.info.id,
                                type: currPeer.entity.info.type,
                                name,
                                namespace,
                            },
                            ingress: property.ingress,
                            port: property.port,
                            protocol: property.protocol,
                            state: filterLabels[filterState],
                        };
                        acc.push({
                            peer,
                            status: networkFlowStatus.BASELINE,
                        });
                    });
                    return acc;
                }, []);
                setResult({ data, error: null, isLoading: false });
            })
            .catch((error) => {
                setResult({ data: [], error, isLoading: false });
            });
    }, [deploymentId, filterState]);

    return result;
}

export default useFetchNetworkBaselines;
