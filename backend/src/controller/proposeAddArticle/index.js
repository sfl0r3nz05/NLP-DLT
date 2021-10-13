const invoke = require("../../utils/invoke/invoke");

const proposeAddArticle = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "proposeAddArticle";
        let value = data.addArticle;
        users = await invoke(method, value);
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeAddArticle };
