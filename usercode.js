const moment = require('moment');
const axios = require('axios');

module.exports = (data) => {
    data.now = moment().format('MMMM Do YYYY, h:mm:ss a');

    return data;
};
