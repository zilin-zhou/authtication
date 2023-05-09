'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
    }

    async submitTransaction() {
        let rpk='NmUxZjBlYmQ5MjllNjUyYzYzNmU5YzY5MGUwNGI5MTYzZDVkMzlmZWY1NTcwMzVlZjMwNDJkNzIwZWI1MmY5NQ=='
        let rsign='MTEwMzIzMTUxODExMTcyNTg2NDMzMzkyMDY1NTE0MDY2MDE3MTE2MTY0MDQzMzgyMTUxNjY3NjI4NTc0NzMzNzE4ODc1NjkxODQ0Nzc3'
        let ssign='OTc1MjAwODUzNjU5NTgxMTM5MjQ1MTc1NDI1OTUxMTkwODgzNjI0NDc0MzY5OTg1NTc4OTQ3MDgxOTIyNjczODgyMDk2ODM4MTEzMjY='
        let CT='BF0Rgf1x6PsBcOZvaiXQLeWL005bU+Og7IUqA2thEXa6I9YNPPP3pjX1BnR88RZi88e+6XHwvjKi31FfJ/RmmEIwZfZfrjIWsKJLuqYitlOikaCZvepdL6Gz3z6RNigiKLZXUKVOtCzvMl6jLrDwMDk1HbL6H5Q1QJ62j8boi55wXhvq2T/5xKBTgLoSqX3xyajVRbX2+jAsI84YHQBqwGxVTcm4rscs56uzhUxDrhEeFpFL2YvDwS4KYoPTrgFecFSUSKIBDqXArrzwgRI3klbAr8xUXwMezwCbeZvHeC2LR8JErSsVwib7aozDQObQ5acZIQ=='
        let eventID='eventIDaddInfo'
        const myArgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'addInfo',
            invokerIdentity: 'User1',
            contractArguments: [rpk,rsign,ssign,CT,eventID],   
            readOnly: true
        };

        await this.sutAdapter.sendRequests(myArgs);
    }

    async cleanupWorkloadModule() {
        let rpk='NmUxZjBlYmQ5MjllNjUyYzYzNmU5YzY5MGUwNGI5MTYzZDVkMzlmZWY1NTcwMzVlZjMwNDJkNzIwZWI1MmY5NQ=='
        let eventID='eventIDdeleteInfo'
        for (let i=0; i<this.roundArguments.assets; i++) {
            const assetID = `${this.workerIndex}_${i}`;
            console.log(`Worker ${this.workerIndex}: Deleting Info ${assetID}`);
            const request = {
                contractId: this.roundArguments.contractId,
                contractFunction: 'deleteinfo',
                invokerIdentity: 'User1',
                contractArguments: [rpk,eventID],
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