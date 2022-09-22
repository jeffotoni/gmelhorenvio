import http from 'k6/http';
import { SharedArray } from "k6/data";
import { sleep } from 'k6';

export default function() {
    var url = `https://gmelhorenvio-ggiobvkpca-uk.a.run.app/v1/frete/calc`;
    var params = {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': __ENV.API_STATIC_TOKEN
        }
    };

    const data = {
        from: {
            postal_code: "96020360"
        },
        to: {
            postal_code: "01018020"
        },
        products: [{
            id: "x100001",
            height: 4,
            width: 12,
            length: 17,
            weight: 0.3,
            insurance_value: 10.1,
            quantity: 1
        }],
        options: {
            receipt: false,
            own_hand: false
        },
        services: "1,2,18"
    };
    http.post(url, JSON.stringify(data), params);
}
