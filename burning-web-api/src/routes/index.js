module.exports = ({ router }) => {

    router.use('/stop', require('./api-routes'));

    // getting the home route
    router.get('/ping', (ctx, next) => ctx.body = 'pong');
};