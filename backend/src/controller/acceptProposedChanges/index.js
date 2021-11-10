const recoverMNO = require("../../utils/recoverMNO");
const readBuffer = require("../../utils/buffer/readBuffer");
const invokeEvents = require("../../utils/invoke/invokeEvents");

const acceptProposedChanges = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "acceptProposedChanges";

        let user = data.userDetails;
        if (!user) {
            res.sendStatus(201);
            res.end("201");
            return
        }

        let mno = await recoverMNO(user.username)
        const selectEnv = 1;
        const objs = await readBuffer(selectEnv);
        var output = objs.filter(function (obj) { return obj.ra_id == data.list[0].ra_id })
        var articles = output[0].articles
        var article = (articles).filter(function (article) { return article.articleId == data.selectedRow.articleId })
        if (mno == article[0].proposedBy) {
            res.sendStatus(202);
            res.end("202");
            return
        }

        if (data.formVariables[0].key === '' && data.formVariables[0].value === '' && data.formCustomText[0].value === '') {
            console.log("here");
            //if (!eventHf[0]) {
            //    res.sendStatus(403);
            //    res.end("403");
            //    return
            //}
            //await updatePROD(eventHf[1])
        }
        else {
            res.sendStatus(203);
            res.end("203");
            return
        }
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { acceptProposedChanges };
