const invoke = require("../../utils/invoke/invoke");

const proposeAgreementInitiation = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "addOrg";
        let value = data.createAgreement;
        if (data.createAgreement.mno1 != "TELEFONICA" || data.createAgreement.mno2 != "ORANGE") {
            res.sendStatus(400);
        } else {
            users = await invoke(method, value);
            res.sendStatus(200);
        }
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeAgreementInitiation };
