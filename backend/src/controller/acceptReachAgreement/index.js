const invoke = require("../../utils/invoke/invoke");

const acceptReachAgreement = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "acceptAgreement";
        let value = data.acceptAgreement;
        users = await invoke(method, value);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { acceptReachAgreement };
