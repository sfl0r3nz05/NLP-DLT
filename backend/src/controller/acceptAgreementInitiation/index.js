const recoverMNO = require("../../utils/recoverMNO");
const updatePROD = require("../../utils/data/updatePROD");
const readBuffer = require("../../utils/buffer/readBuffer");
const invokeEvents = require("../../utils/invoke/invokeEvents");

const acceptAgreementInitiation = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let userDetails = data.userDetails;
        let mno1 = await recoverMNO(userDetails.username)
        let mno2 = data.acceptAgreement.mno;
        if (!mno2) {
            res.sendStatus(201);
            res.end("201");
            return
        } else if (mno1 == mno2) {
            res.sendStatus(202);
            res.end("202");
            return
        }
        const selectEnv = 1;
        const objs = await readBuffer(selectEnv);
        var output = objs.filter(function (obj) { return obj.mno1 == mno1 && obj.mno2 == mno2 })
        if (output[0]) {
            res.sendStatus(202);
            res.end("202");
            return
        }
        var output = objs.filter(function (obj) { return ((obj.mno1 == mno1 && obj.mno2 == mno2) || (obj.mno1 == mno2 && obj.mno2 == mno1)); })
        let arg1 = output[0].ra_id
        let arg2 = "";
        let arg3 = "";
        noArgs = 1
        let method = "acceptAgreementInitiation";
        let event_name = "confirmation_ra_started"
        eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, arg3, userDetails.username);
        if (!eventHf[0]) {
            res.sendStatus(403);
            res.end("403");
            return
        }
        await updatePROD(eventHf[1])
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { acceptAgreementInitiation };
