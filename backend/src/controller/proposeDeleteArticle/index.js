const invoke = require("../../utils/invoke/invoke");

const proposeDeleteArticle = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "proposeDeleteArticle";
        let value = data.delteArticle;
        users = await invoke(method, value);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeDeleteArticle };
