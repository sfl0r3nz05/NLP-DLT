const query = require("../../utils/query/query");
const invokeEvents = require("../../utils/invoke/invokeEvents");
const populatePROD = require("../../utils/data/populatePROD");
const recoverMNO = require("../../utils/recoverMNO");

const proposeAgreementInitiation = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let userDetails = data.userDetails;
        let mno1 = await recoverMNO(userDetails.username)
        let noArgs = 1
        let arg0 = mno1
        let arg1 = mno1
        let arg2 = data.createAgreement.mno2
        let arg3 = data.createAgreement.nameRA
        if (!arg2 || !arg3) {
            res.sendStatus(201);
            res.end("201");
            return
        }
        let method = "queryMNOforID";
        queried_value1 = await query(method, noArgs, arg0, userDetails.username);
        if (queried_value1 == "FALSE" || arg0 == arg2) {
            res.sendStatus(202);
            res.end("402");
            return
        }
        noArgs = 3
        method = "proposeAgreementInitiation";
        let event_name = "started_ra"
        eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, arg3, userDetails.username);
        if (!eventHf[0]) {
            res.sendStatus(403);
            res.end("403");
            return
        }
        await populatePROD(eventHf[1])
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeAgreementInitiation };
