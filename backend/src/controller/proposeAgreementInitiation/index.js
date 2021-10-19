const invoke = require("../../utils/invoke/invoke");

const proposeAgreementInitiation = async (req, res) => {
    try {
        let data = req.body; // params from POST
        console.log(data);
        let method = "addOrg";
        let value = data.createAgreement;
        users = await invoke(method, value);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeAgreementInitiation };
