const recoverMNO = require("../../utils/recoverMNO");
const updatePROD = require("../../utils/data/updatePROD");
const readBuffer = require("../../utils/buffer/readBuffer");
const invokeEvents = require("../../utils/invoke/invokeEvents");

const acceptAgreementInitiation = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let userDetails = data.userDetails;
        let mno1 = await recoverMNO(userDetails.username)
        let mno2 = data.list[(data.list).length - 1].mno2;
        if (mno1 === mno2) { mno2 = data.list[(data.list).length - 1].mno1 }
        if (!mno2) {
            res.sendStatus(201);
            res.end("201");
            return
        }

        console.log(mno1);
        console.log(mno2);

        const selectEnv = 1;
        const objs = await readBuffer(selectEnv);
        var output = objs.filter(function (obj) { return ((obj.mno1 == mno1 && obj.mno2 == mno2) || (obj.mno1 == mno2 && obj.mno2 == mno1)); })
        var articles = output[0].articles
        //console.log(output[0]);

        if (articles.length === 0) {
            let arg1 = output[0].ra_id
            let arg2 = "";
            let arg3 = "";
            let arg4 = "";
            let arg5 = "";
            let arg6 = "";
            let arg7 = "";
            noArgs = 1
            let method = "acceptAgreementInitiation";
            let event_name = "confirmation_ra_started"
            eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, arg3, arg4, arg5, arg6, arg7, userDetails.username);
            if (!eventHf[0]) {
                res.sendStatus(403);
                res.end("403");
                return
            }
            await updatePROD(eventHf[1])
        } else if ((articles.length > 0) && (output[0].ra_status != "accepted_ra")) {
            let arg1 = output[0].ra_id
            let arg2 = "";
            let arg3 = "";
            let arg4 = "";
            let arg5 = "";
            let arg6 = "";
            let arg7 = "";
            noArgs = 1
            let method = "proposeReachAgreement";
            let event_name = "proposal_accept_ra"
            eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, arg3, arg4, arg5, arg6, arg7, userDetails.username);
            if (!eventHf[0]) {
                res.sendStatus(403);
                res.end("403");
                return
            }
            await updatePROD(eventHf[1])
            res.sendStatus(204);
            res.end("204");
            return
        } else if ((articles.length > 0) && (output[0].ra_status === "accepted_ra") && (output[0].acceptRAproposedBy != mno1)) {
            let arg1 = output[0].ra_id
            let arg2 = "";
            let arg3 = "";
            let arg4 = "";
            let arg5 = "";
            let arg6 = "";
            let arg7 = "";
            noArgs = 1
            let method = "acceptReachAgreement";
            let event_name = "confirmation_accepted_ra"

            eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, arg3, arg4, arg5, arg6, arg7, userDetails.username);
            if (!eventHf[0]) {
                res.sendStatus(403);
                res.end("403");
                return
            }
            await updatePROD(eventHf[1])
            res.sendStatus(205);
            res.end("205");
            return
        } else {
            res.sendStatus(203);
            res.end("203");
            return
        }

        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { acceptAgreementInitiation };
