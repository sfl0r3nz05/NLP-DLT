const invokeEvents = require("../../utils/invoke/invokeEvents");

const acceptAgreementInitiation = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "acceptAgreementInitiation";
        let arg1 = data.acceptAgreement.mno;
        let arg2 = "";
        let arg3 = "";
        let userDetails = data.userDetails;
        noArgs = 1
        method = "proposeAgreementInitiation";
        value = data.createAgreement;
        let event_name = "confirmation_ra_started"
        eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, arg3, userDetails.username);
        if (!eventHf[0]) {
            res.sendStatus(403);
            res.end("403");
            return
        }
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { acceptAgreementInitiation };
