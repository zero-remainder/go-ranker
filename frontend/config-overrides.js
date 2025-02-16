const path = require('path');

module.exports = function override(config, env) {
    // Change the output directory to "../resource/build"
    config.output.path = path.join(__dirname, '../resource/build');
    return config;
};