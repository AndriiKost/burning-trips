// export default router.routes();
module.exports = ({ router }) => {
    router.get('/', (ctx, next) => ctx.body = 'pong');
}