const query = require("../../utils/query/query");

const recoverMNO = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "recoverMNO";
        let noArg = 0
        users = await query(method, noArg, "", data.user);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { recoverMNO };