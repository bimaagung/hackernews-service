module.exports = {
  apps: [
    {
      name: 'hackernews-service',
      script: './bin/main',
      watch: true,
      ignore_watch: ['node_modules'],
      watch_options: {
        followSymlinks: false
      },
      env: {
        GOPATH: '/home/ubuntu/go',
      },
      env_production: {
        GOPATH: '/home/ubuntu/go' 
      },
    },
  ],
};
