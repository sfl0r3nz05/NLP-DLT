const invoke = require("../../utils/invoke/invoke");

const acceptProposedChanges = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "acceptProposedChanges";
        let value = data.acceptProposedChanges;
        users = await invoke(method, value);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { acceptProposedChanges };
