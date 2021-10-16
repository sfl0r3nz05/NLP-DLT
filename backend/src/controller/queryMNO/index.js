const query = require("../../utils/query/query");

const queryMNO = async (req, res) => {
    try {
        let data = req.body; // params from POST
        //let value = data.feature;
        console.log(data.mno_name);
        users = await query(data.mno_name);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { queryMNO };
