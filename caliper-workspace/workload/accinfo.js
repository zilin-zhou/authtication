'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
    }

    // async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
    //     await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);

    //     for (let i=0; i<this.roundArguments.assets; i++) {
    //         const assetID = `${this.workerIndex}_${i}`;
    //         console.log(`Worker ${this.workerIndex}: Creating Info ${assetID}`);
    //         const request = {
    //             contractId: this.roundArguments.contractId,
    //             contractFunction: 'addAccInfo',
    //             invokerIdentity: 'User1',
    //             contractArguments: ['ZDU3MjQ0YWIxMGE3NmJlMTM5MzdkYjc2NjQ4ZjUwNGY0MzdlNmUwYTdjMzEwYjQwNzNkNmMwYmRjZTNjNGVmNQ==',
    //                                 'MTg4ODY0MzcyMDg3NDkzMTU5ODE2NTE1MzIyNDQ4OTEyODI1OTA4OTg0ODk0NzYzNDIzNDcxMjA0ODgxMDg2NjAwOTQxNDA3NDg4NzQ0MDk3OTg5ODg4ODY5NTI2ODY5ODU4NDEwMTM2ODc2OTExNTcwMzI3MjY2OTMyMDk3NjgzNjY5NDk2MjIyOTgxNjMwMDQ0OTk3NDc2NzU1MjU5NDY0MDcxMjc4Njg2MTg2NjMyMzU1MTE5MzYzOTY0OTM5NDMyNTQyNjg4MzcyNzczNTU4NTY1MzkwMzQ0MTA0NzUwNTM2ODUxNDM2NDYyMzE0NzkwODU2NDIyODU3MDExODc1NzU4NDg5ODA0MDQxMzUzODE0ODMzMjg3NTMzODQ5MjY1NDYzMTQzNDQ0OTM4ODM2ODc3OTM5ODc1MDA4MTE4Nzg0NzIxMzk4MjYzNTQ3NTk4NjM0NTQzMjA3NjA4Nzk1MTM1NDI5NjcwNDY1OTE2MTk5Njg4MTYzNzkwNzcyODA3NDE2NzY0Njk1MDcxMDc0MTI2MjY3ODMyMjYyMzM5MjIzMjc1MjA1MzUxNDE2MTAwNzIwOTUzOTA4NDQ2Njc1MjE3NjMxODA5Mzk5MDE2ODkwMTQ3MzIyNzgyNzY3NDA1NzA1Njk3OTA4MzM3Njc5OTk0MTE4ODIzNzI4NzEzNDQxMjI4NDg3NjgwMzQyNTM2MzIyOTQxNDg3Mjc0MDY0NjYzOTU5MTg1OTQ5NjQ3NTQ3OTc0Mjk0OTQ2ODU5NDMxMTA1MTA0NjkzMzQ0ODYwMDY4MzM4ODQ3NzkzMzg3NTk3NTA0MTM=',
    //                                 'eventIDaddAccInfo'],
                
    //             readOnly: false
    //         };

    //         await this.sutAdapter.sendRequests(request);
    //     }
    // }

    async submitTransaction() {
        // const randomId = Math.floor(Math.random()*this.roundArguments.assets);
        let lpkhash='ZDU3MjQ0YWIxMGE3NmJlMTM5MzdkYjc2NjQ4ZjUwNGY0MzdlNmUwYTdjMzEwYjQwNzNkNmMwYmRjZTNjNGVmNQ=='
        let acc='MTg4ODY0MzcyMDg3NDkzMTU5ODE2NTE1MzIyNDQ4OTEyODI1OTA4OTg0ODk0NzYzNDIzNDcxMjA0ODgxMDg2NjAwOTQxNDA3NDg4NzQ0MDk3OTg5ODg4ODY5NTI2ODY5ODU4NDEwMTM2ODc2OTExNTcwMzI3MjY2OTMyMDk3NjgzNjY5NDk2MjIyOTgxNjMwMDQ0OTk3NDc2NzU1MjU5NDY0MDcxMjc4Njg2MTg2NjMyMzU1MTE5MzYzOTY0OTM5NDMyNTQyNjg4MzcyNzczNTU4NTY1MzkwMzQ0MTA0NzUwNTM2ODUxNDM2NDYyMzE0NzkwODU2NDIyODU3MDExODc1NzU4NDg5ODA0MDQxMzUzODE0ODMzMjg3NTMzODQ5MjY1NDYzMTQzNDQ0OTM4ODM2ODc3OTM5ODc1MDA4MTE4Nzg0NzIxMzk4MjYzNTQ3NTk4NjM0NTQzMjA3NjA4Nzk1MTM1NDI5NjcwNDY1OTE2MTk5Njg4MTYzNzkwNzcyODA3NDE2NzY0Njk1MDcxMDc0MTI2MjY3ODMyMjYyMzM5MjIzMjc1MjA1MzUxNDE2MTAwNzIwOTUzOTA4NDQ2Njc1MjE3NjMxODA5Mzk5MDE2ODkwMTQ3MzIyNzgyNzY3NDA1NzA1Njk3OTA4MzM3Njc5OTk0MTE4ODIzNzI4NzEzNDQxMjI4NDg3NjgwMzQyNTM2MzIyOTQxNDg3Mjc0MDY0NjYzOTU5MTg1OTQ5NjQ3NTQ3OTc0Mjk0OTQ2ODU5NDMxMTA1MTA0NjkzMzQ0ODYwMDY4MzM4ODQ3NzkzMzg3NTk3NTA0MTM='
        let eventID='eventIDaddAccInfo'
        const myArgs = {
            contractId: 'zero_gnark',
            contractFunction: 'addAccInfo',
            invokerIdentity: 'User1',
            contractArguments: [lpkhash,acc,eventID], 
            readOnly: true
        };

        await this.sutAdapter.sendRequests(myArgs);
    }

    async cleanupWorkloadModule() {
      let lpkhash='ZDU3MjQ0YWIxMGE3NmJlMTM5MzdkYjc2NjQ4ZjUwNGY0MzdlNmUwYTdjMzEwYjQwNzNkNmMwYmRjZTNjNGVmNQ=='
      let eventID='eventIDdeleteAccInfo'  
      for (let i=0; i<this.roundArguments.assets; i++) {
            const assetID = `${this.workerIndex}_${i}`;
            console.log(`Worker ${this.workerIndex}: Deleting Info ${assetID}`);
            const request = {
                contractId: this.roundArguments.contractId,
                contractFunction: 'deleteinfo',
                invokerIdentity: 'User1',
                contractArguments: [lpkhash,eventID],
                readOnly: false
            };

            await this.sutAdapter.sendRequests(request);
        }
    }
}

function createWorkloadModule() {
    return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;