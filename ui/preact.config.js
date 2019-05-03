export default (config, env, helpers) => {
    config.devServer = {
        quiet: true,
        proxy: [
            {
                path: '/**',
                target: 'http://localhost:8080',
                // ...any other stuff...
            }
        ]
    }
}