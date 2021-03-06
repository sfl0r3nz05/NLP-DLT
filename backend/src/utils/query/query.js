const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const dotenv = require('dotenv');

const configPath = path.resolve("data", 'config.json');
const configJSON = fs.readFileSync(configPath, 'utf8');
const config = JSON.parse(configJSON);

let ccpPath;
let ccpJSON;
let ccp;

module.exports = async function query(method, noArgs, arg, user) {
    try {
        let result
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
        //console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists(user);
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: user, discovery: { enabled: false } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(config.channel.channelName);

        // Get the contract from the network.
        const contract = network.getContract(config.channel.contract);
        //console.log(contract);

        // Submit the transaction.
        if (noArgs == 0) {
            //console.log(`Querying value for key:`)
            result = await contract.evaluateTransaction(method);
            //console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        } else if (noArgs == 1) {
            //console.log(`Querying value for key:`, arg)
            //let arg_ = JSON.stringify(arg);
            result = await contract.evaluateTransaction(method, arg);
            //console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        }
        // Disconnect from the gateway.
        await gateway.disconnect();
        return result.toString();
    } catch (error) {
        return error;
    }
};