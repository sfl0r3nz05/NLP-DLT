const invoke = require("../../utils/invoke/invoke");

const proposeAgreementInitiation = async (req, res) => {
    try {
        let data = req.body; // params from POST
        console.log(data);
        let method = "addOrg";
        let value = data.createAgreement;
        value = await invoke(method, value, userDetails.username);
        if (!value) {
            res.sendStatus(403);
        }
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeAgreementInitiation };
