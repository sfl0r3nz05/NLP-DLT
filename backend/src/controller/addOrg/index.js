const invoke = require("../../utils/invoke/invoke");

const addOrg = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "addOrg";
        let value = data.feature;
        console.log(value);
        users = await invoke(method, value);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { addOrg };
