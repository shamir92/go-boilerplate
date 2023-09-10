import http from 'k6/http';
import {sleep} from 'k6';

export const options = {
  // Key configurations for avg load test in this section
  stages: [
    { duration: '30s', target: 100 }, // traffic ramp-up from 1 to 100 users over 5 minutes.
    { duration: '1m', target: 100 }, // stay at 100 users for 30 minutes
    { duration: '30s', target: 0 }, // ramp-down to 0 users
  ],
};

export default () => {
//   const urlRes = http.get('http://localhost:8080/ping/');
//   const urlRes = http.get('http://localhost:8080/ping/controller');
  const urlRes = http.get('http://localhost:8080/v1/ping/usecase');
//   const urlRes = http.get('http://localhost:8080/ping/repository');

};

