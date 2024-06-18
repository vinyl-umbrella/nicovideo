export default (config, env, helpers) => {
  if (env.isProd) {
    config.devtool = false;
  }
}
