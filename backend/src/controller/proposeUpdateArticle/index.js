const invoke = require("../../utils/invoke/invoke");

const proposeUpdateArticle = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "proposeUpdateArticle";
        let value = data.updateArticle;
        users = await invoke(method, value);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeUpdateArticle };
