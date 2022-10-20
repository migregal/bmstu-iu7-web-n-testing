import { check } from "k6";
import http from 'k6/http';

import { FormData } from 'https://jslib.k6.io/formdata/0.0.2/index.js';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import { URL } from 'https://jslib.k6.io/url/1.0.0/index.js';



const binStructure = open('./structure.json.gz', 'b');
const binWeights = open('./weights.json.gz', 'b');


const baseScenario = {
    executor: "constant-vus",
    duration: "1m",
    vus: 10,
};

export const options = {
    insecureSkipTLSVerify: true,
    systemTags: ["scenario", "check"],
    scenarios: {
        fullUserFlow: Object.assign(
            {
                exec: "fullUserFlow",
                env: {
                    URL: `${__ENV.HOST}/api/v1`,
                },
            },
            baseScenario
        ),
    },
};

export const fullUserFlow = () => {
    const login_url = `${__ENV.URL}/registration`;
    const params = {
        headers: {
            "Content-Type": "application/json",
        },
    };

    const login_body = JSON.stringify({
        "username": randomString(10),
        "email": `${randomString(12)}@neural.com`,
        "password": randomString(12),
        "fullname": randomString(24),
    });

    const login_resp = http.post(login_url, login_body, params)
    check(login_resp, {
        "status is 200": (r) => r.status === 200,
        "token is present": (r) => r.json().token != "",
    });

    const auth_token = login_resp.json().token;
    params.headers["Authorization"] = `Bearer ${auth_token}`;

    const fd = new FormData();
    fd.append('title', randomString(20));
    fd.append('structure', http.file(binStructure, 'structure.json.gz'));
    fd.append('weights', http.file(binWeights, 'weights.json.gz'));

    params.headers["Content-Type"] = `multipart/form-data; boundary= ${fd.boundary}`;

    const upload_url = `${__ENV.URL}/models`;
    const upload_resp = http.post(upload_url, fd.body(), params);

    check(upload_resp, {
        "status is 200": (r) => r.status === 200,
        "id is present": (r) => r.json().id != "",
    });

    const model_id = upload_resp.json().id;
    const delete_url = new URL(`${__ENV.URL}/models`);
    delete_url.searchParams.append('id', model_id);

    delete params.headers["Content-Type"];

    const delete_resp = http.del(delete_url.toString(), null, params)
    check(delete_resp, {
        "status is 200": (r) => r.status === 200,
    });
};
