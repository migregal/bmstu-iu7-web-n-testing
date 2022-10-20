import { check } from "k6";
import http from 'k6/http';

import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

const baseScenario = {
    executor: "constant-arrival-rate",
    rate: 15000,
    timeUnit: "1s",
    duration: "1m",
    gracefulStop: "0s",
    preAllocatedVUs: 100,
    maxVUs: 300,
};

export const options = {
    insecureSkipTLSVerify: true,
    systemTags: ["scenario", "check"],
    scenarios: {
        registrationFlow: Object.assign(
            {
                exec: "registrationFlow",
                env: {
                    URL: `${__ENV.HOST}/api/v1`,
                },
            },
            baseScenario
        ),
    },
};

export const registrationFlow = () => {
    const url = `${__ENV.URL}/registration`;
    const params = {
        headers: {
            "Content-Type": "application/json",
        },
    };

    const registrationData = JSON.stringify({
        "username": randomString(10),
        "email": `${randomString(12)}@neural.com`,
        "password": randomString(12),
        "fullname": randomString(24),
    });

    const requests = {
        regData: {
            method: "POST",
            url,
            params,
            body: registrationData
        },
    };

    const responses = http.batch(requests);
    const regResp = responses.regData;

    check(regResp, {
        "status is 200": (r) => r.status === 200,
    });
};
