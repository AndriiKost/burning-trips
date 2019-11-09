'use strict'
import mongoose from 'mongoose';
import Koa from 'koa';
import logger from 'koa-logger';
import Router from 'koa-router';

require('dotenv').config();

const dbConnectionString = process.env.MONGO_DB_CONNECTION_STRING;
const port = process.env.PORT;

mongoose.connect(dbConnectionString, { useNewUrlParser: true, useUnifiedTopology: true })
    .then(dbData => console.log('connected to database ', dbConnectionString))
    .catch(err => handleError(err));

mongoose.Connection.on('error', err => handleError(err));

const app = new Koa();
const router = new Router({
    prefix: '/api'
});

// require('./routes')({ router });

app
    .use(logger())
    .use(router.routes())
    .use(router.allowedMethods())
    .use(async (ctx, next) => globalErrorHandlerMiddleware(ctx, next))
    .listen(port);

console.log(`app listening on port ${port}`);

function handleError(err) {
    console.warn(err);
    // TODO
}

async function globalErrorHandlerMiddleware(ctx, next) {
    try {
      await next();
    } catch (err) {
      ctx.status = err.status || 500;
      ctx.body = err.message;
      ctx.app.emit('error', err, ctx);
    }
  }