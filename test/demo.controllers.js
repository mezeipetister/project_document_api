import chai from 'chai';
import request from 'supertest';
import Server from '../server';

export default describe('Hektor', () => {
    it('should get Hektor', () =>
      request(Server)
        .get('/api/v1/hektor')
        .then(function(res){
            chai.expect(res.text)
            .to.be.equal('Hektor');
        }));
  });