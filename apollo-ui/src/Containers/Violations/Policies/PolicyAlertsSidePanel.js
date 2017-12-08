import React, { Component } from 'react';
import Table from 'Components/Table';
import ReactModal from 'react-modal';
import * as Icon from 'react-feather';
import emitter from 'emitter';
import axios from 'axios';

class MainSidePanel extends Component {
    constructor(props) {
        super(props);

        this.state = {
            showPanel: false,
            policy: {},
            table: {
                columns: [
                    { key: 'deployment.name', label: 'Deployment' },
                    { key: 'time', label: 'Time' },
                    { key: 'deployment.image.registry', label: 'Registry' }
                ],
                rows: []
            },
            modal: {
                showModal: false,
                data: {}
            }
        }

        this.hidePanel = this.hidePanel.bind(this);
        this.handleOpenModal = this.handleOpenModal.bind(this);
        this.handleCloseModal = this.handleCloseModal.bind(this);
    }

    getAlerts(data) {
        if (!data) {
            this.setState({ showPanel: false, data: {}, alerts:[] }); 
            return;
        }
        this.clearData();
        axios.get('/v1/alerts', {
            params: {
                policy_name: data.name
            }
        }).then((response) => {
            if(!response.data || !response.data.alerts.length) return;
            var table = this.state.table;
            table.rows = response.data.alerts.map((alert) => {
                alert.severity = alert.severity.replace('_', ' ');
                return alert;
            });
            this.setState({ data: data, table: table });
        }).catch((error) => {
            this.setState({ data: {}, alerts: [] });
        });
    }

    displayHeader() {
        if(!this.state.data) return "";
        return (
            <div className="flex flex-row">
                {/*<span className="font-semibold">Policy: </span> 
                <span>{this.state.data.name}</span>*/}
                <span className="flex flex-1">Alerts for "{this.state.data.name}"</span>
                <Icon.X className="cursor-pointer" color="white" size={20} onClick={() => { this.hidePanel() }} />
            </div>
        );
    }

    displayModalHeader() {
        if (this.state.modal.data === {} || !this.state.modal.data || !this.state.modal.data.deployment) return "";
        return (
            <header className="flex flex-row w-full p-4 border-b border-grey-light font-bold">
                <span className="flex flex-1">{this.state.modal.data.deployment.name} ({this.state.modal.data.deployment.id})</span>
                <Icon.X className="cursor-pointer" color="black" size={20} onClick={() => { this.handleCloseModal() }} />
            </header>
        );
    }

    displayModalBody() {
        if (this.state.modal.data === {} || !this.state.modal.data || !this.state.modal.data.deployment) return "";
        return (
            <div className="flex flex-row bg-grey-lighter overflow-auto">
                <div className="flex flex-col w-1/2 border-r border-grey-light">
                    <div className="bg-white m-4">
                        <header className="w-full p-4 border-b border-grey-light font-bold">Alert Summary</header>
                        <div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Description:</span> {this.state.modal.data.policy.imagePolicy.description}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Deployment ID:</span> {this.state.modal.data.deployment.id}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Severity:</span> {this.state.modal.data.severity}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Time:</span> {this.state.modal.data.time}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Type:</span> {this.state.modal.data.deployment.type}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Last Updated:</span> {this.state.modal.data.deployment.updatedAt}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Version:</span> {this.state.modal.data.deployment.version}</div>
                        </div>
                    </div>
                    <div className="bg-white m-4">
                        <header className="w-full p-4 border-b border-grey-light font-bold">Image Summary</header>
                        <div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Registry:</span> {this.state.modal.data.deployment.image.registry}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Remote:</span> {this.state.modal.data.deployment.image.remote}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">SHA:</span> {this.state.modal.data.deployment.image.sha}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Metadata:</span> {this.state.modal.data.deployment.image.metadata}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Scan:</span> {this.state.modal.data.deployment.image.scan}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Tag:</span> {this.state.modal.data.deployment.image.tag}</div>
                        </div>
                    </div>
                </div>
                <div className="flex flex-col w-1/2">
                    <div className="bg-white m-4">
                        <header className="w-full p-4 border-b border-grey-light font-bold">Policy Details</header>
                        <div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Name:</span> {this.state.modal.data.policy.name}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Description:</span> {this.state.modal.data.policy.imagePolicy.description}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Category:</span> {this.state.modal.data.policy.category}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Severity:</span> {this.state.modal.data.policy.imagePolicy.severity}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Disabled:</span> {String(this.state.modal.data.policy.imagePolicy.disabled)}</div>
                            <div className="py-2 px-4 text-xs truncate"><span className="font-bold">Scan Age Day:</span> {this.state.modal.data.policy.imagePolicy.scanAgeDays}</div>
                        </div>
                    </div>
                    <div className="bg-white m-4">
                        <header className="w-full p-4 border-b border-grey-light font-bold">Violations</header>
                        <div>
                            {this.state.modal.data.policy.violations.map((violation, i) => { return <div key={'policy-alerts-violation-' + i} className="py-2 px-4 text-xs break-words">{violation.message}</div>; }) }
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    clearData() {
        this.setState({ showPanel: true, alerts: [] });
    }

    hidePanel() {
        this.setState({ showPanel: false });
    }

    handleOpenModal(row) {
        var modal = this.state.modal;
        modal.showModal = true;
        modal.data = row;
        console.log(modal.data);
        this.setState({ modal: modal });
    }

    handleCloseModal() {
        var modal = this.state.modal;
        modal.showModal = false;
        this.setState({ modal: modal });
    }

    render() {
        return (
            <aside className={"pin-r h-full bg-white border-l border-grey-light " + ((this.state.showPanel) ? 'flex flex-col' : 'hidden')}>
                <div className="p-4 bg-black text-white border-b border-grey-light w-full">{this.displayHeader()}</div>
                <Table columns={this.state.table.columns} rows={this.state.table.rows} onRowClick={this.handleOpenModal.bind(this)}></Table>
                <ReactModal
                    isOpen={this.state.modal.showModal}
                    onRequestClose={this.handleCloseModal}
                    contentLabel="Modal"
                    ariaHideApp={false}
                    overlayClassName="ReactModal__Overlay react-modal-overlay"
                    className="ReactModal__Content h-95 w-2/3 m-0-auto p0 bg-white border border-grey-light mt-4 mb-4 overflow-auto">
                    {this.displayModalHeader()}
                    {this.displayModalBody()}
                </ReactModal>
            </aside>
        );
    }

    componentDidMount() {
        // set up event listeners for this componenet
        this.tableRowSelectedListener = emitter.addListener('Table:row-selected', (data) => {
            this.getAlerts(data);
        });
    }

    componentWillUnmount() {
        // remove event listeners
        this.tableRowSelectedListener.remove();
    }

}

export default MainSidePanel;
