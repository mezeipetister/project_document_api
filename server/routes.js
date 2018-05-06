import examplesRouter from './api/controllers/examples/router';
import hektorRouter from './api/controllers/examples/hektor';

export default function routes(app) {
  app.use('/api/v1/examples', examplesRouter);
  app.use('/api/v1/hektor', hektorRouter);
}
