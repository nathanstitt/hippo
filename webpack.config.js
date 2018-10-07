const webpack = require('webpack');
const ManifestPlugin = require('webpack-manifest-plugin');
const path = require('path');

const devMode = process.env.NODE_ENV !== 'production';
//const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;

const production = process.env.NODE_ENV === 'production';
const devServerPort = 3808;
const host = process.env.HOST || 'localhost'
const publicURLPath = production ?
    `https://truelinetitle.com/public/webpack/` :
    `http://${host}:${devServerPort}/webpack/`;


module.exports = {
    entry: {
        application: 'application.js',
        home: 'home.js',
    },
    mode: production ? 'production' : 'development',
    module: {
        rules: [
            {
                test: /\.(sa|sc|c)ss$/,
                use: [
                    'style-loader',
                    'css-loader',
                    'sass-loader',
                ],
            }, {
                test: /\.jsx?$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-env']
                    },
                },
            },
        ]
    },
    resolve: {
        extensions: ['*', '.js', '.jsx', '.css'],
        alias: {
            mobx: __dirname + "/node_modules/mobx/lib/mobx.es6.js"
        },
        modules: [
            'node_modules',
            'client',
        ]
    },
    output: {
        path: __dirname + '/public/webpack',
        publicPath: production ?  "/webpack/" : publicURLPath,
        filename: '[name]-[hash].js',
    },
    plugins: [
        new webpack.HotModuleReplacementPlugin(),
        new ManifestPlugin({
            writeToFileEmit: true,
            publicPath: production ? "/public/webpack/" : publicURLPath,
        }),
    ],
    devServer: {
        contentBase: './dist',
        port: devServerPort,
        headers: { 'Access-Control-Allow-Origin': '*' },
        hot: true
    }
};
