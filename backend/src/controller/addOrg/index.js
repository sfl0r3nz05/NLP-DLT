const populateLER = require("../../utils/data/populateLER");
const invoke = require("../../utils/invoke/invoke");

const addOrg = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "addOrg";
        let value = data.feature;
        let userDetails = data.userDetails;
        let noArg = 1
        value = await invoke(method, noArg, value, "", "", userDetails.username);
        if (!value) {
            res.sendStatus(403);
        }
        await populateLER(value)
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { addOrg };
