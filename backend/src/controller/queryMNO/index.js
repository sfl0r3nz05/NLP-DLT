const query = require("../../utils/query/query");

const queryMNO = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let noArgs = 1
        users = await query(data.method, noArgs, data.mno_name, data.user);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { queryMNO };
