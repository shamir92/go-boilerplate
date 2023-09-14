import http from 'k6/http';
import { check, sleep} from 'k6';

export const options = {
    vus: 3, // Key for Smoke test. Keep it at 2, 3, max 5 VUs
    duration: '2s', // This can be shorter or just a few iterations
};

export default () => {
    const urlRes = http.get('http://localhost:8080/v1/ping/');
    sleep(1);

    const urlRes1 = http.get('http://localhost:8080/v1/ping/controller');
    sleep(1);

    const urlRes2 = http.get('http://localhost:8080/v1/ping/usecase');
    sleep(1);

    const urlRes3 = http.get("http://localhost:8080/v1/ping/repository")
    sleep(1)

  // MORE STEPS
  // Here you can have more steps or complex script
  // Step1
  // Step2
  // etc.
};