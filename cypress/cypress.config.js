const { defineConfig } = require('cypress')

module.exports = defineConfig({
    port: 7002,
    e2e: {
        supportFile: false
    }
})