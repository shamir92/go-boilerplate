import http from 'k6/http';
import { check, sleep} from 'k6';

export const options = {
  vus: 3, // Key for Smoke test. Keep it at 2, 3, max 5 VUs
  duration: '2s', // This can be shorter or just a few iterations
};

export default () => {
  const urlRes1 = http.get('http://localhost:8080/v1/invitation-status');
  sleep(1);
  const urlRes2 = http.get('http://localhost:8080/v1/invitation-status/8f00b0d2-6da5-4c1a-8c99-e294f6541a16');
  sleep(1);
  const urlRes3 = http.get('http://localhost:8080/v1/gathering/type');
  sleep(1);
  const urlRes4 = http.get("http://localhost:8080/v1/gathering/type/99ef6d74-f8c2-4663-9105-76aaec388b4c")

  // MORE STEPS
  // Here you can have more steps or complex script
  // Step1
  // Step2
  // etc.
};