const recoverMNO = require("../../utils/recoverMNO");
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

        let mno = await recoverMNO(user)
        if (mno == data.list.mno2) {
            res.sendStatus(202);
            res.end("202");
            return
        }

        if (data.formVariables[0].key === '' && data.formVariables[0].value === '' && data.formCustomText[0].value === '') {
            //if (!eventHf[0]) {
            //    res.sendStatus(403);
            //    res.end("403");
            //    return
            //}
            //await updatePROD(eventHf[1])
        }
        else {
            res.sendStatus(202);
            res.end("202");
            return
        }
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { acceptProposedChanges };
