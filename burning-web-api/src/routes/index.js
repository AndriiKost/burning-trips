module.exports = ({ router }) => {
    // getting the home route
    router.get('/ping', (ctx, next) => ctx.body = 'pong');
};