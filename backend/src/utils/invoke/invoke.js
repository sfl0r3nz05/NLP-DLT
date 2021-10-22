const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const dotenv = require('dotenv');
const { Console } = require('console');

const configPath = path.resolve("data", 'config.json');
const configJSON = fs.readFileSync(configPath, 'utf8');
const config = JSON.parse(configJSON);

let ccpPath;
let ccpJSON;
let ccp;

module.exports = async function invoke(method, noArgs, arg1, arg2, arg3, user) {
    try {
        let payLoad
        let listener
        let arg_1 = JSON.stringify(arg1);
        let arg_2 = JSON.stringify(arg2);
        let arg_3 = JSON.stringify(arg3);

        dotenv.config();
        if (process.env.NETWORK != undefined) {
            config.connection_profile = config.connection_profile.replace("basic", process.env.NETWORK);
        }
        if (process.env.CHANNEL != undefined) {
            config.channel.channelName = config.channel.channelName.replace("mychannel", process.env.CHANNEL);
        }

        ccpPath = path.resolve("data", config.connection_profile);
        ccpJSON = fs.readFileSync(ccpPath, 'utf8');
        ccp = JSON.parse(ccpJSON);

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);

        // Check to see if we've already enrolled all the users.
        var userExists = await wallet.exists(user);
        if (!userExists) {
            console.log('An identity for the user does not exist in the wallet: ', user);
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: user, discovery: { enabled: false } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(config.channel.channelName);
        //console.log(network);

        // Get the contract from the network.
        const contract = network.getContract(config.channel.contract);

        await contract.addContractListener('my-contract-listener', 'created_org', (err, event, blockNumber, transactionId, status) => {
            if (err) {
                console.error(err);
                return;
            }
            listener = event
            console.log(`Event: ${event.payload.toString()} Block Number: ${blockNumber} Transaction ID: ${transactionId} Status: ${status}`);
        })

        // Submit the transaction.
        if (noArgs == 1) {
            try {
                payLoad = await contract.submitTransaction(method, arg_1);
                console.log(`Transaction has been submitted: ${user}\t${method}\t${arg_1}`);
                console.log(payLoad.toString())
            } catch (error) {
                return false
            }
        } else if (noArgs == 2) {
            try {
                payLoad = await contract.submitTransaction(method, arg_1, arg2);
                console.log(`Transaction has been submitted: ${user}\t${method}\t${arg_1}\t${arg_2}`);
                console.log(payLoad.toString())
            } catch (error) {
                return false
            }
        } else if (noArgs == 3) {
            try {
                payLoad = await contract.submitTransaction(method, arg1, arg2, arg3);
                console.log(`Transaction has been submitted: ${user}\t${method}\t${arg_1}\t${arg_2}\t${arg_3}`);
                console.log(payLoad.toString())
            } catch (error) {
                return false
            }
        }
        // Disconnect from the gateway.
        await gateway.disconnect();
        console.log(listener.payload.toString());
        return payLoad.toString();
    } catch (error) {
        return error;
    }
};