var fs = require('fs')
const readBuffer = require("../buffer/readBuffer");

module.exports = async function populatePROD(valueToUpdate) {
    try {
        value = JSON.parse(valueToUpdate);
        selectEnv = 4;
        countries = await readBuffer(selectEnv);
        country1 = countries.find((country1) => country1.name.toUpperCase() === value.country1);
        country2 = countries.find((country2) => country2.name.toUpperCase() === value.country2);

        fs.readFile(__dirname + "/../../data/listOfMNOs.json", function (err, data) {
            var obj = JSON.parse(data)
            obj.push({
                "mno1": value.mno1,
                "country_mno1": country1.flag,
                "mno2": value.mno2,
                "country_mno2": country2.flag,
                "ra_name": value.raname,
                "ra_id": value.raid,
                "ra_status": value.rastatus,
                "timestamp": value.timestamp
            })
            var json = JSON.stringify(obj)
            fs.writeFile(__dirname + "/../../data/listOfMNOs.json", json, function (err) {
                if (err) throw err;
            })
        })
    } catch (error) {

    }
}