const query = require("../../utils/query/query");
const invoke = require("../../utils/invoke/invoke");

const proposeAgreementInitiation = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "queryMNOforID";
        let arg0 = data.createAgreement.mno2;
        let userDetails = data.userDetails;
        queried_value = await query(method, arg0, userDetails.username);
        if (queried_value == "True") {
            res.sendStatus(202);
            res.end("402");
        }
        let noArgs = 3
        let arg1 = "TELEFONICA"
        let arg2 = data.createAgreement.mno2
        let arg3 = data.createAgreement.nameRA
        method = "proposeAgreementInitiation";
        value = data.createAgreement;
        received_value = await invoke(method, noArgs, arg1, arg2, arg3, userDetails.username);
        if (!received_value) {
            res.sendStatus(403);
            res.end("403");
        }
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeAgreementInitiation };
