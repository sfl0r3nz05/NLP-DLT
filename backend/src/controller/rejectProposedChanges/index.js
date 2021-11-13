const recoverMNO = require("../../utils/recoverMNO");
const updatePROD = require("../../utils/data/updatePROD");
const readBuffer = require("../../utils/buffer/readBuffer");
const invokeEvents = require("../../utils/invoke/invokeEvents");

const rejectProposedChanges = async (req, res) => {
    try {
        let data = req.body; // params from POST
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
            let method = "rejectProposedChanges";
            let event_name = "reject_proposed_changes";
            let noArgs = 2
            let arg1 = output[0].ra_id;
            let arg2 = article[0].articleId;
            eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, "", "", "", "", "", user.username);
            if (!eventHf[0]) {
                res.sendStatus(403);
                res.end("403");
                return
            }
            console.log(eventHf);
            await updatePROD(eventHf[1])
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

module.exports = { rejectProposedChanges };