import { check } from "k6";
import http from 'k6/http';

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
        checkSwggerDocs: Object.assign(
            {
                exec: "checkSwaggerDocs",
                env: {
                    URL: `${__ENV.HOST}/api/v1`,
                },
            },
            baseScenario
        ),
    },
};

export const checkSwaggerDocs = () => {
    const url = __ENV.URL;
    const params = {
        headers: {
            "Content-Type": "application/json",
        },
    };

    const requests = {
        docs: {
            method: "GET",
            url,
            params,
        },
    };

    const responses = http.batch(requests);
    const docsResp = responses.docs;

    check(docsResp, {
        "status is 200": (r) => r.status === 200,
    });
};
