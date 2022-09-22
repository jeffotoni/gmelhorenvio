import http from 'k6/http';
import { SharedArray } from "k6/data";
import { sleep } from 'k6';

export default function() {
    var url = `http://localhost:8080/v1/frete/calc`;
    var params = {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': __ENV.API_STATIC_TOKEN
        }
    };

    const data = {
        from: {
            postal_code: "90570020"
        },
        to: {
            postal_code: "90570020"
        },
        package: {
            height: 4,
            width: 12,
            length: 17,
            weight: 0.3
        }
    };
    http.post(url, JSON.stringify(data), params);
}