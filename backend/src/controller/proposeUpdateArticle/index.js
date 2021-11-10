const recoverMNO = require("../../utils/recoverMNO");
const updatePROD = require("../../utils/data/updatePROD");
const readBuffer = require("../../utils/buffer/readBuffer");
const invokeEvents = require("../../utils/invoke/invokeEvents");

const proposeUpdateArticle = async (req, res) => {
    try {
        let data = req.body; // params from POST
        console.log(data);
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

        if (data.formVariables[0].key != '' || data.formVariables[0].value != '' || data.formCustomText[0].value != '') {
            let method = "proposeUpdateArticle";
            let event_name = "proposed_update_article";
            let noArgs = 7
            let arg1 = output[0].ra_id;
            let arg2 = article[0].articleId;
            let arg3 = data.selectedRow.articleName
            let value4 = JSON.stringify(data.formVariables)
            console.log(value4);
            let arg4 = Buffer.from(value4).toString('base64');
            let value5 = JSON.stringify(data.selectedArticlesVariation);
            //let arg5 = Buffer.from(value5).toString('base64');
            arg5 = ""
            let value6 = JSON.stringify(data.selectedArticlesStdClause)
            //let arg6 = Buffer.from(value6).toString('base64');
            arg6 = ""
            let value7 = JSON.stringify(data.formCustomText)
            console.log(value7);
            let arg7 = Buffer.from(value7).toString('base64');

            eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, arg3, arg4, arg5, arg6, arg7, user.username);
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

module.exports = { proposeUpdateArticle };
