import { visit } from '../../../helpers/visit';
import {
    fillAndSubmitExceptionForm,
    selectSingleCveForException,
    verifyExceptionConfirmationDetails,
    verifySelectedCvesInModal,
    visitWorkloadCveOverview,
} from '../workloadCves/WorkloadCves.helpers';
import { selectors as workloadCVESelectors } from '../workloadCves/WorkloadCves.selectors';
import { selectors as vulnSelectors } from '../vulnerabilities.selectors';

const basePath = '/main/vulnerabilities/exception-management';
export const pendingRequestsPath = `${basePath}/pending-requests`;
export const approvedDeferralsPath = `${basePath}/approved-deferrals`;
export const approvedFalsePositivesPath = `${basePath}/approved-false-positives`;
export const deniedRequestsPath = `${basePath}/denied-requests`;

export function visitExceptionManagement() {
    visit(pendingRequestsPath);

    cy.get('h1:contains("Exception management")');
    cy.location('pathname').should('eq', pendingRequestsPath);
}

export function deferAndVisitRequestDetails({
    comment,
    expiry,
    scope,
}: {
    comment: string;
    expiry: string;
    scope: string;
}) {
    visitWorkloadCveOverview();
    cy.get(vulnSelectors.clearFiltersButton).click(); // Note: This is a workaround to prevent a lack of CVE data from causing the test to fail in CI

    // defer a single cve
    selectSingleCveForException('DEFERRAL').then((cveName) => {
        verifySelectedCvesInModal([cveName]);
        fillAndSubmitExceptionForm({
            comment,
            expiryLabel: expiry,
        });
        verifyExceptionConfirmationDetails({
            expectedAction: 'Deferral',
            cves: [cveName],
            scope,
            expiry,
        });
        cy.get(workloadCVESelectors.copyToClipboardButton).click();
        cy.get(workloadCVESelectors.copyToClipboardTooltipText).contains('Copied');
        // @TODO: Can make this into a custom cypress command (ie. getClipboardText)
        cy.window()
            .then((win) => {
                return win.navigator.clipboard.readText();
            })
            .then((url) => {
                visit(url);
            });
    });
}

// @TODO: We could possibly just use a single function for deferral/false positive
export function markFalsePositiveAndVisitRequestDetails({
    comment,
    scope,
}: {
    comment: string;
    scope: string;
}) {
    visitWorkloadCveOverview();
    cy.get(vulnSelectors.clearFiltersButton).click(); // Note: This is a workaround to prevent a lack of CVE data from causing the test to fail in CI

    // mark a single cve as false positive
    selectSingleCveForException('FALSE_POSITIVE').then((cveName) => {
        verifySelectedCvesInModal([cveName]);
        fillAndSubmitExceptionForm({ comment });
        verifyExceptionConfirmationDetails({
            expectedAction: 'False positive',
            cves: [cveName],
            scope,
        });
        cy.get(workloadCVESelectors.copyToClipboardButton).click();
        cy.get(workloadCVESelectors.copyToClipboardTooltipText).contains('Copied');
        // @TODO: Can make this into a custom cypress command (ie. getClipboardText)
        cy.window()
            .then((win) => {
                return win.navigator.clipboard.readText();
            })
            .then((url) => {
                visit(url);
            });
    });
}
