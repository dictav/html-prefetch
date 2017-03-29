import npm from 'rollup-plugin-node-resolve'
import commonjs from 'rollup-plugin-commonjs'

export default {
  entry: 'jssrc/beacon.js',
  moduleName: "beacon",
  format: "umd",
  plugins: [
    npm({
      jsnext: true, // if provided in ES6
      main: true, // if provided in CommonJS
      browser: true,  // if provided for browsers
      extensions: [ '.js' ]
    }),
    commonjs()
  ]
}
