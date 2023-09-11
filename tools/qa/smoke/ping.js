import http from 'k6/http';
import { check, sleep} from 'k6';

export const options = {
    vus: 3, // Key for Smoke test. Keep it at 2, 3, max 5 VUs
    duration: '2s', // This can be shorter or just a few iterations
};

export default () => {
    const urlRes = http.get('http://localhost:8080/v1/ping/');
    console.log(urlRes)
    sleep(1);

    const urlRes1 = http.get('http://localhost:8080/v1/ping/controller');
    console.log(urlRes1)
    sleep(1);

    const urlRes2 = http.get('http://localhost:8080/v1/ping/usecase');
    console.log(urlRes2)
    sleep(1);

    const urlRes3 = http.get("http://localhost:8080/v1/ping/repository")
    console.log(urlRes3)
    sleep(1)

  // MORE STEPS
  // Here you can have more steps or complex script
  // Step1
  // Step2
  // etc.
};